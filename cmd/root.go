package cmd

import (
	"os"
	"runtime"
	"runtime/debug"

	"github.com/Bertie690/gh-pr-list/filter"
	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/cli/go-gh/v2/pkg/term"
	color "github.com/mgutz/ansi"
	"github.com/spf13/cobra"
)

var (
	Version = "0.0.0"
	Commit  = ""
	Date    = ""
	BuiltBy = ""
)

var (
	// The location of the current config file.
	cfgFile string
	// The command-line flag for the GitHub repository to be crawled.
	repoFlag string

	rootCmd = &cobra.Command{
		Use:     "gh pr-list [flags]",
		Short:   "A gh extension providing a simple interface for listing active PRs.",
		Version: buildVersion(Version, Commit, Date, BuiltBy),
	}
)

// Execute executes the command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}

func buildVersion(version, commit, date, builtBy string) string {
	result := version
	if commit != "" {
		result += "\nCommit: " + commit
	}
	if date != "" {
		result = "\nBuilt at: " + date
	}
	if builtBy != "" {
		result = "\nBuilt by: " + builtBy
	}
	result += "\nGOOS: " + runtime.GOOS + "\nGOARCH: " + runtime.GOARCH
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result += "\nModule version: " + info.Main.Version + "\nChecksum: %s" + info.Main.Sum
	}
	return result
}

func init() {
	// Disable colors if GH's color support says we shouldn't
	color.DisableColors(!term.FromEnv().IsColorEnabled())

	initFlags()
	rootCmd.Flags().SortFlags = true
	rootCmd.RunE = runCmd
}

// runCmd runs the gh command
func runCmd(_ *cobra.Command, args []string) (err error) {
	repo, err := getRepo()
	if err != nil {
		return
	}
	return filter.CreateList(repo, args)
}

// getRepo returns the GitHub repo to parse based on the CLI flag.
func getRepo() (repo repository.Repository, err error) {
	if repoFlag != "" {
		repo, err = repository.Parse(repoFlag)
	} else {
		repo, err = repository.Current()
	}
	return
}
