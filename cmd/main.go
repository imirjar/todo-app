package main

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.initRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}