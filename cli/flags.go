package cli

func initFlags() {
	rootCmd.Flags().BoolP(
		"help",
		"h",
		false,
		"Show this help message",
	)
	rootCmd.Flags().SortFlags = true
}
