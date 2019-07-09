package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

var val = os.Getenv("VAL")

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello " + val + "!")
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
