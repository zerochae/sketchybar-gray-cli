# gsbar

CLI tool for managing sketchybar configuration.

## Features

- ✅ Interactive TUI mode
- ✅ Simple bash config management
- ✅ Theme management (10 themes available)
- ✅ Sketchybar item control (show/hide/toggle)
- ✅ camelCase or SCREAMING_SNAKE_CASE support
- ✅ Works with or without gsbar installed

## Installation

```bash
git clone https://github.com/zerochae/gsbar
cd gsbar
make install
```

## Usage

### Interactive TUI Mode

```bash
gsbar        # Launch TUI mode (default)
gsbar tui    # Same as above
```

### Configuration Management

```bash
# Initialize config
gsbar init

# Get value
gsbar get fontFamily
gsbar get SBAR_FONT_FAMILY  # both work

# Set value
gsbar set fontFamily "SF Pro"
gsbar set SBAR_ICON_FONT_SIZE "20.0"

# List all
gsbar list
```

### Theme Management

```bash
# Show current theme and available themes
gsbar theme

# Set theme
gsbar theme nord
gsbar theme tokyonight
```

**Available Themes:**
- Dark: nord, tokyonight, ayudark, githubdark, onedark
- Light: onelight, ayulight, gruvboxlight, blossomlight, githublight

### Sketchybar Item Control

```bash
# Show item
gsbar show config

# Hide item
gsbar hide config

# Toggle item
gsbar toggle config
```

### Reload

```bash
# Reload sketchybar
gsbar reload
```

## Configuration

gsbar manages `~/.config/sketchybar/user.sketchybarrc`:

```bash
export SBAR_FONT_FAMILY="SpaceMono Nerd Font Mono"
export SBAR_ICON_FONT_SIZE="18.0"
export SBAR_LABEL_FONT_SIZE="12.0"
export SBAR_APP_ICON_FONT_SIZE="13.5"
export SBAR_CLOCK_FORMAT="MM/DD HH:mm"
export SBAR_WEATHER_LOCATION="Seoul"
export SBAR_NETSTAT_SHOW_GRAPH="true"
export SBAR_NETSTAT_SHOW_SPEED="false"
```

**You can also edit this file manually!**

## How It Works

```
gsbar
  ↓ manages
user.sketchybarrc (bash exports)
  ↓ sourced by
sketchybarrc
  ↓ launches
sketchybar
```

**Cascade**: `user.sketchybarrc` → `sketchybarrc` (defaults)

## Development

```bash
make build    # Build to bin/
make install  # Install to ~/.local/bin
make clean    # Clean artifacts
```

## License

MIT
