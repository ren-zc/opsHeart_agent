package file

import (
	"os"
)

// Append bytes to file, if file is not exist create it.
func (fd *FDesc) Append(dat []byte) error {
	// set default perm
	if fd.Perm <= 0 {
		fd.Perm = os.FileMode(0644)
	}

	f, err := os.OpenFile(fd.Name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, fd.Perm)
	if err != nil {
		return err
	}
	_, err = f.Write(dat)
	return err
}

// Append string to file by call Append().
func (fd *FDesc) AppendStr(s string) error {
	b := []byte(s)
	return fd.Append(b)
}
