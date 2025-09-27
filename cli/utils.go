package cli

// import "github.com/spf13/cobra"

// func noArgsOrOneValidArg(cmd *cobra.Command, args []string) error {
// 	if len(args) == 0 {
// 		return nil
// 	}

// 	if err := cobra.ExactArgs(1)(cmd, args); err != nil {
// 		return err
// 	}

// 	return cobra.OnlyValidArgs(cmd, args)
// }

// silenceUsage toggles the `silenceUsage` parameter on rootCmd.
func silenceUsage(silent bool) {
	rootCmd.SilenceUsage = silent
}
