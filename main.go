package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/mitubaEX/gowk/command"
	"github.com/mitubaEX/gowk/service"
	"github.com/mitubaEX/gowk/utils"
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

	// change input source
	var scanner *bufio.Scanner
	if fs.NArg() > 0 {
		file, err := os.Open(fs.Args()[0])
		if err != nil {
			panic(err)
		}
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	opt := utils.NewOptions(*c, *d, *f, *v)

	switch arg1 {
	case "sum":
		service.Perform(command.NewSum(opt), opt, scanner)
	case "filter":
		service.Perform(command.NewFilter(opt), opt, scanner)
	case "frequency":
		service.Perform(command.NewFrequency(opt), opt, scanner)
	case "max":
		service.Perform(command.NewMax(opt), opt, scanner)
	case "min":
		service.Perform(command.NewMin(opt), opt, scanner)
	case "length":
		service.Perform(command.NewLength(opt), opt, scanner)
	case "intersection":
		service.Perform(command.NewIntersection(opt), opt, scanner)
	case "distinct":
		service.Perform(command.NewDistinct(opt), opt, scanner)
	default:
		help()
	}
}
