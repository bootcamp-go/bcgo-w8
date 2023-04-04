package handlers

import "github.com/bootcamp-go/bcgo-w8/02-go-web/clase-02-TT/live/internal/domain"

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

func (request CreateProductRequest) ToDomain() domain.Product {
	return domain.Product{
		ID:          0,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}
}
