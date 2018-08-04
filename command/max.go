package command

import (
	"github.com/mitubaEX/gowk/utils"
	"math"
	"fmt"
	"strings"
	"log"
)

type Max struct {
	maxMap map[int]float64
	options *utils.Options
}

func NewMax(options *utils.Options) *Max {
	return &Max{map[int]float64{}, options}
}

func (max *Max) Perform(targetIndex int, targetVal string) {
	if utils.IsFloat(targetVal) {
		targetValToFloat, err := utils.StringToFloat(targetVal)
		if err != nil {
			panic(err)
		}

		if val, ok := max.maxMap[targetIndex]; ok {
			max.maxMap[targetIndex] = math.Max(val, targetValToFloat)
		} else {
			max.maxMap[targetIndex] = targetValToFloat
		}
	} else {
		log.Fatalln("string value is not available")
	}
}

func (max *Max) Print() {
	var printSlice []string
	for _, k := range max.options.Column {
		if _, ok := max.maxMap[k] ; ok {
			printSlice = append(printSlice, utils.FloatToString(max.maxMap[k]))
		}
	}
	fmt.Println(strings.Join(printSlice, ","))
}
