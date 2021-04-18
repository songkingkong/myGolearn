package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloword")
}
func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":80", nil)
}
