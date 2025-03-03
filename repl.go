package main

import (
  "strings"
  "bufio"
  "os"
  "fmt"
  "github.com/codybstrange/pokedexcli/internal/api"
)

type cliCommand struct {
  name string
  description string
  callback func(*config) error
}

type config struct {
  client api.Client
  nextLocationsURL *string
  prevLocationsURL *string
}

func getCommands() map[string]cliCommand {
  return map[string]cliCommand{
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
    "map": {
      name:         "map",
      description:  "Displays next 20 locations in Pokemon world",
      callback:     commandMapF,
    },
    "mapb": {
      name:         "mapb",
      description:  "Displays the last 20 locations in Pokemon world",
      callback:     commandMapB,
    },
    "explore": {
      name:         "explore",
      description:  "Display all the Pokemon located in a given location",
      callback:     commandExplore,
    },
  }
}


func startRepl(cfg *config) {
  reader := bufio.NewScanner(os.Stdin)
  commands := getCommands()
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
    
    if command, found := commands[commandName]; found {
      command.callback(cfg)
    }

  }
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

