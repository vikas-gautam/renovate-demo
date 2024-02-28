package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
