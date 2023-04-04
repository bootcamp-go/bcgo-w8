package handlers

import (
	"github.com/bootcamp-go/bcgo-w8/02-go-web/clase-02-TT/live/internal/products"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func (router *Router) Setup() {
	// Set default middlewares.
	router.Engine.Use(gin.Logger())
	router.Engine.Use(gin.Recovery())

	// Set routes.
	router.SetProductsRoutes()
}

func (router *Router) SetProductsRoutes() {
	// Setup components.
	repository := &products.SliceBasedRepository{}

	service := products.DefaultService{
		Storage: repository,
	}

	handler := ProductsHandler{
		Service: service,
	}

	// Set routes.
	group := router.Engine.Group("/products")
	{
		group.POST("", handler.Create())
		group.GET("", handler.GetAll())
	}
}
