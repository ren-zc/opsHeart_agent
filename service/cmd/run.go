package cmd

import (
	"bytes"
	"context"
	"opsHeart_agent/logger"
	"os/exec"
	"syscall"
	"time"
)

// Command object.
type C struct {
	//Shell   string
	Name    string
	Arg     []string
	Timeout time.Duration
}

func (c *C) Run() (string, string, int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Timeout))
	defer cancel()
	cc := exec.Command(c.Name, c.Arg...)

	outPipe := &bytes.Buffer{}
	errPipe := &bytes.Buffer{}
	exitCode := -1
	cc.Stdout = outPipe
	cc.Stderr = errPipe

	var cmdOut string
	var cmdErr string
	ch := make(chan struct{})

	go func() {
		err := cc.Run()
		cmdOut = outPipe.String()
		cmdErr = errPipe.String()
		if err != nil {
			logger.AgentLog.Errorf("Command %v, run error123: %s\n", c, err)
			if exitErr, ok := err.(*exec.ExitError); ok {
				waitStatus := exitErr.Sys().(syscall.WaitStatus)
				exitCode = waitStatus.ExitStatus()
			} else {
				logger.AgentLog.Errorf("Can not get exit status for: %s, %s\n", c.Name, c.Arg)
				if cmdErr == "" {
					cmdErr = err.Error()
				}
			}
		} else {
			waitStatus := cc.ProcessState.Sys().(syscall.WaitStatus)
			exitCode = waitStatus.ExitStatus()
		}
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		_ = cc.Process.Kill()
		_ = cc.Process.Release()
		cmdErr = "timeout"
		exitCode = -2
	case <-ch:
	}
	return cmdOut, cmdErr, exitCode
}
