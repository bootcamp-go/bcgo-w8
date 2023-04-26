package main

import (
	"ejemploswagger/docs"
	"ejemploswagger/handler"
	"ejemploswagger/middleware"
	"os"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title EJEMPLO SWAGGER
// @version 1.0
// @description Ejemplo de como implementar swagger en Golang
// @termsOfService https://terminosyservicios.com

// contact.me AlanCano
// contact.url http://linkedin.com

// @license.name Apache 2.0
// @license.url http://licencia.com

// @host localhost/8080
// rutas de la api con los servicios que va a consumir
func main() {
	router := gin.Default()

	//ruta de la documentacion
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.MiddlewareUno())
	router.GET("/albums", handler.GetAll())
	router.POST("/albums", handler.Create())
	router.Run("localhost:8080")
}
