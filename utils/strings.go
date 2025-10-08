package utils

import (
	"strings"
)

// RemoveWhitespace removes all whitespace from a string.
func RemoveWhitespace(s string) string {
	return strings.Join(strings.Fields(s), "")
}