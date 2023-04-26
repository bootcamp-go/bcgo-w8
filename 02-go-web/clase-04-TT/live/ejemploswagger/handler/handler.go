package handler

import (
	"ejemploswagger/model"
	"ejemploswagger/service"
	"ejemploswagger/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

//funcion que trae todos los albums
//comentario propio que no afecta a la documentacion
// @Summary List Albums
// @Tags Albums
// @Description Gets all albums without filter
// @Produce json
// @Param token header string true "token"
// @Sucess 200 {object}	model.Album
// @Failure 400 {object} web.Response
// @Router /albums [GET]
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		albums := service.GetAlbums()

		if albums == nil {

			c.AbortWithStatusJSON(http.StatusBadRequest, &web.Response{
				Status: "No se logro manejar la request",
				Code:   400,
			})

			return
		}

		c.IndentedJSON(http.StatusOK, albums)
	}
}

//funcion de creacion
// @Summary Create Album
// @Tags Albums
// @Description Create an album
// @Produce json
// @Param token header string true "token"
// @Param album body model.Album true "album"
// @Sucess 200 {object}	model.Album
// @Failure 400 {object} web.Response
// @Router /albums [POST]
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var album model.Album
		err := c.BindJSON(&album)
		if err != nil {

			c.AbortWithStatusJSON(http.StatusBadRequest, &web.Response{
				Status: "Error en carga de request",
				Code:   400,
			})
			return
		}
		exito := service.Create(album)
		c.IndentedJSON(http.StatusOK, exito)
	}
}
