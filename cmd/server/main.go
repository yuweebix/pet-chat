package main

import (
	"log"

	"github.com/yuweebix/pet-chat/pkg/repository"
)

func main() {
	_, err := repository.InitDB()
	if err != nil {
		log.Fatal("Failed to initialise the database:", err)
	}

}
