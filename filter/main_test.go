package filter

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run();

	// Don't move files on CI runs if the target files already exist
	// (since that likely indicates gotestsum rerunning failed cases)
	if _, err := os.Stat("../tmp/got_filter.jsonl"); os.Getenv("CI") == "" || errors.Is(err, os.ErrNotExist) {
		fmt.Println("Moving diff files after filter package run")
		os.Rename("../tmp/got.jsonl", "../tmp/got_filter.jsonl")
		os.Rename("../tmp/want.jsonl", "../tmp/want_filter.jsonl")
		os.Rename("../tmp/diff.jsonl", "../tmp/diff_filter.jsonl")
	}
}