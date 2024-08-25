package main

import (
  "fmt"
  "strings"

  "github.com/chzyer/readline"
)

func main() {
  rl, err := readline.New("> ")
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

    res, err := evaluate(line)
    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Println("=", res)
  }
  }
}
