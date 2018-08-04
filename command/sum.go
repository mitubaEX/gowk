package command

import (
	"fmt"
	"github.com/mitubaEX/gowk/utils"
	"strings"
	"sort"
)

type Sum struct {
	floatMap map[int]float64
	strMap map[int]string
	indexMap map[int]int
}

func NewSum() *Sum {
	return &Sum{map[int]float64{}, map[int]string{}, map[int]int{}}
}


func (sum *Sum) Perform(targetIndex int, targetVal string) {
	sum.indexMap[targetIndex] += 1
	// sum
	if utils.IsFloat(targetVal) {
		targetValToFloat, err := utils.StringToFloat(targetVal)
		if err != nil {
			panic(err)
		}
		if val, ok := sum.floatMap[targetIndex]; ok {
			sum.floatMap[targetIndex] = val + targetValToFloat
		} else {
			sum.floatMap[targetIndex] = targetValToFloat
		}
	} else {
		if val, ok := sum.strMap[targetIndex]; ok {
			sum.strMap[targetIndex] = val + targetVal
		} else {
			sum.strMap[targetIndex] = targetVal
		}
	}
}

func (sum *Sum) Print() {
	var keys []int
	for k := range sum.indexMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// output
	var printSlice []string
	for _, k := range keys {
		if _, ok := sum.floatMap[k] ; ok {
			printSlice = append(printSlice, utils.FloatToString(sum.floatMap[k]))
		}
		if _, ok := sum.strMap[k] ; ok {
			printSlice = append(printSlice, sum.strMap[k])
		}
	}
	fmt.Println(strings.Join(printSlice, ","))
}
