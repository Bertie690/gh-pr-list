package utils

import (
	"errors"
	"os"
)

// IsFile reports whether a file exists at path and is a directory.
// A non-existent file will return false, nil.
func IsFile(path string) (exists bool, err error) {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}