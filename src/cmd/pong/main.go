package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{ \"message\": \"pong\" }\n")
	})

	log.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
