package cmd

import (
	"testing"
	"time"
)

func TestC_Run(t *testing.T) {
	//she := "/bin/bash"
	c := C{
		//Shell:   she,
		Name:    "/bin/ls",
		Arg:     []string{"-al", "/tmp"},
		Timeout: 2 * time.Second,
	}
	o, e, s := c.Run()
	t.Logf("stdout: %s\n", o)
	t.Logf("stderr: %s\n", e)
	t.Logf("return code: %d\n", s)
}
