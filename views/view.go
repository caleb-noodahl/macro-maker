package views

import (
	"fmt"
	"strings"

	"github.com/caleb-noodahl/macro-maker/commands"
	"github.com/jroimartin/gocui"
)

type View interface {
	Render(g *gocui.Gui) error
}

func WriteGreen(msg string, v *gocui.View) {
	v.Write([]byte(fmt.Sprintf("\033[32;4m%s\033[0m\n", msg)))
}

func WriteRed(msg string, v *gocui.View) {
	v.Write([]byte(fmt.Sprintf("\033[31;6m%s\033[0m\n", msg)))
}

func WriteYellow(msg string, v *gocui.View) {
	v.Write([]byte(fmt.Sprintf("\033[33;8m%s\033[0m\n", msg)))
}

func WriteWhite(msg string, v *gocui.View) {
	v.Write([]byte(fmt.Sprintf("%s\n", msg)))
}

func SetupCommandHandlers(g *gocui.Gui, cmds []commands.Command) error {
	g.DeleteKeybinding("input", gocui.KeyEnter, gocui.ModNone)
	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		input := v.Buffer()
		output, err := g.View("output")
		if err != nil {
			return err
		}
		split := strings.Split(input, " ")
		for _, cmd := range cmds {
			if cmd.Name() == strings.TrimSpace(split[0]) {
				res, err := cmd.Invoke(input)
				if err != nil {
					WriteRed(err.Error(), output)
				}
				WriteGreen(res, output)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
