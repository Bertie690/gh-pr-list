// package projectpath contains internal variables to get the current project root dir.
//
// It is only useful in testing and should not be used in production.
package projectpath

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../..")
)
