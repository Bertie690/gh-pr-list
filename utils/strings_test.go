// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import "testing"

func TestRemoveWhitespace(t *testing.T) {
	tests := []struct {
		name string
		str    string
		want string
	}{
		{
			name: "No whitespace",
			str: "12345",
			want: "12345",
		},
		{
			name: "Leading & Trailing",
			str: "\n\r\t 12345\r\t\n",
			want: "12345",
		},
		{
			name: "Middle whitespace",
			str: "\nf l u\nb\tb\re\n\t r ",
			want: "flubber",
		},
		{
			name: "JSON",
			str: `[
	{
		"foo": 1,
		"mergeable": "CONFLICTING"
	}
]`,
			want: `[{"foo": 1,"mergeable":"CONFLICTING"}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveWhitespace(tt.str); got != tt.want {
				t.Errorf("RemoveWhitespace() output for input %q was %v, wanted %v", tt.str, got, tt.want)
			}
		})
	}
}
