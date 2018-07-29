package utils

import "strings"

type Options struct {
	Column []string
	Delimiter string
}

func NewOptions (column string, delimiter string) *Options {
	targetColumn := strings.Split(column, ",")
	return &Options{targetColumn, delimiter}
}
