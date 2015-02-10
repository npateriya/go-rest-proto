package handlers

import (
  "net/http"
  "fmt"
)


func RecoverHandler(next http.Handler) http.Handler {

  // Return http handler for recovery function.
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    // We want to run this after executing service function so defering
    defer func() {
      // If any chained hanlder paniced return internalserver error and log it
      if rec := recover(); rec != nil {
          w.WriteHeader(http.StatusInternalServerError)
          fmt.Println("PANIC recovered err: ", rec)
          fmt.Fprintf(w, "PANIC server with error : %s\n", rec)
      }
    }()

    // Execute passed handler function chain before running defered recovery
    next.ServeHTTP(w, r)
  })
}
