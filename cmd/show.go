package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zerochae/gsbar/internal/sketchybar"
)

var showCmd = &cobra.Command{
	Use:   "show [item]",
	Short: "Show sketchybar item",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		item := args[0]

		switch item {
		case "config":
			if err := sketchybar.ShowConfigPopup(); err != nil {
				return fmt.Errorf("failed to show config popup: %w", err)
			}
			fmt.Println("Config popup shown")
		default:
			return fmt.Errorf("unknown item: %s", item)
		}

		return nil
	},
}

var hideCmd = &cobra.Command{
	Use:   "hide [item]",
	Short: "Hide sketchybar item",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		item := args[0]

		switch item {
		case "config":
			if err := sketchybar.HideConfigPopup(); err != nil {
				return fmt.Errorf("failed to hide config popup: %w", err)
			}
			fmt.Println("Config popup hidden")
		default:
			return fmt.Errorf("unknown item: %s", item)
		}

		return nil
	},
}

var toggleCmd = &cobra.Command{
	Use:   "toggle [item]",
	Short: "Toggle sketchybar item",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		item := args[0]

		switch item {
		case "config":
			if err := sketchybar.ToggleConfigPopup(); err != nil {
				return fmt.Errorf("failed to toggle config popup: %w", err)
			}
			fmt.Println("Config popup toggled")
		default:
			return fmt.Errorf("unknown item: %s", item)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(hideCmd)
	rootCmd.AddCommand(toggleCmd)
}
