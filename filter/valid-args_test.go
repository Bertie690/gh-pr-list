// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_getRequiredFields(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		template string
		want     string
	}{
		{
			name:     "template only",
			query:    "",
			template: `{{range .}}{{.title}} - {{.url}}{{end}}`,
			want:     "title,url",
		},
		{
			name:     "jq filter only",
			query:    `map(select(.author == "foo" and (.assignees | length > 0) and .labels | contains("bug")))`,
			template: "",
			want:     "author,assignees,labels",
		},
		{
			name:     "filter and template with overlap",
			query:    `.[] | map(select(.state == "OPEN" and .author == "foo"))`,
			template: `{{range .}}{{printf "%s %s" .title .author .mergeable}}{{end}}`,
			want:     "state,author,title,mergeable",
		},
		{
			name:     "complex filter/template",
			query:    `map(select(.state == "OPEN" and .mergeable != "CONFLICTING" and (.isDraft | not) and (.statusCheckRollup | all(.conclusion != "FAILURE"))))`,
			template: `{{range .}}{{tablerow (printf "[%s](<%s>)" .title .url)}}{{end}}`,
			want:     "state,mergeable,isDraft,statusCheckRollup,title,url",
		},
		{
			name:     "uses colorstate",
			query:    `map(select(.state == "OPEN"))`,
			template: `{{range .}}{{tablerow ((autocolor (colorstate .) (printf "#%v" .number)) | hyperlink .url) ((timeago .updatedAt) | autocolor "blue")}}{{end}}`,
			want:     "state,number,url,updatedAt,isDraft",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRequiredFields(tt.query, tt.template)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("getRequiredFields() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
