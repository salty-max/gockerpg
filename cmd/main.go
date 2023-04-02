package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := mux.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "1664"
	}

	fmt.Printf("📡 Listening on port %s\n", port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
