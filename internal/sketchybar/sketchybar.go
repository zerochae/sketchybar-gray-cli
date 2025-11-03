package sketchybar

import (
	"os/exec"
)

func ReloadSketchybar() error {
	cmd := exec.Command("sketchybar", "--reload")
	return cmd.Run()
}

func ShowConfigPopup() error {
	cmd := exec.Command("sketchybar", "--set", "config", "drawing=on")
	return cmd.Run()
}

func HideConfigPopup() error {
	cmd := exec.Command("sketchybar", "--set", "config", "drawing=off")
	return cmd.Run()
}

func ToggleConfigPopup() error {
	cmd := exec.Command("sketchybar", "--set", "config", "drawing=toggle")
	return cmd.Run()
}
