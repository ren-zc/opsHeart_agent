package file

import "io/ioutil"

// Read string from file.
func (fd *FDesc) ReadStr() (string, error) {
	b, err := ioutil.ReadFile(fd.Name)
	return string(b), err
}
