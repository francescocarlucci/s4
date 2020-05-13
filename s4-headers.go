package main

import (
  "log"
  "net/http"
)

func serve(w http.ResponseWriter, r *http.Request) {

  // eg. caching, security, custom headers
  w.Header().Add("X-Header", "Value")

  p := "./public" + r.URL.Path

  http.ServeFile(w, r, p)

}

func main() {

  http.HandleFunc("/", serve)

  err := http.ListenAndServe(":8020", nil)

  if err != nil {
    log.Fatal(err)
  }

}
