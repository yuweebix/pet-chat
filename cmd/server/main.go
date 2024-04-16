package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yuweebix/pet-chat/pkg/handlers"
	"github.com/yuweebix/pet-chat/pkg/middleware"
	"github.com/yuweebix/pet-chat/pkg/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env:", err)
	}

	db, err := repository.InitDB()
	if err != nil {
		log.Fatal("Failed to initialise the database:", err)
	}

	mux, err := handlers.NewRouter(db)
	if err != nil {
		log.Fatal("Failed to initialise routers:", err)
	}

	stack := middleware.CreateStack(
		middleware.Logging,
		// middleware.IsAuthed,
		// middleware.CheckPermissions,
	)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("DOMAIN"), os.Getenv("PORT")),
		Handler: stack(mux),
	}
	log.Fatal(server.ListenAndServe())
}
