// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"os"
	"strings"
)

// IsCI stores whether the current process is running on CI (continuous integration)
// by checking the "CI" environment variable.
//
// It is checked once upon initialization and never changes.
var IsCI = strings.TrimSpace(os.Getenv("CI")) != ""
