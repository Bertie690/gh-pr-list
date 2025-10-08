package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/Bertie690/gh-pr-list/utils"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// is_CI reports whether the current process is running in CI (continuous integration)
// by checking the "CI" environment variable.
func is_CI() bool {
	return os.Getenv("CI") != ""
}

// Run gofumpt code quality checks.
func Lint() error {
	fmt.Println("Running gofumpt linting checks...")
	cmd := exec.Command("go", "tool", "gofumpt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func cleanTmpDir() error {
	if err := os.RemoveAll("tmp"); err != nil {
		return mg.Fatalf(1, "error cleaning out tmp dir: \n%w", err)
	}
	if err := os.Mkdir("tmp", 0755); err != nil {
		return mg.Fatalf(1, "error recreating tmp dir: \n%w", err)
	}
	return nil
}

// Run backend tests using gotestsum, passing args to "go test".
// CI runs will always run all tests across all packages,
// whereas non-CI runs can specify which package(s) to run as part of goTestArgs.
// If a package identifier is omitted on non-CI runs,
// it will default to running everything ("./...").
func Test(goTestArgs string) error {
	// NB: This _should_ take ...string[] as an argument, but mage doesn't support variadic arguments at the moment
	// a workaround is passing a "" as the first argument
	fmt.Println("Running backend tests...")
	mg.Deps(cleanTmpDir)

	// read gotestsum config args from text file;
	// use CI config if on CI; else regular config
	var filePath string
	if is_CI() {
		fmt.Println("CI run detected; using CI config")
		filePath = "gotestsum/gotestsum_ci.config.txt"
	} else {
		fmt.Println("Non-CI run detected; using default config")
		filePath = "gotestsum/gotestsum.config.txt"
	}

	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		return mg.Fatalf(1, "error reading gotestsum config file: \n%w", err)
	}

	// extract config values delimited by commas and whitespace
	configVals := strings.FieldsFunc(string(configBytes), func(r rune) bool {
		return (r == ',' || r == ' ' || r == '\r' || r == '\n')
	})
	fmt.Printf("Config file at %s successfully read.\n", filePath)

	// If the user forgot to add a package mark for non-CI runs,
	// do them a favor rather than outright failing.
	// CI runs are exempt from this due to rerun-fails requiring an explicit package argument
	// (not to mention their entire *job* is to test everything)
	args := strings.Fields(goTestArgs)
	if !is_CI() && slices.IndexFunc(args, func(s string) bool {
		return strings.HasPrefix(s, "./")
	}) == -1 {
		fmt.Println("No package identifier found; defaulting to running everything")
		args = append([]string{"./..."}, args...)
	}

	// tack on whatever config vals were passed by the user
	configVals = append(configVals, args...)

	// get package name, deferring to $GITHUB_REPOSITORY (workflow runs) if present
	pkgName := "gh-pr-list"
	if r := strings.TrimSpace(os.Getenv("GITHUB_REPOSITORY")); r != "" {
		pkgName = r
	}

	// merge together any temporary json files together once we're done testing.
	// We do this now to save time - if the prior steps fail,
	// there won't be any JSON files to merge)
	defer func() {
		if err := Merge_Temp_JSON(); err != nil {
			fmt.Printf("error merging temp JSON diffs after test run:\n%v\n", err)
		}
	}()

	return sh.RunWithV(map[string]string{"PKGNAME": pkgName},
		configVals[0], configVals[1:]...) // "go", "tool", "gotest.tools/gotestsum"...
}

// Remove all temp json files produced during tests and merge them together.
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
		if !strings.HasSuffix(fileName, ".jsonl") {
			continue
		}

		prefix, pkgName, found := strings.Cut(fileName, "_")
		if !found {
			// file name has no underscores, so it 100% isn't a preformatted JSON file
			continue
		}

		// cut out file extension to extract package name
		pkgName, _ = strings.CutSuffix(pkgName, ".jsonl")

		// grab file data
		fileBytes, err := os.ReadFile(fullName)
		if err != nil {
			return mg.Fatalf(1, "error during os.ReadFile: \n%w", err)
		}
		path := filepath.Join("tmp", prefix+".jsonl") // got.jsonl, want.jsonl, etc.

		// Add a header mentioning which package we're in to the start of the file
		contents := "//* " +
			strings.ToUpper(pkgName) + "\n" +
			string(fileBytes)
		if count == 0 {
			// truncate file if it already exists; otherwise add a newline delimiter
			if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
				return mg.Fatalf(1, "error during os.WriteFile: \n%w", err)
			}
		} else if err := utils.AppendFile(path, "\n"+contents); err != nil {
			return mg.Fatalf(1, "error during test.AppendFile: \n%w", err)
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
