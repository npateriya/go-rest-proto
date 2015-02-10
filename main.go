package main

import(
  "log"
  "net/http"
  "shipped/datastore"
  "shipped/config"
  "shipped/services"
  )

func main() {
  router :=   NewRouter()
  config.Context.DS = datastore.NewDatastore("sqlite3","shipped.db")
  services.BuildPackDBInit()
  log.Fatal(http.ListenAndServe(":8888", router))
}
