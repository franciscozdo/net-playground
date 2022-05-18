package main

import (
  "net/http"
  "fmt"
)

func test(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "test success")
}

func main() {
  http.HandleFunc("/test", test)

  http.ListenAndServe(":80", nil)
}
