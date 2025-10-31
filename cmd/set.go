package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zerochae/gsbar/internal/config"
	"github.com/zerochae/gsbar/internal/sketchybar"
)

var setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set configuration value",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		value := args[1]

		userCfg := config.NewUser()
		if userCfg == nil {
			return fmt.Errorf("failed to get user config path")
		}

		if err := userCfg.Load(); err != nil {
			return err
		}

		userCfg.Set(key, value)

		if err := userCfg.Save(); err != nil {
			return err
		}

		fmt.Printf("Set %s = %s\n", key, value)

		reload, _ := cmd.Flags().GetBool("reload")
		if reload {
			if err := sketchybar.ReloadSketchybar(); err != nil {
				return fmt.Errorf("failed to reload sketchybar: %w", err)
			}
			fmt.Println("Sketchybar reloaded")
		}

		return nil
	},
}

func init() {
	setCmd.Flags().BoolP("reload", "r", false, "Reload sketchybar after setting")
	rootCmd.AddCommand(setCmd)
}
