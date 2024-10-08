package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
  scanner := bufio.NewScanner(os.Stdin)
  for {
    fmt.Printf("Pokedex > ")

    scanner.Scan()
    text := scanner.Text()

    cleaned := cleanInput(text)
    if len(cleaned) == 0 {
      continue
    }

    comandName := cleaned[0]

    availableCommands := getCommands()

    command, ok := availableCommands[comandName]

    if !ok {
    fmt.Println("invalid command")
      continue
    }

    err := command.callback()
    if err != nil {
      fmt.Println(err)
    }
  }
}

type cliCommand struct {
  name		  string
  description	  string
  callback	  func() error
}

func getCommands() map[string]cliCommand {
  return map[string]cliCommand{
    "help": {
      name: "help",
      description: "prints the help menu",
      callback: callbackHelp,
    },
    "exit": {
	name: "exit",
	description: "turns off pokidex",
	callback: callbackExit,
    },
  }
}

func cleanInput(str string) []string {
  lowered := strings.ToLower(str)
  words := strings.Fields(lowered)
  return words
}
