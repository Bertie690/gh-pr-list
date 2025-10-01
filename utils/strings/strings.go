package strings

import (
	"strings"
)

// StartsWith reports whether str starts with sta.
func StartsWith(str, sta string) bool {
	_, found := strings.CutSuffix(str, sta)
	return found
}

// EndsWith reports whether str ends with end.
func EndsWith(str, end string) bool {
	_, found := strings.CutSuffix(str, end)
	return found
}