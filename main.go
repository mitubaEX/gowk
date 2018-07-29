package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/mitubaEX/gowk/utils"
	"github.com/mitubaEX/gowk/service"
	"github.com/mitubaEX/gowk/command"
)

func help() {
	flag.Usage()
	os.Exit(2)
}

func main() {

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var (
		c = fs.String("c", "0", "target column number")
		d = fs.String("d", ",", "delimiter for line")
	)

	// -hオプション用文言
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage: %s <command> [<options>] <file>...
Options
`, os.Args[0])
		fs.PrintDefaults()
	}

	if len(os.Args) <= 1 || os.Args[1] == "-h" {
		help()
	}

	arg1 := os.Args[1]

	fs.Parse(os.Args[2:])
	//fmt.Println("arg1:", arg1)
	//fmt.Println("opt1:", *c)
	//fmt.Println("opt2:", *d)
	//fmt.Println("args:", fs.Args())

	opt := utils.NewOptions(*c, *d)

	switch arg1 {
	case "sum":
		service.Perform(command.NewSum(), opt)
	default:
		help()
	}
}
