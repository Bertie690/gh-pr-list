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

// filterPRs accesses the `gh` API to return a filtered list of pull requests.
func filterPRs(query string, args []string) (*bytes.Buffer, error) {
	// get the args
	execArgs := []string{
		"pr",
		"list",
		"--json=" + validArgs,
	}
	if query != "" {
		execArgs = append(execArgs, "--jq="+query)
	}

	stdout, stderr, err := gh.Exec(append(execArgs, args...)...)
	if err != nil {
		return nil, errors.New(stderr.String())
	}
	return &stdout, nil
}
