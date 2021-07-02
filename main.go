package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	hook "github.com/robotn/gohook"
)

type keytype int

const (
	keyboard   keytype = 0
	mouseclick         = 1
)

type keypress struct {
	KeyType keytype `json:"key_type"`
	Length  int64   `json:"length"`
	Code    uint16  `json:"code"`
	Key     rune    `json:"key"`
	X       int16   `json:"x"`
	Y       int16   `json:"y"`
}

type macro struct {
	last_time int64
	KPS       []keypress `json:"kps"`
}

func (m *macro) last() rune {
	if len(m.KPS) == 0 {
		return rune(' ')
	}
	return m.KPS[len(m.KPS)-1].Key
}

func timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (m *macro) push(k keypress) {
	switch k.KeyType {
	case keyboard:
		now := timestamp()
		k.Length = now - m.last_time
		m.KPS = append(m.KPS, k)
		m.last_time = timestamp()
	case mouseclick:
		m.KPS = append(m.KPS, k)
	}
	fmt.Printf("last: %+v\n", m.KPS[len(m.KPS)-1])
}

func (m *macro) save() {
	path := "C:/data/output.json"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	fmt.Printf("%+v\n", m)
	if err != nil {
		fmt.Println("could not write macro to file")
	}
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("could not marshel macro object")
	}
	file.Write(data)
	return
}

type hooked_event struct {
}

func main() {
	test()
}

func test() {
	hook.Register(hook.KeyDown, []string{"alt", "1"}, start)
	s := hook.Start()
	<-hook.Process(s)
}

func start(e hook.Event) {
	fmt.Printf("recording macro:\n")
	hook.End()
	m := macro{last_time: timestamp(), KPS: []keypress{}}
	echan := hook.Start()
	for ev := range echan {
		if ev.Keychar == 96 {
			fmt.Println("exit key pressed, saving macro")
			go m.save()
			hook.End()
			test()
			return
		}
		switch ev.Kind {
		case hook.KeyDown:
			m.push(keypress{Code: ev.Rawcode, Key: ev.Keychar, KeyType: keyboard})
		case hook.MouseDown:
			m.push(keypress{Code: 0, Key: ' ', X: ev.X, Y: ev.Y, KeyType: mouseclick})
		}
	}
}
