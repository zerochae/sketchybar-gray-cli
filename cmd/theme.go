package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/zerochae/gsbar/internal/config"
)

var themeCmd = &cobra.Command{
	Use:   "theme [name]",
	Short: "Get or set the sketchybar theme",
	Long: `Get or set the sketchybar theme.

Available themes:
  - nord
  - tokyonight
  - ayudark
  - githubdark
  - onedark
  - onelight
  - ayulight
  - gruvboxlight
  - blossomlight
  - githublight

Examples:
  gsbar theme              # Show current theme
  gsbar theme nord         # Set theme to nord
  gsbar theme tokyonight   # Set theme to tokyonight`,
	Run: runTheme,
}

func init() {
	rootCmd.AddCommand(themeCmd)
}

func runTheme(cmd *cobra.Command, args []string) {
	cfg := config.NewUser()
	if cfg == nil {
		fmt.Fprintln(os.Stderr, "Failed to get user config path")
		os.Exit(1)
	}

	if err := cfg.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
		os.Exit(1)
	}

	if len(args) == 0 {
		currentTheme, ok := cfg.Get("SBAR_THEME")
		if !ok {
			currentTheme = "onedark (default)"
		}
		fmt.Printf("Current theme: %s\n", currentTheme)
		fmt.Println("\nAvailable themes:")
		fmt.Println("  Dark:  nord, tokyonight, ayudark, githubdark, onedark")
		fmt.Println("  Light: onelight, ayulight, gruvboxlight, blossomlight, githublight")
		return
	}

	themeName := args[0]

	validThemes := map[string]bool{
		"nord":          true,
		"tokyonight":    true,
		"ayudark":       true,
		"githubdark":    true,
		"onedark":       true,
		"onelight":      true,
		"ayulight":      true,
		"gruvboxlight":  true,
		"blossomlight":  true,
		"githublight":   true,
	}

	if !validThemes[themeName] {
		fmt.Fprintf(os.Stderr, "Invalid theme: %s\n", themeName)
		fmt.Println("Available themes:")
		fmt.Println("  Dark:  nord, tokyonight, ayudark, githubdark, onedark")
		fmt.Println("  Light: onelight, ayulight, gruvboxlight, blossomlight, githublight")
		os.Exit(1)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get home directory: %v\n", err)
		os.Exit(1)
	}

	themeFile := filepath.Join(home, ".config", "sketchybar", "tokens", "themes", themeName+".sh")
	if _, err := os.Stat(themeFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Theme file not found: %s\n", themeFile)
		os.Exit(1)
	}

	cfg.Set("SBAR_THEME", themeName)

	if err := cfg.Save(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Theme set to: %s\n", themeName)

	cmd2 := exec.Command("sketchybar", "--reload")
	if err := cmd2.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to reload sketchybar: %v\n", err)
	} else {
		fmt.Println("Sketchybar reloaded successfully")
	}
}
