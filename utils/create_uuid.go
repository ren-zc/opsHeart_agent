package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func CreateUUID() (string, error) {
	uuidObj := uuid.NewV4()
	v, err := uuidObj.Value()
	if err != nil {
		return "", err
	}
	u := fmt.Sprintf("%s", v)
	return u, nil
}
