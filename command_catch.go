package main

import (
  "fmt"
  "errors"
  "math/rand"
)

func commandCatch(cfg *config, args ...string) error {
  if len(args) != 1 {
    return errors.New("you must provide a pokemon name")
  }

  name := args[0]
  pokemonResp, err := cfg.client.GetPokemon(name)
  if err != nil {
    return err
  }
  
  fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)

  r := rand.New(rand.NewSource(1))
  p_catch := 50.0/float64(pokemonResp.BaseExperience)
  if r.Float64() < p_catch {
    fmt.Printf("%s was caught!\n", pokemonResp.Name)
    cfg.pokedex[pokemonResp.Name] = pokemonResp
  } else {
    fmt.Printf("%s escaped!\n", pokemonResp.Name)
  }

  return nil
}

