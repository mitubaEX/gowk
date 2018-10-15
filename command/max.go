package command

import (
	"errors"
	"fmt"
	"github.com/mitubaEX/gowk/utils"
	"math"
	"strings"
)

type Max struct {
	maxMap  map[int]float64
	options *utils.Options
}

func NewMax(options *utils.Options) *Max {
	return &Max{map[int]float64{}, options}
}

func (max *Max) Perform(line []string) error {
	for _, v := range max.options.Column {
		targetIndex := v
		targetVal := line[v]

		if utils.IsFloat(targetVal) {
			targetValToFloat, err := utils.StringToFloat(targetVal)
			if err != nil {
				return err
			}

			if val, ok := max.maxMap[targetIndex]; ok {
				max.maxMap[targetIndex] = math.Max(val, targetValToFloat)
			} else {
				max.maxMap[targetIndex] = targetValToFloat
			}
		} else {
			return errors.New("string value is not available")
		}
	}
	return nil
}

func (max *Max) Print() {
	var printSlice []string
	for _, k := range max.options.Column {
		if _, ok := max.maxMap[k]; ok {
			printSlice = append(printSlice, utils.FloatToString(max.maxMap[k]))
		}
	}
	fmt.Println(strings.Join(printSlice, ","))
}
