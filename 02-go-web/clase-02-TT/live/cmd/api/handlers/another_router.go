package handlers

import (
	"github.com/gin-gonic/gin"
)

type AnotherRouter struct {
	Engine *gin.Engine

	ProductsHandler ProductsHandler
}

func (router *AnotherRouter) Setup() {
	// Set default middlewares.
	router.Engine.Use(gin.Logger())
	router.Engine.Use(gin.Recovery())

	// Set routes.
	router.SetProductsRoutes()
}

func (router *AnotherRouter) SetProductsRoutes() {
	// Set routes.
	group := router.Engine.Group("/products")
	{
		group.POST("", router.ProductsHandler.Create())
		group.GET("", router.ProductsHandler.GetAll())
	}
}
