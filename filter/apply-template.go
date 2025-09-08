package filter

import (
	"bytes"

	"github.com/cli/go-gh/v2/pkg/term"
	"github.com/cli/go-gh/v2/pkg/template"
)

func applyTemplate(queried *bytes.Buffer, tmpl string) (output string, err error) {
	var out bytes.Buffer
	tm := template.New(&out, 120, term.FromEnv().IsTrueColorSupported()).Funcs(getTemplateFuncs())
	if err = tm.Parse(tmpl); err != nil {
		return
	}
	if err = tm.Execute(queried); err != nil {
		return
	}
	return out.String(), nil
}
