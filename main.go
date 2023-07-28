package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/go-chi/chi"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading env file.")
	}
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in environment.")
	}

	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(router)
}
