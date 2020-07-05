package script

import "testing"

// sudo go test to test
func TestScript_Run(t *testing.T) {
	spt := Script{
		Shell:   "/bin/bash",
		Name:    "/tmp/echo123.sh",
		User:    "jacenren",
		Timeout: 5,
		//Timeout: 2, // timeout
	}

	r := spt.Run()
	t.Logf("r is: %v", r)
}
