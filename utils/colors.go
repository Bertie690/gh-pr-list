// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"strconv"
	"strings"

	"github.com/cli/go-gh/v2/pkg/term"
	"github.com/fatih/color"
)

// SprintColorHex colors a string according to a hex code and returns it.
// It supports formatting directives similar to [fmt.Sprintf].
func SprintColorHex(hex string, s string, args ...any) string {
	return color.RGB(rgbToHex(hex)).Sprintf(s, args...)
}

// ColorHex colors a string according to a hex code and returns it.
func ColorHex(hex string, s string) string {
	return color.RGB(rgbToHex(hex)).Sprint(s)
}

func rgbToHex(hex string) (r int, g int, b int) {
	// Remove leading hashtag if found
	hex = strings.TrimLeft(hex, "#")

	values, _ := strconv.ParseInt(hex, 16, 32)

	return int(values >> 16), int((values >> 8) & 0xFF), int(values & 0xFF)
}

func init() {
	// Disable colors if GH's color support says we shouldn't
	color.NoColor = !term.FromEnv().IsTrueColorSupported()
}
