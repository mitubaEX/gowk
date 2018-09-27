package service

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/mitubaEX/gowk/command"
	"github.com/mitubaEX/gowk/utils"
)

func Perform(com command.Command, options *utils.Options) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), options.Delimiter)

		for _, v := range options.Column {
			if err := com.Perform(v, line[v]); err != nil {
				log.Println(err)
			}
		}
	}
	com.Print()
}
