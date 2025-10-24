// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package cli

import (
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/Bertie690/gh-pr-list/filter"
	"github.com/spf13/cobra"
)

// Version information, injected during `release.yml` via `-ldflags -X`
var (
	version   = "0.0.0-develop"
	commit    = "local"
)

var rootCmd = &cobra.Command{
	// TODO: Make filter and template optional with default template matching `gh pr list`
	Use:   "gh pr-list [flags] filter template [-- ...args]",
	Short: "A gh extension providing a simple interface for listing active PRs.",
	Long: `A gh extension providing a simple interface for listing active PRs.

Any additional arguments after filter and template will be passed directly to ` + "`gh pr list`" + `.

For more information about JQ or Go template formatting, see ` + "`gh help formatting`.",
	Args:         cobra.MinimumNArgs(2),
	SilenceUsage: false,
}

// Execute executes the command.
func Execute() {
	rootCmd.Version = versionText(version, commit)
	rootCmd.SetVersionTemplate(`gh-pr-list {{printf "%s\n" .Version}}`)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// versionText gets the versioning text from the version, commit and build time.
func versionText(version, commit string) (result string) {
	if versionWithoutV, _ := strings.CutPrefix(version, "v"); versionWithoutV != "" {
		result += "\nVersion: " + versionWithoutV
	}
	if commit != "" {
		result += "\nCommit: " + commit
	}
	result += "\nGOOS: " + runtime.GOOS + "\nGOARCH: " + runtime.GOARCH
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result += "\nModule version: " + info.Main.Version + "\nChecksum: %s" + info.Main.Sum
	}
	return result
}

func init() {
	initFlags()
	rootCmd.RunE = runCmd
}

// runCmd runs the gh command.
func runCmd(cmd *cobra.Command, args []string) (err error) {
	silenceUsage(true)

	return filter.CreateList(args[0], args[1], args[2:])
}
