package main

import (
  "log"
  "net/http"
  "encoding/json"
  "os"
  cpf "github.com/otaviofcs/cpf_dv/cpf"
)

func handler(w http.ResponseWriter, r *http.Request) {
  var baseCpf string = r.URL.Path[len("/cpfdv/"):]
  cpfdv, _ := cpf.CpfComDV(baseCpf)
  v := map[string] interface{}{
    "cpf": cpfdv,
  }
  enc := json.NewEncoder(w)
  if err := enc.Encode(&v); err != nil {
    log.Println(err)
  }
}

func main() {
    http.HandleFunc("/cpfdv/", handler)
    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
