package commands

type Command interface {
	Invoke(input string) (string, error)
	Name() string
}
