/*
Copyright © 2024 jake-young-dev
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "kelp", //kelp command has no functionality and will simply display the "help" message
	Short:   "A terminal remote console client for Minecraft servers",
	Long:    `Kelp is a terminal-friendly remote console client to manage Minecraft servers.`,
	Version: "v0.0.6",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//dont show completion subcommand in help message, makes the syntax confusing
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}
