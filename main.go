package main

import (
	"log"
	"os"
	"todo/pkg/db"
	"todo/pkg/server"
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

	// Запуск сервера
	server.Run()
}
}
