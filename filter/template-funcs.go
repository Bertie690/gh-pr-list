package filter

import (
	"github.com/Bertie690/gh-pr-list/utils"
)

func getTemplateFuncs() map[string]any {
	return map[string]any{
		"colorhex": utils.ColorHex,
		"colorstate": colorPrState,
	}
}

type prStub struct {
	State prState
	IsDraft bool
}

type prState string

const (
	stateOpen prState = "OPEN"
	stateClosed prState = "CLOSED"
	stateMerged prState = "MERGED"
)

func colorPrState(pr prStub) string {
	switch pr.State {
	case stateOpen:
		if pr.IsDraft {
			return "gray"
		}
		return "green"
	case stateClosed:
		return "red"
	case stateMerged:
		return "magenta"
	default:
		return ""
	}
}
