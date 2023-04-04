package main

import (
	"github.com/bootcamp-go/bcgo-w8/02-go-web/clase-02-TT/live/cmd/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Generar un nuevo router en Gin.
	server := gin.New()

	// Configurar el router.
	router := handlers.Router{
		Engine: server,
	}
	router.Setup()

	// Iniciar el servidor.
	server.Run(":8080")
}
