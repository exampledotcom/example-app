package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":9090"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello!\nRequest path: %s\nRequest method: %s\n", r.URL.Path, r.Method)
    })
    fmt.Printf("Listening on %s", port)
    http.ListenAndServe(port, nil)
}
