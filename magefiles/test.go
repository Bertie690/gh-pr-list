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

	"github.com/Bertie690/gh-pr-list/utils"
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// isCI reports whether the current process is running in CI (continuous integration)
// by checking the "CI" environment variable.
func isCI() bool {
	return os.Getenv("CI") != ""
}

// Run golangci-lint code quality checks.
func Lint() error {
	path, err := exec.LookPath("golangci-lint")
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("golangci-lint executable not found in PATH")
	} else if err != nil {
		return fmt.Errorf("Unknown error found when retrieving golangci-lint path: %w", err)
	}

	fmt.Println("Found golangci-lint binary at", path)

	cmd := exec.Command("golangci-lint", "run")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func cleanTmpDir() error {
	if err := os.RemoveAll("tmp"); err != nil {
		return mg.Fatalf(1, "error cleaning out tmp dir: \n%w", err)
	}
	if err := os.Mkdir("tmp", 0o755); err != nil {
		return mg.Fatalf(1, "error recreating tmp dir: \n%w", err)
	}
	return nil
}

// Run backend tests using gotestsum, passing any provided args to "go test".
// CI runs will always run all tests across all packages,
// whereas non-CI runs can specify which package(s) to run as part of goTestArgs.
// If a package identifier is omitted on non-CI runs,
// it will default to running everything ("./...").
func Test(goTestArgs string) error {
	// NB: This _should_ take ...string as an argument, but mage doesn't support variadic arguments at the moment.
	// A workaround is passing a "" as the first argument

	mg.Deps(cleanTmpDir)

	// Change args if on CI (different reporter for gh actions, etc.)
	var args []string
	if isCI() {
		fmt.Println(utils.ColorHex("CI run detected; using CI config", "#ffa500"))
		args = testArgsCI
	} else {
		fmt.Println(utils.ColorHex("Non-CI run detected; using normal config", "#ffa500"))
		args = testArgsLocal
	}

	// If the user forgot to add a package mark for non-CI runs,
	// do them a favor rather than outright failing.
	// CI runs are exempt from this due to rerun-fails requiring an explicit package argument
	// (not to mention their entire *job* is to test everything)
	cliArgs := strings.Fields(goTestArgs)
	if !isCI() && slices.IndexFunc(cliArgs, func(s string) bool {
		return strings.HasPrefix(s, "./")
	}) == -1 {
		fmt.Println(color.BlueString("No package identifier found; defaulting to running everything"))
		cliArgs = append(cliArgs, "./...")
	}

	// tack on whatever config vals were passed by the user, separated by a double dash
	args = append(args, cliArgs...)

	// Get package name for JUnit XML reports, deferring to $GITHUB_REPOSITORY env var if present
	pkgName := "gh-pr-list"
	if r := strings.TrimSpace(os.Getenv("GITHUB_REPOSITORY")); r != "" {
		pkgName = r
	}

	// Merge together any temporary json files together once we're done testing.
	// We do this now to save time - if the prior steps fail,
	// there won't be any JSON files to merge
	defer func() {
		if err := Merge_Temp_JSON(); err != nil {
			fmt.Printf("error merging temp JSON diffs after test run:\n%v\n", err)
		}
	}()

	return sh.RunWithV(map[string]string{"PKGNAME": pkgName},
		"go", args...) // "go", "tool", "gotest.tools/gotestsum", ...
}

// Remove all temporary json files produced during tests and merge them together.
// This takes all files matching the format "XXX_**.jsonl",
// and merges them together into a single file named "XXX.jsonl".
// Comments are added between failing tests from different packages.
func Merge_Temp_JSON() error {
	tmpdir, err := os.Open("tmp")
	if err != nil {
		return mg.Fatalf(1, "error while opening temp folder: \n%w", err)
	}
	fileNames, err := tmpdir.Readdirnames(-1)
	if err != nil {
		return mg.Fatalf(1, "error while reading temp folder files: \n%w", err)
	}

	if len(fileNames) == 0 {
		fmt.Println("No files were found inside ./tmp, exiting")
		return nil
	}

	count := 0
	for _, fileName := range fileNames {
		fullName := filepath.Join("tmp", fileName)
		// skip non JSONL files or ones without underscores
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
		fileBytes, err := os.ReadFile(fullName)
		if err != nil {
			return mg.Fatalf(1, "error during os.ReadFile: \n%w", err)
		}
		path := filepath.Join("tmp", prefix+".jsonl") // got.jsonl, want.jsonl, etc.

		// Add a header mentioning which package we're in to the start of the file
		contents := fmt.Sprintf("// %s\n%s", strings.ToUpper(pkgName), string(fileBytes))
		if count == 0 {
			// truncate file if it already exists; otherwise add a newline delimiter
			if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
				return mg.Fatalf(1, "error truncating file contents: \n%w", err)
			}
		} else if err := utils.AppendFile(path, "\n"+contents); err != nil {
			return mg.Fatalf(1, "error during utils.AppendFile: \n%w", err)
		}

		count++
		// remove test file after merging
		if err := sh.Rm(fullName); err != nil {
			return err
		}
	}

	var message string
	if count > 0 {
		message = fmt.Sprintf("Successfully merged a total of %d temp json files together.", count)
	} else {
		message = "No JSON files to merge were found."
	}
	fmt.Println(message, "\nHave a nice day.")
	return nil
}
