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

func (max *Max) Perform(targetIndex int, targetVal string) error {
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
		return nil
	}

	return errors.New("string value is not available")
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
