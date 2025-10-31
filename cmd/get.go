package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zerochae/gsbar/internal/config"
)

var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get configuration value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		value, err := config.GetValueCascade(key)
		if err != nil {
			return err
		}
		fmt.Println(value)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
