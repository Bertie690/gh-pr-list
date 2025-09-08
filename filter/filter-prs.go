package filter

import (
	"bytes"

	"github.com/cli/go-gh/v2/pkg/jq"
)

func filterJSON(json *bytes.Buffer, query string) (*bytes.Buffer, error) {
	var b bytes.Buffer
	if err := jq.Evaluate(json, &b, query); err != nil {
		return nil, err
	}
	return &b, nil
}
