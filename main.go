package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type card struct {
	id string
}

func main() {

	var change int
	var mute bool

	flag.BoolVar(&mute, "toggle-mute", mute, "Mute all outputs")
	flag.IntVar(&change, "change-volume", change, "Change the volume by a given percent (negative to decrease volume) on all outputs")
	flag.Parse()

	cards, err := getCards()
	if err != nil {
		fmt.Printf("Error discovering sound cards: %s\n", err)
		os.Exit(1)
	}

	if !mute && change == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, c := range cards {
		if mute {
			if err := c.ToggleMute(); err != nil {
				fmt.Printf("Error toggling mute on card #%s: %s\n", c.id, err)
			}
		} else if change != 0 {
			if err := c.ChangeVolume(change); err != nil {
				fmt.Printf("Error adjusting volume on card #%s: %s\n", c.id, err)
			}
		}
	}

}

func (c card) ToggleMute() error {
	return exec.Command("pactl", "set-sink-mute", c.id, "toggle").Run()
}

func (c card) ChangeVolume(changePct int) error {

	changeArg := fmt.Sprintf("%d%%", changePct)
	if changePct >= 0 {
		changeArg = fmt.Sprintf("+%s", changeArg)
	}

	return exec.Command("pactl", "set-sink-volume", c.id, changeArg).Run()
}

func getCards() ([]card, error) {

	output, err := exec.Command("pactl", "list").Output()
	if err != nil {
		return nil, err
	}

	cards := []card{}

	var c *card
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Card #") {
			if c != nil {
				cards = append(cards, *c)
			}
			c = &card{
				id: line[6:],
			}
		}
	}

	if c != nil {
		cards = append(cards, *c)
	}

	return cards, nil

}
