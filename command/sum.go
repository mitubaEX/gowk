package command

import (
	"fmt"
	"github.com/mitubaEX/gowk/utils"
	"strings"
)

type Sum struct {
	resultMap map[int]int
}

func NewSum() *Sum {
	return &Sum{map[int]int{}}
}


func (sum *Sum) Perform(targetIndex int, targetVal string) {
	// sum
	targetValToInt := utils.StringToInt(targetVal)
	if val, ok := sum.resultMap[targetIndex]; ok {
		sum.resultMap[targetIndex] = val + targetValToInt
	} else {
		sum.resultMap[targetIndex] = targetValToInt
	}
}

func (sum *Sum) Print() {
	// output
	var printSlice []string
	for _, v := range sum.resultMap {
		printSlice = append(printSlice, utils.IntToString(v))
	}
	fmt.Println(strings.Join(printSlice, ","))
}
