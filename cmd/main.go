package main

import (
	"log"

	todo "github.com/0vermotivated/todoapp"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
