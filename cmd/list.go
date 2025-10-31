package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/zerochae/gsbar/internal/config"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configuration values",
	RunE: func(cmd *cobra.Command, args []string) error {
		userCfg := config.NewUser()
		if userCfg == nil {
			return fmt.Errorf("failed to get user config path")
		}

		if err := userCfg.Load(); err != nil {
			return err
		}

		configs := userCfg.List()

		if len(configs) == 0 {
			fmt.Println("No user configuration found")
			return nil
		}

		keys := make([]string, 0, len(configs))
		for k := range configs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Printf("%s=%s\n", k, configs[k])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
