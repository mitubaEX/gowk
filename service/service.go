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

		if err := com.Perform(line); err != nil {
			log.Println(err)
		}
	}
	com.Print()
}
