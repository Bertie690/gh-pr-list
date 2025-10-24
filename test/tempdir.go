// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Bertie690/gh-pr-list/internal/projectpath"
	"github.com/Bertie690/gh-pr-list/utils"
)

// The directory where temporary diffs and other test results are stored.
//
// Located within the project root directory for non-CI runs to allow for easier manual checking,
// or inside a subdirectory of [os.TempDir] on CI for performance.
var ResultsDir string

func init() {
	if utils.IsCI {
		ResultsDir = filepath.Join(os.TempDir(), "test-results")
	} else {
		ResultsDir = filepath.Join(projectpath.Root, "tmp")
		fmt.Println(ResultsDir)
	}
}
