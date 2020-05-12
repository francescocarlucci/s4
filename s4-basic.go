package main

import (
  "log"
  "net/http"
)

func main() {

  http.Handle("/", http.FileServer(http.Dir("./public")))

  err := http.ListenAndServe(":8000", nil)

  if err != nil {
    log.Fatal(err)
  }

}
