package sketchybar

import (
	"os/exec"
)

func ReloadSketchybar() error {
	cmd := exec.Command("sketchybar", "--reload")
	return cmd.Run()
}
