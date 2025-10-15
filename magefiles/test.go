// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/Bertie690/gh-pr-list/test"
	"github.com/Bertie690/gh-pr-list/utils"
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Run golangci-lint code quality checks.
func Lint() error {
	path, err := exec.LookPath("golangci-lint")
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("golangci-lint executable not found in PATH")
	} else if err != nil {
		return fmt.Errorf("Unknown error found when retrieving golangci-lint path: %w", err)
	}

	fmt.Println("Found golangci-lint binary at:", path)

	cmd := exec.Command("golangci-lint", "run")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Run backend tests using gotestsum, passing any provided args to "go test".
// CI runs will always run all tests across all packages,
// whereas non-CI runs can specify which package(s) to run as part of goTestArgs.
// If a package identifier is omitted on non-CI runs,
// it will default to running everything ("./...").
func Test(goTestArgs string) error {
	// NB: This _should_ take ...string as an argument, but mage doesn't support variadic arguments at the moment.
	// A workaround is passing a "" as the first argument

	args := getTestArgs(goTestArgs)

	// Merge together any temporary json files together once we're done testing.
	// We do this now to save time - if the prior steps fail,
	// there won't be any JSON files to merge
	defer func() {
		if err := Merge_Temp_JSON(); err != nil {
			fmt.Printf("error merging temp JSON diffs after test run:\n%v\n", err)
		}
	}()

	logIfVerbose("Running gotestsum with args: %q\n", args)

	// Exit immediately if the command ran and a non-zero exit code was reported.
	// This avoids printing a "XYZ failed" error message on top of the gotestsum report
	ran, err := sh.Exec(nil, os.Stdout, os.Stderr, "go", args...) // "go", "tool", "gotest.tools/gotestsum", ...
	if ran && mg.ExitStatus(err) != 0 {
		os.Exit(mg.ExitStatus(err))
	}
	return err
}

// getTestArgs obtains the arguments passed to gotestsum.
func getTestArgs(goTestArgs string) []string {
	var baseArgsStr string
	// Change baseline args if on CI (different reporter for gh actions, etc.)
	if utils.IsCI {
		logIfVerbose("CI run detected; using CI config")
		baseArgsStr = testArgsCI
	} else {
		logIfVerbose("Non-CI run detected; using normal config")
		baseArgsStr = testArgsLocal
	}

	baseArgs := processBaseArgs(baseArgsStr)
	cliArgs := strings.Fields(goTestArgs)

	// If the user forgot to add a package marker for non-CI runs,
	// do them a favor rather than outright failing.
	// CI runs are exempt from this due to `rerun-fails` requiring an explicit package argument
	// (not to mention their entire *job* is to test everything)
	if !utils.IsCI && slices.IndexFunc(cliArgs, func(s string) bool {
		return strings.HasPrefix(s, "./")
	}) == -1 {
		logIfVerbose(color.BlueString("No package identifier found; defaulting to running everything"))
		cliArgs = append(cliArgs, "./...")
	}

	return append(baseArgs, cliArgs...)
}

func processBaseArgs(baseArgs string) []string {
	// Get package name for JUnit XML reports, deferring to $GITHUB_REPOSITORY env var if present
	pkgName := "gh-pr-list"
	if r := strings.TrimSpace(os.Getenv("GITHUB_REPOSITORY")); r != "" {
		pkgName = r
	}

	// Replace tokens within base args
	baseArgs = strings.ReplaceAll(baseArgs, pkgNameToken, pkgName)
	baseArgs = strings.ReplaceAll(baseArgs, resultsDirToken, test.ResultsDir)

	return strings.Fields(baseArgs)
}

// Remove all temporary json files produced during tests and merge them together.
// This takes all files matching the format "XXX_**.jsonl",
// and merges them together into a single file named "XXX.jsonl".
// Comments are added between failing tests from different packages.
func Merge_Temp_JSON() error {
	tmpdir, err := os.Open(test.ResultsDir)
	// No temp dir makes our job easy
	if errors.Is(err, os.ErrNotExist) {
		return nil
	} else if err != nil {
		return mg.Fatalf(1, "error while opening temp folder: \n%w", err)
	}
	fileNames, err := tmpdir.Readdirnames(-1)
	if err != nil {
		return mg.Fatalf(1, "error while reading temp folder files: \n%w", err)
	}

	count := 0
	for _, fileName := range fileNames {
		fullPath := filepath.Join(test.ResultsDir, fileName)
		// skip non-JSONL files or ones without underscores
		if !strings.HasSuffix(fileName, ".jsonl") {
			continue
		}
		prefix, pkgName, found := strings.Cut(fileName, "_")
		if !found {
			continue
		}

		// cut out file extension to extract actual package name
		pkgName, _ = strings.CutSuffix(pkgName, ".jsonl")

		// grab file data
		fileBytes, err := os.ReadFile(fullPath)
		if err != nil {
			return mg.Fatalf(1, "error during os.ReadFile: \n%w", err)
		}
		path := filepath.Join(test.ResultsDir, prefix+".jsonl") // got.jsonl, want.jsonl, etc.

		// Add a header mentioning which package we're in to the start of the file
		contents := fmt.Sprintf("// %s\n%s", strings.ToUpper(pkgName), string(fileBytes))

		// create/truncate file if running first time; otherwise append with newline delimiter
		if count == 0 {
			if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
				return mg.Fatalf(1, "error during os.WriteFile: \n%w", err)
			}
		} else if err := utils.AppendFile(path, "\n"+contents); err != nil {
			return mg.Fatalf(1, "error during utils.AppendFile: \n%w", err)
		}

		count++
		// remove temp file after merging
		if err := sh.Rm(fullPath); err != nil {
			return mg.Fatalf(1, "error during sh.Rm: \n%w", err)
		}
	}

	var message string
	if count > 0 {
		message = fmt.Sprintf("Successfully merged a total of %d temp json files together at %q.", count, test.ResultsDir)
	} else {
		message = "No JSON files to merge were found."
	}
	logIfVerbose(message)
	return nil
}
