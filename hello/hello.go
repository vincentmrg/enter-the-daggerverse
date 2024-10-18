package main

import (
	"fmt"
	"net/http"
)

type Toto struct {}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":666", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
