package views

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
)

type HomeView struct {
	ViewList map[string]View
}

func (h *HomeView) Render(g *gocui.Gui) error {
	g.SetManagerFunc(h.layout)
	g.Cursor = true
	if err := h.initKeybindings(g); err != nil {
		log.Fatalln(err)
	}
	return nil
}

func (h *HomeView) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("cmds", maxX-32, 0, maxX-3, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "cmds"
		for key := range h.ViewList {
			if key == "home" {
				WriteGreen(key, v)
			} else {
				WriteWhite(key, v)
			}

		}
	}

	if v, err := g.SetView("output", 0, 0, maxX-34, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Autoscroll = true
		v.Editable = false
		v.Wrap = true
	}

	if v, err := g.SetView("input", 0, maxY-3, maxX-34, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView("input"); err != nil {
			return err
		}
		v.Editable = true
		v.Wrap = true
	}

	return nil
}

func (h *HomeView) initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlA, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			v.Mask ^= '*'
			return nil
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			//append intput to output
			input := v.Buffer()
			outputView, err := g.View("output")
			if err != nil {
				return err
			}
			split := strings.Split(input, " ")
			if view, ok := h.ViewList[strings.TrimSpace(split[0])]; ok {
				WriteGreen(fmt.Sprintf("> %s", input), outputView)
				if err := view.Render(g); err != nil {
					WriteRed(err.Error(), outputView)
				}
				//after we're done with that view render the home view
				//h.Render(g)
			} else {
				WriteRed(fmt.Sprintf("unknown view: %s", split[0]), outputView)
			}
			v.Clear()
			v.SetCursor(0, 0)
			return nil
		}); err != nil {
		return err
	}
	return nil
}
