// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package cli

import (
	"os"
	"runtime"
	"runtime/debug"

	"github.com/Bertie690/gh-pr-list/filter"
	"github.com/spf13/cobra"
)

// TODO: Find a way of auto-including version information rather than manually updating it each release

var (
	Version = "1.0.1"
	Commit  = ""
	Date    = "2025/09/30"
)

var rootCmd = &cobra.Command{
	// TODO: Make filter and template optional with default template matching `gh pr list`
	Use:   "gh pr-list [flags] filter template [-- ...args]",
	Short: "A gh extension providing a simple interface for listing active PRs.",
	Long: `A gh extension providing a simple interface for listing active PRs.

Any additional arguments after filter and template will be passed directly to ` + "`gh pr list`" + `.

For more information about JQ or Go template formatting, see ` + "`gh help formatting`.",
	Version:      buildVersion(Version, Commit, Date),
	Args:         cobra.MinimumNArgs(2),
	SilenceUsage: false,
}

// Execute executes the command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func buildVersion(version, commit, date string) string {
	result := version
	if commit != "" {
		result += "\nCommit: " + commit
	}
	if date != "" {
		result = "\nBuilt at: " + date
	}
	result += "\nGOOS: " + runtime.GOOS + "\nGOARCH: " + runtime.GOARCH
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result += "\nModule version: " + info.Main.Version + "\nChecksum: %s" + info.Main.Sum
	}
	return result
}

func init() {
	initFlags()
	rootCmd.SetVersionTemplate(`gh-pr-list {{printf "version %s\n" .Version}}`)
	rootCmd.RunE = runCmd
}

// runCmd runs the gh command.
func runCmd(cmd *cobra.Command, args []string) (err error) {
	silenceUsage(true)

	return filter.CreateList(args[0], args[1], args[2:])
}
