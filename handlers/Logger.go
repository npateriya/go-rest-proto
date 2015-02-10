package handlers

import (
  "log"
  "net/http"
  "os"
  "time"
)

type Logger struct {
  *log.Logger
}

func NewLogger() *Logger {
  return &Logger{log.New(os.Stdout, "[shipped] ", 0)}
}

func (l *Logger) LoggerHandler(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

    next.ServeHTTP(w, r)
    l.Printf( "%s\t%s\t%s", r.Method, r.RequestURI, time.Since(start), )
  })
}
