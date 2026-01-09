// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"bytes"
	"errors"

	"github.com/cli/go-gh/v2"
)

// runCmd runs the given `gh` command in a separate process.
// It returns the command's `stdout` output or an error if the command failed.
func runCmd(execArgs ...string) (output *bytes.Buffer, err error) {
	stdout, stderr, err := gh.Exec(execArgs...)
	if err != nil {
		return nil, errors.New(stderr.String())
	}
	return &stdout, nil
}
