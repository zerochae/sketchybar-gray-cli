# sketchybar-gray-cli

Gray's Sketchybar Configuration Tool - CLI & TUI for managing sketchybar configurations

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

> 💡 This tool is designed for [sketchybar-gray](https://github.com/zerochae/sketchybar-gray) configuration management.

## ✨ Features

### CLI Mode
Fast command-line interface for scripting and automation:
- `gsbar get <key>` - Get configuration value
- `gsbar set <key> <value>` - Set configuration value
- `gsbar list` - List all configurations
- `gsbar reload` - Reload sketchybar

### TUI Mode
Beautiful interactive terminal UI powered by [Bubbletea](https://github.com/charmbracelet/bubbletea):
- Navigate with arrow keys or vim bindings (j/k)
- Visual configuration management
- Real-time feedback
- Custom color palette (39 curated colors)

## 🎨 Color Palette

Custom-designed color scheme with 39 colors organized in 4 categories:
- **Main:** Vibrant primary colors
- **Mild:** Soft, gentle tones
- **Pastel:** Light, subtle hues
- **Dim:** Dark, muted shades

## 🚀 Installation

### From Source

```bash
git clone https://github.com/zerochae/sketchybar-gray-cli.git
cd sketchybar-gray-cli
make build
```

The binary will be created at `bin/gsbar`

### Using Go Install

```bash
go install github.com/zerochae/gsbar@latest
```

## 📖 Usage

### CLI Mode

```bash
# Get a configuration value
gsbar get THEME

# Set a configuration value
gsbar set THEME dark

# Set and reload sketchybar
gsbar set FONT "SF Pro" --reload

# List all configurations
gsbar list

# Reload sketchybar
gsbar reload
```

### TUI Mode

```bash
# Launch interactive TUI (default when no args)
gsbar

# Or explicitly
gsbar tui
```

**TUI Controls:**
- `↑/↓` or `j/k` - Navigate
- `Enter` - Select/Confirm
- `Esc` - Go back
- `q` - Quit (from menu)
- `Ctrl+C` - Force quit

## 🏗️ Project Structure

```
sketchybar-gray-cli/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command (defaults to TUI)
│   ├── get.go             # Get command
│   ├── set.go             # Set command
│   ├── list.go            # List command
│   ├── reload.go          # Reload command
│   └── tui.go             # TUI command
├── tui/                   # TUI implementation
│   └── tui.go
├── internal/              # Internal packages
│   ├── colors/           # Color palette
│   ├── config/           # Configuration management
│   └── sketchybar/       # Sketchybar control
├── main.go               # Entry point
└── Makefile              # Build scripts
```

## 🔧 Development

### Build

```bash
make build       # Build to bin/gsbar
make run         # Build and run
make release     # Optimized release build
make clean       # Clean build artifacts
```

### Requirements

- Go 1.25+
- Sketchybar installed

## 📝 Configuration

gsbar manages configuration files at:
- User config: `~/.config/sketchybar/user.sketchybarrc`
- Default config: `~/.config/sketchybar/sketchybarrc`

Values are read in cascade: user config → default config

## 🔗 Related Projects

- [sketchybar-gray](https://github.com/zerochae/sketchybar-gray) - My sketchybar configuration
- [sketchybar](https://github.com/FelixKratz/SketchyBar) - The amazing macOS status bar

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details

## 🙏 Acknowledgments

- [Bubbletea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions
- [Cobra](https://github.com/spf13/cobra) - CLI framework

---

Built with ❤️ by [zerochae](https://github.com/zerochae)
