package command

import (
	"fmt"
	"github.com/mitubaEX/gowk/utils"
	"go/constant"
	"go/token"
	"go/types"
	"log"
	"strings"
)

type Filter struct {
	strSlice []string
	options  *utils.Options
}

func NewFilter(options *utils.Options) *Filter {
	return &Filter{[]string{}, options}
}

func eval(expr string) (types.TypeAndValue, error) {
	return types.Eval(token.NewFileSet(), types.NewPackage("main", "main"), token.NoPos, expr)
}

func (filter *Filter) Perform(line []string) error {
	for _, v := range filter.options.Column {
		targetVal := line[v]
		if filter.options.IsVerbose {
			log.Printf("current element is %v\n", targetVal)
		}

		// float val condition
		if utils.IsFloat(targetVal) {
			v, err := eval(targetVal + filter.options.Conditions)
			if err != nil {
				return err
			}

			if constant.BoolVal(v.Value) {
				filter.strSlice = append(filter.strSlice, strings.Join(line, ","))
			}
		} else {
			// string val condition
			v, err := eval(`"` + targetVal + `"` + filter.options.Conditions)
			if err != nil {
				return err
			}
			if constant.BoolVal(v.Value) {
				filter.strSlice = append(filter.strSlice, strings.Join(line, ","))
			}
		}
	}
	return nil
}

func (filter *Filter) Print() {
	for _, v := range filter.strSlice {
		fmt.Println(v)
	}
}
