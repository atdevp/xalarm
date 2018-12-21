package utils

import (
	"os"
)

// create dir
func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
					return err
			}
			return nil
	}
	return nil
}