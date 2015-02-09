package services

import (
  "encoding/json"
  "net/http"
)

type Help struct {
  Path        string
  Help        string
}

type Helps []Help

var helps = Helps{
  Help{ "/", "Print help messages",},
}

func Usage(w http.ResponseWriter, r *http.Request) {
  if err := json.NewEncoder(w).Encode(helps); err != nil {
    panic(err)
  }
}
