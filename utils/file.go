// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"errors"
	"fmt"
	"os"
)

// ExistsDir reports whether a file exists at path and is a directory.
// A non-existent file or one that is not a directory will return `false, nil`.
//
// dir is always false if err is non-nil.
func ExistsDir(path string) (isDir bool, err error) {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

// AppendFile appends a string or byte slice to the named file, creating it if necessary.
// It returns any error produced.
func AppendFile[S ~string | ~[]byte](path string, data S) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file %q: \n%w", path, err)
	}
	defer f.Close()

	if _, err := f.Write([]byte(data)); err != nil {
		return fmt.Errorf("error appending data to file %q: \n%w", path, err)
	}
	return nil
}
