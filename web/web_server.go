package web

import (
  "log"
  "net/http"
  "text/template"
)

type WebServer struct {
  addr      string
  handlers  map[string]func(w http.ResponseWriter, r *http.Request)
  templates map[string]string
}

func NewWebServer(addr string) *WebServer {
  return &WebServer{
    addr:      addr,
    handlers:  make(map[string]func(w http.ResponseWriter, r *http.Request)),
    templates: make(map[string]string)
  }
}

func index(w http.ResponseWriter, r *http.Request) {
  log.Print("/")
  tmpl, err := template.ParseFiles("index.html")
  if err != nil {
    log.Fatal(err)
  }
  tmpl.Execute(w, "Hi mom")
}

func (self, WebServer) RunServer() {
  mux := http.NewServeMux()
  for key, handler := range self.handlers {
    mux.HandleFunc(key, handler)
  }
  log.Fatal(http.ListenAndServe(self.addr, mux)) 
}

