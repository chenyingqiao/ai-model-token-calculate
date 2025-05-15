package cmd

import "github.com/spf13/cobra"

func init() {
	// To ensure all subcommands get initialized
	// They each call rootCmd.AddCommand()
	// just by importing their package in main
	// In this simple setup, init in each file works fine
}

func RootCmd() *cobra.Command {
	return rootCmd
}
