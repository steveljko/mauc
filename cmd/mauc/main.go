package main

import (
	"fmt"
	"strings"

	"mauc/internal/utils"

	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "> ",
		HistoryFile:     "/tmp/mauc.tmp",
		AutoComplete:    nil,
		InterruptPrompt: "^C",
	})
	if err != nil {
		fmt.Println("Error creating readline ", err)
		return
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			fmt.Println(err)
			break
		}

		line = strings.TrimSpace(line)
		if line == "exit" {
			break
		}

		res, err := utils.Evaluate(line)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("=", res)
		}
	}
}
