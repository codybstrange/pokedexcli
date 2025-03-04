package api

import (
  "encoding/json"
  "io"
  "net/http"
)

func (c *Client) GetLocation(locationName string) (RespLocation, error) {
  url := baseURL + "/location-area/" + locationName
  if dat, found := c.cache.Get(url); found {
    locationResp := RespLocation{}

    if err := json.Unmarshal(dat, &locationResp); err != nil {
      return RespLocation{}, err
    }
    return locationResp, nil
  }

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return RespLocation{}, err
  }

  resp, err := c.httpClient.Do(req)
  if err != nil {
    return RespLocation{}, err
  }
  defer resp.Body.Close()

  dat, err := io.ReadAll(resp.Body)
  if err != nil {
    return RespLocation{}, err
  }

  locationResp := RespLocation{}
  if err := json.Unmarshal(dat, &locationResp); err != nil {
    return RespLocation{}, err
  }
  c.cache.Add(url, dat)
  return locationResp, nil
}
