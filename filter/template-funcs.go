package filter

import "github.com/Bertie690/gh-pr-list/utils"

func getTemplateFuncs() map[string]any {
	return map[string]any{
		"colorHex": utils.ColorHex,
	}
}