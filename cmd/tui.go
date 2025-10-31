package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zerochae/gsbar/tui"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch interactive TUI mode",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tui.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func RunTUI() {
	if err := tui.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
