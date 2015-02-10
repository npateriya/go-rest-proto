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
  Route{ "BuildPackList", "GET", "/buildpack", services.BuildPackList,"Return list of build packs" },
  Route{ "BuildPackcreate", "POST", "/buildpack", services.BuildPackCreate,"Adds new build packs" },
}

func NewRouter() *mux.Router {

  router := mux.NewRouter().StrictSlash(true)
  logger := handlers.NewLogger()

  for _, route := range routes {
    loggedHandler := logger.LoggerHandler(route.HandlerFunc)
    panicHandler := handlers.RecoverHandler(loggedHandler)
    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(panicHandler)

  }

  return router
}
