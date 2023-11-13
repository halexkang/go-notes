package main

import (
  "log"
  "net/http"
)

func main() {
  mux := http.NewServeMux() // localize servemux for security
  // HandleFunc -> transforms function to a Handler function
  mux.HandleFunc("/", home) // catches all paths
  mux.HandleFunc("/param", param) // func1 exists in handlers.go
  mux.HandleFunc("/post", post) // func2 exists in handlers.go
  log.Print("server running on :5000")
  err := http.ListenAndServe(":5000", mux) // starts an HTTP server with a given address and handler
  log.Fatal(err)
}
