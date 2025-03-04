package main

import (
  "time"
  "github.com/codybstrange/pokedexcli/internal/api"
)

func main() {
  client := api.NewClient(5 * time.Second, 5 * time.Minute)
  cfg := &config{
    pokedex: map[string]api.Pokemon{},
    client: client,
  }
  startRepl(cfg)
}

