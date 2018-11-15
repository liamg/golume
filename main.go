package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type output struct {
	id string
}

func main() {

	var change int
	var mute bool

	flag.BoolVar(&mute, "toggle-mute", mute, "Toggle mute on the active output")
	flag.IntVar(&change, "change-volume", change, "Change the volume by a given percent (negative to decrease volume) on the active output")
	flag.Parse()

	c, err := getActiveOutput()
	if err != nil {
		fmt.Printf("Error discovering active output: %s\n", err)
		os.Exit(1)
	}

	if !mute && change == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if mute {
		if err := c.ToggleMute(); err != nil {
			fmt.Printf("Error toggling mute on output #%s: %s\n", c.id, err)
		}
	}

	if change != 0 {
		if err := c.ChangeVolume(change); err != nil {
			fmt.Printf("Error adjusting volume on output #%s: %s\n", c.id, err)
		}
	}

}

func (c output) ToggleMute() error {
	return exec.Command("pactl", "set-sink-mute", c.id, "toggle").Run()
}

func (c output) ChangeVolume(changePct int) error {

	changeArg := fmt.Sprintf("%d%%", changePct)
	if changePct >= 0 {
		changeArg = fmt.Sprintf("+%s", changeArg)
	}

	return exec.Command("pactl", "set-sink-volume", c.id, changeArg).Run()
}

func getActiveOutput() (*output, error) {

	out, err := exec.Command("pactl", "list", "short").Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.HasSuffix(line, "RUNNING") {
			p := strings.Split(line, " ")
			return &output{id: p[0]}, nil
		}
	}

	return nil, fmt.Errorf("No active output found")
}
