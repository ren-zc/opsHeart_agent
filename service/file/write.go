package file

import (
	"io/ioutil"
	"os"
)

// Truncate file and write data to file, create it if not exist.
func (fd *FDesc) WriteStr(s string) error {
	// set default perm
	if fd.Perm <= 0 {
		fd.Perm = os.FileMode(0644)
	}

	b := []byte(s)
	return ioutil.WriteFile(fd.Name, b, fd.Perm)
}
