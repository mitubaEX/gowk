package command

import (
	"fmt"

	"github.com/mitubaEX/gowk/utils"
)

type Distinct struct {
	distinctMap map[int][]string
	options     *utils.Options
}

func NewDistinct(options *utils.Options) *Distinct {
	return &Distinct{map[int][]string{}, options}
}

func (distinct *Distinct) Perform(line []string) error {
	for _, v := range distinct.options.Column {
		targetIndex := v
		targetVal := line[v]
		distinct.distinctMap[targetIndex] = append(distinct.distinctMap[targetIndex], targetVal)
	}

	return nil
}

func (distinct *Distinct) Print() {
	printMap := map[string]int{}
	keys := make([]int, 0, len(distinct.distinctMap))
	for k := range distinct.distinctMap {
		keys = append(keys, k)
	}

	for _, v := range distinct.distinctMap[keys[0]] {
		if !utils.Contains(v, distinct.distinctMap[keys[1]]) {
			if _, ok := printMap[v]; !ok {
				printMap[v] = 1
			}
		}
	}

	for _, v := range distinct.distinctMap[keys[1]] {
		if !utils.Contains(v, distinct.distinctMap[keys[0]]) {
			if _, ok := printMap[v]; !ok {
				printMap[v] = 1
			}
		}
	}

	for k := range printMap {
		fmt.Println(k)
	}
}
