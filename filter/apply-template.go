// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"bytes"
	"fmt"

	"strings"
	"github.com/cli/go-gh/v2/pkg/template"
	"github.com/cli/go-gh/v2/pkg/term"
)

// TODO: Add preset default templates
func applyTemplate(queried *bytes.Buffer, tmpl string) (output string, err error) {
	var out bytes.Buffer

	// Add a trailing `tablerender` call if the template doesn't end with one.
	// This does nothing if the 
	if !strings.HasSuffix(tmpl, "{{tablerender}}") {
		tmpl += "{{tablerender}}";
	}
	tm := template.New(&out, getLineMax(term.FromEnv().IsTerminalOutput()), term.FromEnv().IsTrueColorSupported()).Funcs(getTemplateFuncs())
	if err = tm.Parse(tmpl); err != nil {
		return
	}
	if err = tm.Execute(queried); err != nil {
		return
	}
	return out.String(), nil
}

func getLineMax(isTerm bool) int {
	// TODO: Make this configurable
	if (isTerm) {
		return 120
	}
	return 99999
}
