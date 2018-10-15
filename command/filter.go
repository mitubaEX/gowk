package command

import (
	"fmt"
	"github.com/mitubaEX/gowk/utils"
	"go/constant"
	"go/token"
	"go/types"
	"strings"
)

type Filter struct {
	strMap  map[int][]string
	options *utils.Options
}

func NewFilter(options *utils.Options) *Filter {
	return &Filter{map[int][]string{}, options}
}

func eval(expr string) (types.TypeAndValue, error) {
	return types.Eval(token.NewFileSet(), types.NewPackage("main", "main"), token.NoPos, expr)
}

func (filter *Filter) Perform(line []string) error {
	for _, v := range filter.options.Column {
		targetIndex := v
		targetVal := line[v]

		// float val condition
		if utils.IsFloat(targetVal) {
			v, err := eval(targetVal + filter.options.Conditions)
			if err != nil {
				return err
			}

			if constant.BoolVal(v.Value) {
				if val, ok := filter.strMap[targetIndex]; ok {
					val = append(val, targetVal)
					filter.strMap[targetIndex] = val
				} else {
					filter.strMap[targetIndex] = []string{targetVal}
				}
			}
		} else {
			// string val condition
			v, err := eval(`"` + targetVal + `"` + filter.options.Conditions)
			if err != nil {
				return err
			}
			if constant.BoolVal(v.Value) {
				if val, ok := filter.strMap[targetIndex]; ok {
					val = append(val, targetVal)
					filter.strMap[targetIndex] = val
				} else {
					filter.strMap[targetIndex] = []string{targetVal}
				}
			}
		}
	}
	return nil
}

func (filter *Filter) Print() {
	var printSlice []string

	for i := 0; i < len(filter.strMap[filter.options.Column[0]]); i++ {
		printSlice = []string{}
		for _, k := range filter.options.Column {
			if i < len(filter.strMap[k]) {
				printSlice = append(printSlice, filter.strMap[k][i])
			} else {
				printSlice = append(printSlice, "0")
			}
		}
		fmt.Println(strings.Join(printSlice, ","))
	}

}
