package main

import (
  "log"
  "net/http"
  "os"
  fun "imbw.com.br/cpf_dv"
)
/*
 * Environment Variables:
 * PORT = 8080
 *
 * Using Google Functions standard env entries:
 * https://cloud.google.com/functions/docs/env-var
 */
func main() {
  entry_point := "/cpf_dv/"
  os.Setenv("ENTRY_POINT", entry_point)
  http.HandleFunc(entry_point, fun.CpfFull)
  log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
