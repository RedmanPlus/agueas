package web

import (
  "log"
  "net/http"
  "text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
  log.Print("/")
  tmpl, err := template.ParseFiles("index.html")
  if err != nil {
    log.Fatal(err)
  }
  tmpl.Execute(w, "Hi mom")
}

func RunServer() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", index)
  log.Fatal(http.ListenAndServe(":8000", mux)) 
}

