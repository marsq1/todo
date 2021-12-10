package main

import (
	"log"
	"todo"
	"todo/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)

	if err := srv.Run("8080", handlers.InitRouters()); err != nil {
		log.Fatalf("Error run server: %s", err)
	}
}
