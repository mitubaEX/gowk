package command

type Command interface {
	Perform(targetIndex, targetVal int)
	Print()
}
