// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"bytes"

	"github.com/cli/go-gh/v2/pkg/jsonpretty"
)

func formatJSON(json *bytes.Buffer) (*bytes.Buffer, error) {
	var b bytes.Buffer

	if err := jsonpretty.Format(&b, json, "\t", false); err != nil {
		return nil, err
	}
	return &b, nil
}
