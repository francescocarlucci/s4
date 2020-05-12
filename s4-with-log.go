package main

import (
  "log"
  "net/http"
)

func serve(w http.ResponseWriter, r *http.Request) {

  log.Println(r.URL.Path, r.Method) // log to console, not file

  p := "./public" + r.URL.Path

  http.ServeFile(w, r, p)

}

func main() {

  http.HandleFunc("/", serve)

  err := http.ListenAndServe(":8040", nil)

  if err != nil {
    log.Fatal(err)
  }

}
