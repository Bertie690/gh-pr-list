// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

// Variable tokens populated during argument pre-processing.
const (
	resultsDirToken = "{{RESULTS_DIR}}"
	pkgNameToken    = "{{PKG_NAME}}"
)

// String containing default gotestsum arguments for CI builds.
const testArgsCI = `tool
gotest.tools/gotestsum
--format=github-actions
--format-hide-empty-pkg
--rerun-fails=2
--rerun-fails-run-root-test
--rerun-fails-report=` + resultsDirToken + `/gotestsum-flake-report.txt
--packages ./...
--junitfile=` + resultsDirToken + `/go-test-report.xml
--junitfile-hide-empty-pkg
--junitfile-project-name=` + pkgNameToken + `
--junitfile-testcase-classname=short
--junitfile-testsuite-name=short
--`

// String containing default gotestsum arguments for local builds, separated by newlines.
const testArgsLocal = `tool
gotest.tools/gotestsum
--format=testname
--format-hide-empty-pkg
--format-icons=text
--junitfile=` + resultsDirToken + `/go-test-report.xml
--junitfile-hide-empty-pkg
--junitfile-project-name=` + pkgNameToken + `
--junitfile-testcase-classname=short
--junitfile-testsuite-name=short
--`
