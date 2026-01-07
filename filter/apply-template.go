// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"bytes"
	"strings"

	"github.com/cli/go-gh/v2/pkg/template"
	"github.com/cli/go-gh/v2/pkg/term"
)

// TODO: Add preset default templates
func applyTemplate(queried *bytes.Buffer, tmpl string) (output string, err error) {
	var out bytes.Buffer

	// Add a trailing `tablerender` call if the template doesn't end with one.
	// This does nothing if the template has no tables, so we should be fine to add it
	if !strings.HasSuffix(tmpl, "{{tablerender}}") {
		tmpl += "{{tablerender}}"
	}
	tm := template.New(&out, getLineMax(), term.FromEnv().IsTrueColorSupported()).Funcs(getTemplateFuncs())
	if err = tm.Parse(tmpl); err != nil {
		return output, err
	}
	if err = tm.Execute(queried); err != nil {
		return output, err
	}
	return out.String(), nil
}

// TODO: Make this configurable
func getLineMax() int {
	env := term.FromEnv()
	if !env.IsTerminalOutput() {
		return 99_999
	}
	width, _, err := env.Size()
	if err != nil || width < 0 {
		return 120
	}
	return width
}
