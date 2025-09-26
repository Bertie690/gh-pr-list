package filter

import (
	"github.com/Bertie690/gh-pr-list/utils"
)

func getTemplateFuncs() map[string]any {
	return map[string]any{
		"colorHex": utils.ColorHex,
	}
}

// func ColorForPRState() string {
// 	switch pr.State {
// 	case "OPEN":
// 		if pr.IsDraft {
// 			return "gray"
// 		}
// 		return "green"
// 	case "CLOSED":
// 		return "red"
// 	case "MERGED":
// 		return "magenta"
// 	default:
// 		return ""
// 	}
// }
