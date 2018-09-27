package command

type Command interface {
	Perform(int, string) error
	Print()
}
