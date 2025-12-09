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

func CreateList(query, template string, args []string) (err error) {
	if err := validateExtraArgs(args); err != nil {
		return err
	}

	json, err := filterPRs(query, args)
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
		return errors.New("cannot pass --json flag; all fields are enabled by default")
	}

	return nil
}
