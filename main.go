package main

import (
  "fmt"
  "bufio"
  "strings"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  for ; ; {
    fmt.Printf("Pokedex > ")
    ok := scanner.Scan()
    if !ok {
      fmt.Printf("Something went wrong with the scanning")
    }
    text := scanner.Text()
    output := strings.ToLower(text)
	  words := strings.Fields(output)
	  fmt.Printf("Your command was: %s\n", words[0]) 
  }
}


