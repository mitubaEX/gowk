package command

import (
	"fmt"
	"github.com/mitubaEX/gowk/utils"
	"log"
	"strings"
)

type Sum struct {
	floatMap map[int]float64
	strMap   map[int]string
	options  *utils.Options
}

func NewSum(options *utils.Options) *Sum {
	return &Sum{map[int]float64{}, map[int]string{}, options}
}

func (sum *Sum) Perform(line []string) error {
	for _, v := range sum.options.Column {
		targetIndex := v
		targetVal := line[v]
		if sum.options.IsVerbose {
			log.Printf("current element is %v\n", targetVal)
		}

		// float val condition
		if utils.IsFloat(targetVal) {
			targetValToFloat, err := utils.StringToFloat(targetVal)
			if err != nil {
				return err
			}

			if val, ok := sum.floatMap[targetIndex]; ok {
				sum.floatMap[targetIndex] = val + targetValToFloat
			} else {
				sum.floatMap[targetIndex] = targetValToFloat
			}
		} else {
			// string val condition
			if val, ok := sum.strMap[targetIndex]; ok {
				sum.strMap[targetIndex] = val + targetVal
			} else {
				sum.strMap[targetIndex] = targetVal
			}
		}
	}

	return nil
}

func (sum *Sum) Print() {
	// output
	var printSlice []string
	for _, k := range sum.options.Column {
		if _, ok := sum.floatMap[k]; ok {
			printSlice = append(printSlice, utils.FloatToString(sum.floatMap[k]))
		}
		if _, ok := sum.strMap[k]; ok {
			printSlice = append(printSlice, sum.strMap[k])
		}
	}
	fmt.Println(strings.Join(printSlice, ","))
}
