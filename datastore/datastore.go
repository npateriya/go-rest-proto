package datastore

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  _ "github.com/mattn/go-sqlite3"
  "log"
)

type Datastore struct{
  DB *sql.DB
}

func NewDatastore(driver string, connectString string) (*Datastore){
  DB, err := sql.Open(driver, connectString)
  if err != nil {
    log.Fatal(err)
  }
  return &Datastore{DB}
}
