package utils

import (
	"os"
	"strings"
)

// IsCI stores whether the current process is running in CI (continuous integration)
// by checking the "CI" environment variable.
//
// It is checked once upon initialization and never changes.
var IsCI = strings.TrimSpace(os.Getenv("CI")) != ""
