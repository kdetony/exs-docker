package main
import (
    "fmt"
    "net/http")

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hola Multibuild Docker: %s\n", r.URL.Path)
    })

    http.ListenAndServe("0.0.0.0:8080", nil)
}
