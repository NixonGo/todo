package server

import (
	"net/http"
	"os"

	"todo/pkg/api"
)

func Run() {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	// Static files
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	// API
	api.Init() // регистрирует обработчики, например /api/nextdate

	// Запуск сервера
	http.ListenAndServe(":"+port, nil)
}
