package main

import (
  "fmt"
)

func commandPokedex(cfg *config, args ...string) error {
  if len(cfg.pokedex) < 1 {
    fmt.Println("Your PokeDex is empty")
  } else {
    fmt.Println("Your PokeDex:")
    for _, pokemon := range cfg.pokedex {
      fmt.Printf(" - %s\n", pokemon.Name)
    }
  }
  return nil
}
