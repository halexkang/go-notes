package main

import (
  "fmt"
  "net/http"
  "strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("home!"))
}

func param(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id")) // parse param
  if err != nil || id < 1 {
    http.NotFound(w, r) //404
    return
  }
  fmt.Fprintf(w, "call with param id: %d", id)
}

func post(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost) // display allowed method on error
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    return
  }
  w.Write([]byte("post method"))
}
