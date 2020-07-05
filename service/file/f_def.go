package file

import "os"

// File object.
type FDesc struct {
	Name string
	Perm os.FileMode
}
