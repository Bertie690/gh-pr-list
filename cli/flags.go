// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
