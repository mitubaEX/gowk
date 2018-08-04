package command

import (
	"go/types"
	"go/token"
	"github.com/mitubaEX/gowk/utils"
	"go/constant"
	"fmt"
	"strings"
)

type Filter struct {
	strMap map[int][]string
	options *utils.Options
}

func NewFilter(options *utils.Options) *Filter {
	return &Filter{map[int][]string{}, options}
}

func eval(expr string) (types.TypeAndValue, error) {
	return types.Eval(token.NewFileSet(), types.NewPackage("main", "main"), token.NoPos, expr)
}

func (filter *Filter) Perform(targetIndex int, targetVal string) {
	if utils.IsFloat(targetVal) {
		v, err := eval(targetVal + filter.options.Conditions)
		if err != nil {
			panic(err)
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
		v, err := eval(`"` + targetVal + `"` + filter.options.Conditions)
		if err != nil {
			panic(err)
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
