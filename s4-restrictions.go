package main

import (
  "log"
  "net/http"
  "strings"
)

func serve(w http.ResponseWriter, r *http.Request) {

  // restrict methods
  if r.Method != "GET" {
    return
  }

  stopwords := [5]string {
    "..", // this should be already prevented by design
    "<",
    ">",
    "script",
    "git",
  }

  for _, word := range stopwords {
    if strings.Contains(r.URL.Path, word) {
      return
    }
  }

  p := "./public" + r.URL.Path

  http.ServeFile(w, r, p)

}

func main() {

  http.HandleFunc("/", serve)

  err := http.ListenAndServe(":8030", nil)

  if err != nil {
    log.Fatal(err)
  }

}
