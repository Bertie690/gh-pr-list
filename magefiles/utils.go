// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
)

// logIfVerbose logs a message using [fmt.Sprintf] if and only if mage is running in verbose mode.
//
// Logs will be colored blue by default, and will have a trailing newline appended.
func logIfVerbose(s string, args ...any) {
	if mg.Verbose() {
		fmt.Println(color.BlueString(s, args...))
	}
}
