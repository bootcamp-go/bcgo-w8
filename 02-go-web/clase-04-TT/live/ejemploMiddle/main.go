package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// rutas de la api con los servicios que va a consumir
func main() {
	router := gin.Default()
	//forma para que todas las rutas utilicen un midle
	//router.Use(middlewareUno)
	router.GET("/albums", getAlbums)
	router.GET("/album/:id",middlewareUno,middlewareDos, getIdAlbum, middlewareTres)
	router.Run("localhost:8080")

}

// estructura de las request
type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

// base de datos del repositorio
var albums = []album{
	{ID: "1", Title: "cualquiera", Artist: "cualquiera", Year: 2002},
}

// servicio que se conecta al repo
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getIdAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, estructura := range albums {
		if estructura.ID == id {
			c.IndentedJSON(http.StatusOK, estructura)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No se encontro el album"})

}

//ejemplos de middlewares

func middlewareUno(c *gin.Context) {
	log.Println("Este es el primer middleware")
	token := c.GetHeader("token")
	if token != "123" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "El token no es el mismo"})
		return
	}
	c.Next()
}

func middlewareDos(c *gin.Context) {
	log.Println("Este es el segundo endpoint")
	id := c.Param("id")
	if id == "0" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "El id debe ser distinto de cero")
		return
	}
	c.Next()
}

func middlewareTres(c *gin.Context) {
	log.Println("Este el tercer y ultimo middleware")
	c.String(200, "Todo salio de maravillas!")
	c.Next()
}
