package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var (
		opt1 = fs.String("opt1", "default-value", "First string option")
		opt2 = fs.String("opt2", "default-value", "Second string option")
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
		flag.Usage()
		os.Exit(2)
	}

	arg1 := os.Args[1]
	switch arg1 {
	case "sum":
		fmt.Println("sum")
	}

	// flag.Parse()
	fs.Parse(os.Args[2:])
	fmt.Println("arg1:", arg1)
	fmt.Println("opt1:", *opt1)
	fmt.Println("opt2:", *opt2)
	fmt.Println("args:", fs.Args())
}
