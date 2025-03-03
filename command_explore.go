package main

import (
  "fmt"
)

func commandExplore(cfg *config, locationName string) error {
  locationResp, err := cfg.client.ListLocation(locationName)

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
