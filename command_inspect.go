package main

import (
  "fmt"
  "errors"
)

func commandInspect(cfg *config, args ...string) error {
  if len(args) != 1 {
    return errors.New("you must provide a pokemon name")
  }

  name := args[0]
  if _, found := cfg.pokedex[name]; !found {
    fmt.Printf("you have not caught that pokemon\n")
    return nil
  }
  pokemon := cfg.pokedex[name]

  fmt.Printf("Name:   %s\n", pokemon.Name)
  fmt.Printf("Height: %v\n", pokemon.Height)
  fmt.Printf("Weight: %v\n", pokemon.Weight)
  fmt.Printf("Stats:\n")
  for _, stat := range pokemon.Stats {
    fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
  }
  fmt.Printf("Types:\n")
  for _, t := range pokemon.Types {
    fmt.Printf("  - %s\n", t.Type.Name)
  }
  return nil
}
