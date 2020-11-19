package main

import (
  "fmt"
  "log"
  "net/http"
  cpf "github.com/otaviofcs/cpf_dv/cpf"
)

func handler(w http.ResponseWriter, r *http.Request) {
  var baseCpf string = r.URL.Path[len("/cpfdv/"):]
  cpfdv, _ := cpf.CpfComDV(baseCpf)
  fmt.Fprintf(w, "Hi, complete CPF is: '%s'", cpfdv)
}

func main() {
    http.HandleFunc("/cpfdv/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
