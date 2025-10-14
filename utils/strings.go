// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"strings"
)

// RemoveWhitespace removes all whitespace from a string.
func RemoveWhitespace(s string) string {
	return strings.Join(strings.Fields(s), "")
}