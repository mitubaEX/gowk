package main

import (
	"flag"
	"fmt"
	"github.com/mitubaEX/gowk/command"
	"github.com/mitubaEX/gowk/service"
	"github.com/mitubaEX/gowk/utils"
	"os"
)

func help() {
	flag.Usage()
	os.Exit(2)
}

func main() {

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var (
		c = fs.String("c", "0", "target column number")
		d = fs.String("d", ",", "delimiter for line (delimiter should be rune of golang)")
		f = fs.String("f", "> 0", "function of filter condition {required at filter mode}")
		v = fs.Bool("v", false, "set verbose mode")
	)

	// -hオプション用文言
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage: %s <command> [<options>] <file>...`, os.Args[0])
		fmt.Fprintf(os.Stderr, `
Commands:
	sum
	filter
	frequency
	min
	max
	length
	intersection
	distinct
`)
		fmt.Fprintf(os.Stderr, `
Options:
`)
		fs.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Documentation: https://github.com/mitubaEX/gowk
`)
	}

	if len(os.Args) <= 1 || os.Args[1] == "-h" {
		help()
	}

	arg1 := os.Args[1]

	fs.Parse(os.Args[2:])

	opt := utils.NewOptions(*c, *d, *f, *v)

	switch arg1 {
	case "sum":
		service.Perform(command.NewSum(opt), opt)
	case "filter":
		service.Perform(command.NewFilter(opt), opt)
	case "frequency":
		service.Perform(command.NewFrequency(opt), opt)
	case "max":
		service.Perform(command.NewMax(opt), opt)
	case "min":
		service.Perform(command.NewMin(opt), opt)
	case "length":
		service.Perform(command.NewLength(opt), opt)
	case "intersection":
		service.Perform(command.NewIntersection(opt), opt)
	case "distinct":
		service.Perform(command.NewDistinct(opt), opt)
	default:
		help()
	}
}
