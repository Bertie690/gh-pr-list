// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"fmt"

	"github.com/Bertie690/gh-pr-list/utils"
)

func getTemplateFuncs() map[string]any {
	return map[string]any{
		// reverse argument order so hex comes first (for consistency with `autocolor` template func)
		"colorhex":   func(hex, str string) string { return utils.ColorHex(str, hex) },
		"colorstate": colorPrState,
	}
}

// colorPrState returns the color name used to color the given PR.
// It returns an empty string and an error if pr is of an invalid type.
func colorPrState(pr map[string]any) (string, error) {
	switch pr["state"] {
	case "OPEN":
		if drafted, ok := pr["isDraft"].(bool); !ok {
			return "", fmt.Errorf("invalid PR draft status %v passed to colorstate", pr["isDraft"])
		} else if drafted {
			return "gray", nil
		}
		return "green", nil
	case "CLOSED":
		return "red", nil
	case "MERGED":
		return "magenta", nil
	default:
		return "", fmt.Errorf("invalid PR state %v passed to colorstate", pr["state"])
	}
}
