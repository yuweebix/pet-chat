package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yuweebix/pet-chat/pkg/handlers"
	"github.com/yuweebix/pet-chat/pkg/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env:", err)
	}

	_, err := repository.InitDB()
	if err != nil {
		log.Fatal("Failed to initialise the database:", err)
	}

	userMux := http.NewServeMux()
	if err := handlers.InitUserRoutes(userMux); err != nil {
		log.Fatal("Failed to initialise user routes:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/users/", http.StripPrefix("/users", userMux))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("DOMAIN"), os.Getenv("PORT")), mux))
}
