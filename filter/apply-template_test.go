package filter

import (
	"bytes"
	"strings"
	"testing"
)

func Test_applyTemplate(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		tmpl    string
		want    string
		wantErr bool
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := bytes.NewBufferString(tt.json)
			got, gotErr := applyTemplate(b, tt.tmpl)
			if (gotErr != nil) != tt.wantErr {
				if tt.wantErr {
					t.Error("applyTemplate() succeeded unexpectedly")
				} else {
					t.Errorf("applyTemplate() failed: %v", gotErr)
				}
				return
			}

			got = strings.TrimSpace(got)
			if got != tt.want {
				t.Errorf("applyTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
