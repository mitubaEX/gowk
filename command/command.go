package command

type Command interface {
	Perform(targetIndex int, targetVal string)
	Print()
}
