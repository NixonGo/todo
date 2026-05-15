package main

import (
	"log"
	"os"

	"github.com/NixonGo/todo/pkg/db"
	"github.com/NixonGo/todo/pkg/server"
)

func main() {
	dbFile := "scheduler.db"
	envDB := os.Getenv("TODO_DBFILE")

	if envDB != "" {
		dbFile = envDB
	}

	err := db.Init(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Db.Close()

	// Запуск сервера
	server.Run()
}
