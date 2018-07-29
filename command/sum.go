package command

import "fmt"

type Sum struct {
	resultMap map[int]int
}

func NewSum() *Sum {
	return &Sum{map[int]int{}}
}


func (sum *Sum) Perform(targetIndex, targetVal int) {
	// sum
	if val, ok := sum.resultMap[targetIndex]; ok {
		sum.resultMap[targetIndex] = val + targetVal
	} else {
		sum.resultMap[targetIndex] = targetVal
	}
}

func (sum *Sum) Print() {
	// output
	for _, v := range sum.resultMap {
		fmt.Printf("%d%s", v, ",")
	}
	fmt.Println()
}
