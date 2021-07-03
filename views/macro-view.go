package views

import (
	c "github.com/caleb-noodahl/macro-maker/commands"
	m "github.com/caleb-noodahl/macro-maker/models"
	"github.com/jroimartin/gocui"
)

type MacroView struct {
	current *m.Macro
	cmds    []c.Command
}

func (m *MacroView) Render(g *gocui.Gui) error {
	m.cmds = []c.Command{
		&c.MacroRecord{},
	}

	if err := m.renderCMDsView(g, 0); err != nil {
		return err
	}
	if err := m.renderHelpText(g); err != nil {
		return err
	}
	SetupCommandHandlers(g, m.cmds)

	return nil
}

func (m *MacroView) renderHelpText(g *gocui.Gui) error {
	output, err := g.View("output")
	if err != nil {
		return err
	}
	WriteYellow("type record in the input below and hit enter (the UI will freeze)", output)
	WriteYellow("ex: record <name>\n", output)
	WriteYellow("while record is active, press alt + 1 to start recording the macro", output)
	WriteYellow("press ` to stop and save", output)
	return nil
}

func (m *MacroView) renderCMDsView(g *gocui.Gui, highlight int) error {
	view, err := g.View("cmds")
	if err != nil {
		return err
	}
	view.Clear()
	for i, cmd := range m.cmds {
		if i == highlight {
			WriteGreen(cmd.Name(), view)
		}
	}
	return nil
}
