package command

import (
	"fmt"

	"github.com/mitubaEX/gowk/utils"
)

type Length struct {
	lengthList []int
	options    *utils.Options
}

func NewLength(options *utils.Options) *Length {
	return &Length{[]int{}, options}
}

func (length *Length) Perform(line []string) error {
	length.lengthList = append(length.lengthList, len(line))
	return nil
}

func (length *Length) Print() {
	for _, v := range length.lengthList {
		fmt.Println(v)
	}
}
