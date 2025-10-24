// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/cli/go-gh/v2"
)

func CreateList(query, template string, args []string) (err error) {
	json, err := getPrs(args)
	if err != nil {
		return err
	}

	var queried *bytes.Buffer
	if query == "" {
		queried = json
	} else {
		queried, err = filterJSON(json, query)
		if err != nil {
			return err
		}
	}
	output, err := applyTemplate(queried, template)
	if err != nil {
		return err
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
