package main

import (
	"github.com/imirjar/todo-app"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}