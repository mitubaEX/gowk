package command

import (
	"github.com/mitubaEX/gowk/utils"
	"fmt"
)

type Frequency struct {
	strMap map[int]map[string]int
	options *utils.Options
}

func NewFrequency(options *utils.Options) *Frequency {
	return &Frequency{map[int]map[string]int{}, options}
}

func (frequency *Frequency) Perform(targetIndex int, targetVal string) error {
	if _, ok := frequency.strMap[targetIndex][targetVal]; ok {
		frequency.strMap[targetIndex][targetVal] += 1
	} else {
		val := frequency.strMap[targetIndex]
		if len(val) == 0 {
			frequency.strMap[targetIndex] = map[string]int{targetVal: 1}
		} else {
			val[targetVal] = 1
			frequency.strMap[targetIndex] = val
		}
	}
	return nil
}

func (frequency *Frequency) Print() {
	for _, k := range frequency.options.Column {
		for key, val := range frequency.strMap[k] {
			fmt.Printf("%s: %d\n", key, val)
		}
	}
}
