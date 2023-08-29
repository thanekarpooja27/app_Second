package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Listening on %s...\n", addr)
	http.ListenAndServe(addr, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
