// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

// Array containing default gotestsum arguments for CI builds.
var testArgsCI = []string{
	"tool",
	"gotest.tools/gotestsum",
	"--format=github-actions",
	"--format-hide-empty-pkg",
	"--rerun-fails=2",
	"--rerun-fails-run-root-test",
	"--rerun-fails-report=tmp/test-results/gotestsum-flake-report.txt",
	"--packages ./...",
	"--junitfile=tmp/test-results/go-test-report.xml",
	"--junitfile-hide-empty-pkg",
	"--junitfile-project-name=$PKGNAME",
	"--junitfile-testcase-classname=short",
	"--junitfile-testsuite-name=short",
	"--",
}

// Array containing default gotestsum arguments for local builds.
var testArgsLocal = []string{
	"tool",
	"gotest.tools/gotestsum",
	"--format=testname",
	"--format-hide-empty-pkg",
	"--format-icons=text",
	"--junitfile=tmp/test-results/go-test-report.xml",
	"--junitfile-hide-empty-pkg",
	"--junitfile-project-name=$PKGNAME",
	"--junitfile-testcase-classname=short",
	"--junitfile-testsuite-name=short",
	"--",
}