package command

import (
	"errors"
	"fmt"
	"github.com/mitubaEX/gowk/utils"
	"log"
	"math"
	"strings"
)

type Min struct {
	maxMap  map[int]float64
	options *utils.Options
}

func NewMin(options *utils.Options) *Min {
	return &Min{map[int]float64{}, options}
}

func (min *Min) Perform(line []string) error {
	for _, v := range min.options.Column {
		targetIndex := v
		targetVal := line[v]
		if min.options.IsVerbose {
			log.Printf("current element is %v\n", targetVal)
		}

		if utils.IsFloat(targetVal) {
			targetValToFloat, err := utils.StringToFloat(targetVal)
			if err != nil {
				return err
			}

			if val, ok := min.maxMap[targetIndex]; ok {
				min.maxMap[targetIndex] = math.Min(val, targetValToFloat)
			} else {
				min.maxMap[targetIndex] = targetValToFloat
			}
		} else {
			return errors.New("string value is not available")
		}
	}
	return nil
}

func (min *Min) Print() {
	var printSlice []string
	for _, k := range min.options.Column {
		if _, ok := min.maxMap[k]; ok {
			printSlice = append(printSlice, utils.FloatToString(min.maxMap[k]))
		}
	}
	fmt.Println(strings.Join(printSlice, ","))
}
