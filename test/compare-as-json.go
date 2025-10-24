// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

// Package test contains some useful utility functions for testing.
//
// It should not be used in production.
package test

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Bertie690/gh-pr-list/utils"
	"github.com/nsf/jsondiff"
)

// CompareAsJSON compares 2 objects as JSON outputs for testing.
//
// If the comparison fails, this marks the current test as a failure and
// writes 3 JSONL files to the [ResultsDir] directory:
//   - `got.jsonl` contains serialized versions of `got`
//   - `want.jsonl` contains serialized versions of `want`
//   - `diff.jsonl` contains a pretty-printed difference between the two JSON objects
//     (courtesy of [github.com/nsf/jsondiff]).
//
// This json difference is passed to [testing.T.Errorf] as well for ease of use.
//
// These files are continuously appended to during a test run (sectioned off by test name),
// and must be moved or removed after the package finishes testing via [TestMain] or similar.
// Invocation from parallel tests is untested and not recommended.
//
// A failure to parse JSON will halt test execution and fail immediately.
//
// [TestMain]: https://pkg.go.dev/testing#hdr-Main
// [ResultsDir]: https://pkg.go.dev/github.com/Bertie690/gh-pr-list/test#ResultsDir
func CompareAsJSON(t *testing.T, got, want any) {
	if got == nil && want == nil {
		return
	} else if (got == nil) != (want == nil) { // one is nil and the other isn't
		t.Helper()
		t.Errorf("Unequal values (nilness): got = %v, want = %v", got, want)
	}

	var (
		gotJson, wantJson string
		ok                bool
	)

	// Don't re-serialize objects that are already valid JSON
	if gotJson, ok = got.(string); !ok {
		gotBytes, err := json.MarshalIndent(got, "", "\t")
		if err != nil {
			t.Fatalf("CompareAsJSON could not marshal got to json: \n%v", err)
		}
		gotJson = string(gotBytes)
	}
	if wantJson, ok = want.(string); !ok {
		wantBytes, err := json.MarshalIndent(want, "", "\t")
		if err != nil {
			t.Fatalf("CompareAsJSON could not marshal want to json: \n%v", err)
		}
		wantJson = string(wantBytes)
	}

	// Ignore whitespace while making the comparison
	if utils.RemoveWhitespace(gotJson) == utils.RemoveWhitespace(wantJson) {
		return
	}
	diff, err := parseJSONDiff(gotJson, string(wantJson), t.Name())
	if err != nil {
		t.Fatalf("error creating JSON diffs: \n%v", err)
	}

	// Remove block comments in the stdout version since we value clutter-free output over valid syntax
	r := strings.NewReplacer("/* ", "", " */", ":")

	t.Helper()
	t.Errorf("JSONs not equal; diff between got & want: \n%s", r.Replace(diff))
}

// Parsing options for jsondiff.
// Fun fact: this is guaranteed to produce valid JSONL output in the diff
// so long as the input values are also valid (which should always be the case).
var options = jsondiff.Options{
	Added:            jsondiff.Tag{Begin: "/* Added */ ", End: ""},
	Removed:          jsondiff.Tag{Begin: "/* Removed */ ", End: ""},
	Changed:          jsondiff.Tag{Begin: "/* Changed */ [ ", End: " ]"},
	ChangedSeparator: ", ",
	Indent:           "\t",
	SkipMatches:      true,
}

// parseJSONDiff parses and diffs a pair of JSON outputs, creating files to log values as appropriate.
// It returns the human-readable diff output, as well as any error produced.
//
// A non-nil error is always accompanied by an empty diff string.
func parseJSONDiff(gotJSON, wantJSON, testName string) (diff string, err error) {
	_, diff = jsondiff.Compare([]byte(gotJSON), []byte(wantJSON), &options)
	header := "// " + testName + "\n"

	if err := os.MkdirAll(ResultsDir, 0o755); err != nil {
		return "", err
	}
	for i := range 3 {
		var file, body string
		switch i {
		case 0:
			file = "got.jsonl"
			body = gotJSON
		case 1:
			file = "want.jsonl"
			body = wantJSON
		case 2:
			file = "diff.jsonl"
			body = diff
		}
		path := filepath.Join(ResultsDir, file)

		if _, err = os.Stat(path); err == nil {
			// add extra newline in header to properly delimit sections on existing files
			header = "\n" + header
		} else if !errors.Is(err, os.ErrNotExist) {
			return "", err
		}

		if err = utils.AppendFile(path, header+body+"\n"); err != nil {
			return "", err
		}
	}

	return diff, nil
}
