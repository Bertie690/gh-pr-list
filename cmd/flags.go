package cmd

import (
	"log"

	color "github.com/mgutz/ansi"
)

func initFlags() {
	rootCmd.PersistentFlags().StringVarP(
		&cfgFile,
		"config",
		"c",
		"",
		// TODO: Add more schema variants
		`A path to a configuration file to use.
Currently supported formats: JSON.
For more information about configuration, run: `+color.Color("gh prlist help config", "blue"),
	)
	err := rootCmd.MarkPersistentFlagFilename("config", "yaml", "yml")
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.Flags().BoolP(
		"help",
		"h",
		false,
		"Show this help message",
	)

}