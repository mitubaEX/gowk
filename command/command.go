package command

type Command interface {
	Perform([]string) error
	Print()
}
