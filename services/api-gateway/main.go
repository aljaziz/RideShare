package main

import (
	"aljaziz/RideShare/shared/env"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	godotenv.Load()
	log.Println("Starting API Gateway")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from API Gateway"))
	})

	http.ListenAndServe(httpAddr, nil)
}
