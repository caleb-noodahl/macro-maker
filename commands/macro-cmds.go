package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/caleb-noodahl/macro-maker/models"
	hook "github.com/robotn/gohook"
)

type MacroRecord struct {
	macro *models.Macro
	path  string
}

func (m *MacroRecord) Name() string {
	return "record"
}

func (m *MacroRecord) Invoke(input string) (string, error) {
	split := strings.Split(input, " ")
	m.macro = models.NewMacro()
	hook.Register(hook.KeyDown, []string{"alt", "1"}, m.macro.Record)
	s := hook.Start()
	<-hook.Process(s)
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("macro: %s\n", split[1]))
	for i, step := range m.macro.KPS {
		buf.WriteString(fmt.Sprintf("step: %v, press %c for %v ms\n", i, step.Key, step.Length))
	}
	return buf.String(), nil
}

func (m *MacroRecord) Save() error {
	file, err := os.OpenFile(m.path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	fmt.Printf("%+v\n", m)
	if err != nil {
		return err
	}
	data, err := json.Marshal(m.macro)
	if err != nil {
		return err
	}
	file.Write(data)
	return nil
}
