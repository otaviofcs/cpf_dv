package main

import (
  "log"
  "net/http"
  "os"
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
  http.HandleFunc(entry_point, CpfFull)
  log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
