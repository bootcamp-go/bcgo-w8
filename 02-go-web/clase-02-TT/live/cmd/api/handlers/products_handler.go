package handlers

import (
	"net/http"

	"github.com/bootcamp-go/bcgo-w8/02-go-web/clase-02-TT/live/internal/products"
	"github.com/bootcamp-go/bcgo-w8/02-go-web/clase-02-TT/live/pkg/rest"
	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	Service products.Service
}

func (handler ProductsHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener la petición y validarla.
		var err error
		var request CreateProductRequest

		if err := ctx.ShouldBindJSON(&request); err != nil {
			// TODO: Discriminar errores de validación.
			ctx.JSON(http.StatusBadRequest, rest.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		// Crear el producto.
		productToCreate := request.ToDomain()

		if err = handler.Service.Create(&productToCreate); err != nil {
			switch err {
			case products.ErrProductAlreadyExists:
				ctx.JSON(http.StatusConflict, rest.ErrorResponse{
					Error: "cannot create the given product because it already exists",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{
					Error: "an internal error has occurred",
				})
			}
			return
		}

		// Devolver la respuesta.
		ctx.JSON(http.StatusCreated, rest.SuccessfulResponse{
			Data: CreateProductResponse{
				ID:          productToCreate.ID,
				Name:        productToCreate.Name,
				Description: productToCreate.Description,
				Price:       productToCreate.Price,
			},
		})
	}
}

func (handler ProductsHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener los productos.
		products, err := handler.Service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, rest.ErrorResponse{
				Error: "an internal error has occurred",
			})
			return
		}

		// Devolver la respuesta.
		ctx.JSON(http.StatusOK, rest.SuccessfulResponse{
			Data: products,
		})
	}
}
