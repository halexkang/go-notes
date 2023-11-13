package main
import (
  "log"
  "net/http"
)
func main() {
  mux := http.NewServeMux() // localize servemux for security
  mux.HandleFunc("/", home) // catches all paths
  mux.HandleFunc("/path/to/func1", func1) // func1 exists in handler.go
  mux.HandleFunc("/path/to/func2", func2) // func2 exists in handler.go
  log.Print("server running on :5000")
  err := http.ListenAndServe(":5000", mux) // starts an HTTP server with a given address and handler
  log.Fatal(err)
}
