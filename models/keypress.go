package models

type keytype int

const (
	Keyboard   keytype = 0
	MouseClick         = 1
)

type Keypress struct {
	KeyType keytype `json:"key_type"`
	Length  int64   `json:"length"`
	Code    uint16  `json:"code"`
	Key     rune    `json:"key"`
	X       int16   `json:"x"`
	Y       int16   `json:"y"`
}
