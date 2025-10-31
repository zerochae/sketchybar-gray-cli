package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gsbar",
	Short: "Gray's Sketchybar Configuration Tool",
	Long: `┌─┐┬─┐┌─┐┬ ┬   ┌─┐┬┌─┌─┐┌┬┐┌─┐┬ ┬┬ ┬   ┌┐ ┌─┐┬─┐
│ ┬├┬┘├─┤└┬┘───└─┐├┴┐├┤  │ │  ├─┤└┬┘───├┴┐├─┤├┬┘
└─┘┴└─┴ ┴ ┴    └─┘┴ ┴└─┘ ┴ └─┘┴ ┴ ┴    └─┘┴ ┴┴└─

gsbar is a CLI tool to manage your sketchybar configuration.`,
	Version: "0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		RunTUI()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
