package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zerochae/gsbar/internal/sketchybar"
)

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload sketchybar configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := sketchybar.ReloadSketchybar(); err != nil {
			return fmt.Errorf("failed to reload sketchybar: %w", err)
		}
		fmt.Println("Sketchybar reloaded successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(reloadCmd)
}
