package services

import (
  "encoding/json"
  "net/http"
  "fmt"
  "shipped/config"
)

type BuildPack struct {
  BuildPackId           string `json:"buildpack_id"`
  Name                  string `json:"name"`
  ImageURL              string `json:"image_url"`
  DefaultBuildCommand   string `json:"default_build_command"`
  DefaultTestCommand    string `json:"default_test_command"`
  DefaultCPU            int    `json:"default_cpu"`
  DefaultRAM            int    `json:"default_ram"`
}

type BuildPacks []BuildPack

var createSqlStmt = `create table if not exists buildpack (
                    buildpack_id text not null primary key,
                    name text,
                    image_url text,
                    default_build_command text,
                    default_test_command text,
                    default_cpu int,
                    default_ram int );`
var insertSqlStmt = `Insert into buildpack(buildpack_id, name, image_url, default_build_command,
                    default_test_command, default_cpu, default_ram) values (?,?,?,?,?,?,?);`
var selectSqlStmt = `select buildpack_id, name, image_url, default_build_command,
                    default_test_command, default_cpu, default_ram from buildpack;`


func BuildPackList(w http.ResponseWriter, r *http.Request) {
  var bps BuildPacks
  rows, err := config.Context.DS.DB.Query(selectSqlStmt)
  if err != nil {
    panic(err)
  }
  for rows.Next(){
    var bp BuildPack
    err = rows.Scan(&bp.BuildPackId, &bp.Name,
               &bp.ImageURL, &bp.DefaultBuildCommand, &bp.DefaultTestCommand,
               &bp.DefaultCPU, &bp.DefaultRAM)
    if err != nil {
      panic(err)
    }
    fmt.Println(bp)
    bps = append(bps,bp)
  }
  if err := json.NewEncoder(w).Encode(bps); err != nil {
    panic(err)
  }
}

func BuildPackCreate(w http.ResponseWriter, r *http.Request){
  var bp BuildPack
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&bp)
  if err != nil {
    panic(err)
  }
  fmt.Println(bp)
  _, err = config.Context.DS.DB.Exec(insertSqlStmt,bp.BuildPackId, bp.Name,
             bp.ImageURL, bp.DefaultBuildCommand, bp.DefaultTestCommand,
             bp.DefaultCPU, bp.DefaultRAM)
  if err != nil {
    panic(err)
  }

  if err = json.NewEncoder(w).Encode(bp); err != nil {
    panic(err)
  }
}

func BuildPackDBInit() error{
  _, err := config.Context.DS.DB.Exec(createSqlStmt)
  if err != nil {
    fmt.Println("%q: %s\n", err, createSqlStmt)
    return err
  }
  return nil
}
