package command

import (
	"fmt"

	"github.com/mitubaEX/gowk/utils"
)

type Intersection struct {
	intersectionMap map[int][]string
	options         *utils.Options
}

func NewIntersection(options *utils.Options) *Intersection {
	return &Intersection{map[int][]string{}, options}
}

func (intersection *Intersection) Perform(line []string) error {
	for _, v := range intersection.options.Column {
		targetIndex := v
		targetVal := line[v]
		intersection.intersectionMap[targetIndex] = append(intersection.intersectionMap[targetIndex], targetVal)
	}

	return nil
}

func (intersection *Intersection) Print() {
	printMap := map[string]int{}
	keys := make([]int, 0, len(intersection.intersectionMap))
	for k := range intersection.intersectionMap {
		keys = append(keys, k)
	}

	for _, v := range intersection.intersectionMap[keys[0]] {
		if utils.Contains(v, intersection.intersectionMap[keys[1]]) {
			if _, ok := printMap[v]; !ok {
				printMap[v] = 1
			}
		}
	}

	for k := range printMap {
		fmt.Println(k)
	}
}
