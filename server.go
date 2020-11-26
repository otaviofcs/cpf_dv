package function

import (
  "log"
  "net/http"
  "encoding/json"
  "os"
  cpf "imbw.com.br/cpf_dv/cpf"
)

func CpfFull(w http.ResponseWriter, r *http.Request) {
  var baseCpf string = r.URL.Path[len(os.Getenv("ENTRY_POINT")):]
  cpfdv, _ := cpf.CpfComDV(baseCpf)
  v := map[string] interface{}{
    "cpf": cpfdv,
  }
  enc := json.NewEncoder(w)
  if err := enc.Encode(&v); err != nil {
    log.Println(err)
  }
}
