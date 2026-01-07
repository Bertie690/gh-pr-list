// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

// CreateList creates and outputs a list of PRs based on the provided query, template, and extra CLI args.
// It returns any error produced.
func CreateList(query, template string, args []string) (err error) {
	if err = validateExtraArgs(args); err != nil {
		return err
	}

	fields := getRequiredFields(query, template)
	if fields == "" {
		// Edge case if query and template are both empty
		cmd, err := runCmd("pr", "list")
		if err != nil {
			return err
		}
		fmt.Println(cmd.String())
		return nil
	}

	execArgs := []string{
		"pr",
		"list",
		"--json=" + fields,
	}
	if query != "" {
		execArgs = append(execArgs, "--jq="+query)
	}
	json, err := runCmd(append(execArgs, args...)...)
	if err != nil {
		return err
	}

	output, err := applyTemplate(json, template)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}

// validateExtraArgs performs extra validation to ensure that any extra CLI args are valid.
func validateExtraArgs(args []string) error {
	if slices.ContainsFunc(args, func(s string) bool {
		return strings.Contains(s, "--json")
	}) {
		return errors.New("cannot pass '--json' flag; required fields are inferred from filter/template")
	}

	return nil
}
