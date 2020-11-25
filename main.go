package main

import (
  "log"
  "net/http"
  "os"
)
/*
 * Environment Variables:
 * ENTRY_POINT = "/cpfdv/"
 * PORT = 8080
 *
 * Using Google Functions standard env entries:
 * https://cloud.google.com/functions/docs/env-var
 */
func main() {
    http.HandleFunc(os.Getenv("ENTRY_POINT"), CpfFull)
    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
