package script

import (
	"fmt"
	"opsHeart_agent/service/cmd"
	"time"
)

// Script object.
type Script struct {
	Shell   string
	Name    string
	User    string
	Timeout int // Second
}

// Script exe result.
type Ret struct {
	Name     string
	Start    time.Time
	End      time.Time
	Output   string
	Stderr   string
	ExitCode int
}

func (s *Script) Run() Ret {
	var r Ret
	r.Start = time.Now()
	r.Name = s.Name
	c := cmd.C{
		Name:    "su",
		Arg:     []string{"-", s.User, "-c", fmt.Sprintf("%s %s", s.Shell, s.Name)},
		Timeout: time.Duration(s.Timeout) * time.Second,
	}
	r.Output, r.Stderr, r.ExitCode = c.Run()
	r.End = time.Now()
	return r
}
