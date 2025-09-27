package filter

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/cli/go-gh/v2"
)

func CreateList(query string, template string, args []string) (err error) {
	if len(args) > 0 {
		fmt.Printf("Using extra args %s\n", strings.Join(args, ","))
	}
	json, err := getPrs(args)
	if err != nil {
		return
	}

	var queried *bytes.Buffer
	if query == "" {
		queried = json
		} else {
			queried, err = filterJSON(json, query)
			fmt.Println(queried.String())
		if err != nil {
			return
		}
	}
	output, err := applyTemplate(queried, template)
	if err != nil {
		return
	}
	fmt.Println(output)
	return nil
}

func getPrs(args []string) (*bytes.Buffer, error) {
	if err := validateExtraArgs(args); err != nil {
		return nil, err
	}
	stdout, stderr, err := gh.Exec(append([]string{"pr", "list", "--json=" + strings.Join(validArgs, ",")}, args...)...)
	if err != nil {
		return nil, errors.New(stderr.String())
	}
	return &stdout, nil
}

// validateExtraArgs performs extra validation to ensure that any extra CLI args are valid.
func validateExtraArgs(args []string) error {
	if slices.ContainsFunc(args, func(s string) bool {
		return strings.Contains(s, "--json")
	}) {
		return errors.New("cannot pass --json flag; all fields are enabled by default")
	}

	return nil
}
