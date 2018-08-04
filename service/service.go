package service

import (
	"github.com/mitubaEX/gowk/command"
	"github.com/mitubaEX/gowk/utils"
	"bufio"
	"os"
	"strings"
)

func Perform(com command.Command, options *utils.Options) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), options.Delimiter)

		for _, v := range options.Column {
			com.Perform(v, line[v])
		}
	}
	com.Print()
}
