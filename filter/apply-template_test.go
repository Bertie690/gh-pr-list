// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"bytes"
	"strings"
	"testing"
)

func Test_applyTemplate(t *testing.T) {
	tests := []struct {
		name string
		json string
		tmpl string
		want string
	}{
		{
			name: "1 PR URL/Title",
			json: `[
	{
		"title": "Nice PR",
		"number": 12345,
		"mergeable": "CONFLICTING",
		"isDraft": false,
		"state": "OPEN",
		"url": "https://www.youtube.com/watch?v=XfELJU1mRMg"
	}
]`,
			tmpl: `{{range .}}{{tablerow .title (printf "#%v" .number) .url}}{{end}}{{tablerender}}`,
			want: `Nice PR  #12345  https://www.youtube.com/watch?v=XfELJU1mRMg`,
		},
		{
			name: "No tablerender",
			json: `[
	{
		"title": "Nice PR",
		"number": 12345,
		"mergeable": "CONFLICTING",
		"isDraft": false,
		"state": "OPEN",
		"url": "https://www.youtube.com/watch?v=XfELJU1mRMg"
	}
]`,
			tmpl: `{{range .}}{{tablerow .title (printf "#%v" .number) .url}}{{end}}`,
			want: `Nice PR  #12345  https://www.youtube.com/watch?v=XfELJU1mRMg`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bytes.NewBufferString(tt.json)
			got, gotErr := applyTemplate(b, tt.tmpl)
			if gotErr != nil {
				t.Fatalf("applyTemplate() failed: %v", gotErr)
			}

			got = strings.TrimSpace(got)
			if got != tt.want {
				t.Errorf("applyTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
