package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	port := os.Getenv("TODO_PORT")

	if port == "" {
		port = "7540"
	}

	r := gin.Default()

	// Static files
	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")

	// favicon
	r.StaticFile("/favicon.ico", "./web/favicon.ico")

	// index
	r.StaticFile("/", "./web/index.html")

	r.Run(":" + port)
}
