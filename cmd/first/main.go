package main

import (
  "fmt"
  "net/http"
  "github.com/tt5/first/internal/routes"
)

func main() {
  router := routes.NewRouter()

  port := 8081
  addr := fmt.Sprintf(":%d", port)
  fmt.Printf("server %s\n", addr)
  err := http.ListenAndServe(addr, router)
  if err != nil {
    panic(err)
  }
}
