package main

import (
  "strings"
  "bufio"
  "os"
  "fmt"
)

type cliCommand struct {
  name string
  description string
  callback func() error
}

var commandMap map[string]cliCommand

func commandExit() error {
  fmt.Printf("Closing the Pokedex... Goodbye!\n")
  os.Exit(0)
  return nil
}

func commandHelp() error {
  fmt.Printf("Welcome to the Pokedex!\n")
  fmt.Printf("Usage:\n\n")
  
  for commandKey, command := range commandMap {
    fmt.Printf("%v: %v\n", commandKey, command.description)
  }
  return nil
}

func initCommands() {
  commandMap = map[string]cliCommand{
    "help": {
      name:         "help",
      description:  "Displays a help message",
      callback:     commandHelp,
    },
    "exit": {
      name:         "exit",
      description:  "Exit the Pokedex",
      callback:     commandExit,
    },
  }
  return
}


func StartRepl() {
  initCommands()
  reader := bufio.NewScanner(os.Stdin)

  for ; ; {
    fmt.Printf("Pokedex > ")
    ok := reader.Scan()
    if !ok {
      fmt.Printf("Something went wrong with the scanning")
    }
    words := cleanInput(reader.Text())
    if len(words) == 0 {
      continue
    }
    
    commandName := words[0]
    
    //fmt.Printf("Your command was: %s\n", commandName)

    if command, found := commandMap[commandName]; found {
      command.callback()
    }

  }
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

