package main

import (
  "strings"
  "bufio"
  "os"
  "fmt"
)

func StartRepl() {
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

    //fmt.Printf("Your command was: %s\n", words[0]) 

  }
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

