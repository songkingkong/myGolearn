package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	req := r.Body
	fmt.Fprintf(w, "helloword,%v", req)
}
func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":80", nil)
}
