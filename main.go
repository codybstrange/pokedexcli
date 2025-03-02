package main

import (
  "time"
  "github.com/codybstrange/pokedexcli/internal/api"
)

func main() {
  client := api.NewClient(5 * time.Second)
  cfg := &config{
    client: client,
  }
  startRepl(cfg)
}

