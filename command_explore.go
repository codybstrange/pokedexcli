package main

import (
  "fmt"
  "errors"
)

func commandExplore(cfg *config, args ...string) error {
  if len(args) != 1 {
    return errors.New("you must provide a location name")
  }
  name := args[0]
  locationResp, err := cfg.client.GetLocation(name)

  if err != nil {
    return err
  }

  fmt.Printf("Exploring %s...\n", locationResp.Name)
  
  for i, enc := range locationResp.PokemonEncounters {
    if i == 0 {
      fmt.Println("Found Pokemon:")
    }
    fmt.Printf(" - %s\n", enc.Pokemon.Name)
  }
  return nil
}
