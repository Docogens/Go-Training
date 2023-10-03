package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting server\n")
	http.HandleFunc("/healthcheck", HealthCheckHandle)
	fmt.Printf("Server is now on :8080\n")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func HealthCheckHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "healthcheck")
}
