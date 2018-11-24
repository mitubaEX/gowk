package utils

import (
	"sort"
	"strings"
)

type Options struct {
	Column     []int
	Delimiter  string
	Conditions string
	IsVerbose  bool
}

func NewOptions(column string, delimiter string, conditions string, isVerbose bool) *Options {
	var columnSlice []int

	targetColumn := strings.Split(column, ",")

	for _, v := range targetColumn {
		val := StringToInt(v)
		columnSlice = append(columnSlice, val)
	}
	sort.Ints(columnSlice)
	return &Options{columnSlice, delimiter, conditions, isVerbose}
}
