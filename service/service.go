package service

import (
	"bufio"
	"log"
	"strings"

	"github.com/mitubaEX/gowk/command"
	"github.com/mitubaEX/gowk/utils"
)

func Perform(com command.Command, options *utils.Options, scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := []string{}

		splitFn := func(c rune) bool {

			s := []rune(options.Delimiter)

			var flag = false
			for _, v := range s {
				flag = c == v
			}
			return flag
		}

		for _, v := range strings.FieldsFunc(scanner.Text(), splitFn) {
			line = append(line, strings.TrimSpace(v))
		}

		if err := com.Perform(line); err != nil {
			log.Println(err)
		}
	}
	com.Print()
}
