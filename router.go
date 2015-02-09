package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "shipped/handlers"
  "shipped/services"
)

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
  Help        string
}

type Routes []Route

var routes = Routes{
  Route{ "Index", "GET", "/", services.Usage,"Prints API usage" },
}

func NewRouter() *mux.Router {

  router := mux.NewRouter().StrictSlash(true)
  logger := handlers.NewLogger()
  for _, route := range routes {
    loggedHandler := logger.LoggerHandler(route.HandlerFunc)
    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(loggedHandler)

  }

  return router
}
