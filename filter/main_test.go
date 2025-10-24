// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"errors"
	"os"
	"testing"
)

// Cleanup temp JSON files for build processing.
func TestMain(m *testing.M) {
	_ = m.Run()

	// Don't move files on CI runs if the target files already exist
	// (since that likely indicates gotestsum rerunning failed cases)
	if _, err := os.Stat("../tmp/got_filter.jsonl"); os.Getenv("CI") == "" || errors.Is(err, os.ErrNotExist) {
		_ = os.Rename("../tmp/got.jsonl", "../tmp/got_filter.jsonl")
		_ = os.Rename("../tmp/want.jsonl", "../tmp/want_filter.jsonl")
		_ = os.Rename("../tmp/diff.jsonl", "../tmp/diff_filter.jsonl")
	}
}
