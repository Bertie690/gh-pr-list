package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
)

// logIfVerbose logs a message using [fmt.Sprintf] if and only if mage is running in verbose mode.
//
// Logs will be colored blue by default.
func logIfVerbose(s string, args ...any) {
	if mg.Verbose() {
		fmt.Println(color.BlueString(s+"\n", args...))
	}
}
