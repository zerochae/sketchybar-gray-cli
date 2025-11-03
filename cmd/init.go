package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zerochae/gsbar/internal/config"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config.yaml with default values",
	Long:  `Creates ~/.config/gsbar/config.yaml with default configuration values.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewUser()
		if cfg == nil {
			fmt.Fprintln(os.Stderr, "Error: Failed to get config path")
			os.Exit(1)
		}

		defaults := map[string]string{
			"SBAR_FONT_FAMILY":        "SpaceMono Nerd Font Mono",
			"SBAR_ICON_FONT_SIZE":     "18.0",
			"SBAR_LABEL_FONT_SIZE":    "12.0",
			"SBAR_APP_ICON_FONT_SIZE": "13.5",
			"SBAR_CLOCK_FORMAT":       "MM/DD HH:mm",
			"SBAR_WEATHER_LOCATION":   "Seoul",
			"SBAR_NETSTAT_SHOW_GRAPH": "true",
			"SBAR_NETSTAT_SHOW_SPEED": "false",
		}

		for key, value := range defaults {
			cfg.Set(key, value)
		}

		if err := cfg.Save(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to save config: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("✓ Initialized ~/.config/sketchybar/user.sketchybarrc with default values")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
