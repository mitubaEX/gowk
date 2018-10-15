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
		line := []string{}
		for _, v := range strings.Split(scanner.Text(), options.Delimiter) {
			line = append(line, strings.TrimSpace(v))
		}

		if err := com.Perform(line); err != nil {
			log.Println(err)
		}
	}
	com.Print()
}
