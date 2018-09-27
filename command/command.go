package command

type Command interface {
	Perform(targetIndex int, targetVal string) error
	Print()
}
