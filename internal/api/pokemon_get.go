package api

import (
  "encoding/json"
  "io"
  "net/http"
)

func (c *Client) GetPokemon(pokemonName string) (RespPokemon, error) {
  url := baseURL + "/pokemon/" + pokemonName

  if dat, found := c.cache.Get(url); found {
    pokemonResp := RespPokemon{}
    if err := json.Unmarshal(dat, &pokemonResp); err != nil {
      return RespPokemon{}, err
    }
    return pokemonResp, nil
  }

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return RespPokemon{}, err
  }

  resp, err := c.httpClient.Do(req)
  if err != nil {
    return RespPokemon{}, err
  }

  defer resp.Body.Close()

  dat, err := io.ReadAll(resp.Body)
  if err != nil {
    return RespPokemon{}, err
  }

  pokemonResp := RespPokemon{}
  if err := json.Unmarshal(dat, &pokemonResp); err != nil {
    return RespPokemon{}, err
  }
  c.cache.Add(url, dat)
  return pokemonResp, nil
} 
