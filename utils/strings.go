// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"strings"
)

// RemoveWhitespace removes any and all all whitespace from a string.
// This is distinct from [strings.TrimSpace], which only strips leading and trailing whitespace
// (leaving ones in the middle intact).
func RemoveWhitespace(str string) string {
	return strings.Join(strings.Fields(str), "")
}
