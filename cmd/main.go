package main

import (
	"log"

	"github.com/drakond/todo-app"
	"github.com/drakond/todo-app/pkg/handler"

	"github.com/drakond/todo-app/pkg/repository"
	"github.com/drakond/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// Создаем новый экземпляр сервера приложения.
	srv := new(todo.Server)

	// Указываем порт для прослушивания.
	port := ":8000"

	// Запускаем сервер и проверяем ошибки.
	if err := srv.Run(port, handlers.InitRoutes()); err != nil { // Изменен вызов InitRoutes()
		// Если произошла ошибка, выводим сообщение об ошибке и завершаем работу приложения.
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
