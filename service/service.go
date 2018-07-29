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
			targetIndex := utils.StringToInt(v)
			//targetVal := utils.StringToInt(line[targetIndex])

			com.Perform(targetIndex, line[targetIndex])
		}
	}
	com.Print()
}
