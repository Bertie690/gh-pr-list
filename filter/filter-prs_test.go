// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"bytes"
	"testing"

	"github.com/Bertie690/gh-pr-list/test"
)

func Test_filterJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		query   string
		want    string
		wantErr bool
	}{
		{
			name: "all mergeable",
			json: `[
	{
		"foo": 1,
		"mergeable": "CONFLICTING"
	},
	{
		"foo": 2,
		"mergeable": "MERGEABLE"
	}
]`,
			query: `map(select(.mergeable == "CONFLICTING"))`,
			want: `[
	{
		"foo": 2,
		"mergeable": "CONFLICTING"
	}
]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBufferString(tt.json)
			got, gotErr := filterJSON(buf, tt.query)
			if (gotErr != nil) != tt.wantErr {
				if tt.wantErr {
					t.Errorf("filterJSON() succeeded unexpectedly!")
				} else {
					t.Errorf("filterJSON() failed: %v", gotErr)
				}
				return
			}
			gotStr := got.String()
			test.CompareAsJSON(t, gotStr, tt.want)
		})
	}
}
