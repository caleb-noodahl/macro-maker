package models

import (
	u "github.com/caleb-noodahl/macro-maker/utilities"
	"github.com/jroimartin/gocui"
	hook "github.com/robotn/gohook"
)

type Macro struct {
	ui        *gocui.Gui
	last_time int64
	KPS       []Keypress `json:"kps"`
}

func NewMacro() *Macro {
	return &Macro{
		last_time: u.Timestamp(),
		KPS:       []Keypress{},
	}
}

func (m *Macro) last() rune {
	if len(m.KPS) == 0 {
		return rune(' ')
	}
	return m.KPS[len(m.KPS)-1].Key
}

func (m *Macro) push(k Keypress) {
	switch k.KeyType {
	case Keyboard:
		now := u.Timestamp()
		k.Length = now - m.last_time
		m.KPS = append(m.KPS, k)
		m.last_time = u.Timestamp()
	case MouseClick:
		m.KPS = append(m.KPS, k)
	}
}

func (m *Macro) Record(e hook.Event) {
	hook.End()
	echan := hook.Start()
	for ev := range echan {
		if ev.Keychar == 96 {
			hook.End()
			return
		}
		switch ev.Kind {
		case hook.KeyDown:
			m.push(Keypress{Code: ev.Rawcode, Key: ev.Keychar, KeyType: Keyboard})
		case hook.MouseDown:
			m.push(Keypress{Code: 0, Key: ' ', X: ev.X, Y: ev.Y, KeyType: MouseClick})
		}
	}
}
