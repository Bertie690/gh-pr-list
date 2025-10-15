package test

import (
	"os"
	"path"

	"github.com/Bertie690/gh-pr-list/utils"
)

// The directory where temporary diffs and other test results are stored.
//
// Located within the current directory for non-CI runs to allow for easier manual checking,
// or inside a subdirectory of [os.TempDir] on CI for performance.
var ResultsDir string

func init() {
	if utils.IsCI {
		ResultsDir = path.Join(os.TempDir(), "test-results")
	} else {
		ResultsDir = "tmp"
	}
}
