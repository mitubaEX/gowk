package utils

import (
	"strings"
	"sort"
)

type Options struct {
	Column []int
	Delimiter string
	Conditions string
}

func NewOptions (column string, delimiter string, conditions string) *Options {
	var columnSlice []int

	targetColumn := strings.Split(column, ",")

	for _, v := range targetColumn {
		val := StringToInt(v)
		columnSlice = append(columnSlice, val)
	}
	sort.Ints(columnSlice)
	return &Options{columnSlice, delimiter, conditions}
}
