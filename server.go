package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	fmt.Fprintf(w, buf.String())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
