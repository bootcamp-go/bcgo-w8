package products

import "github.com/bootcamp-go/bcgo-w8/02-go-web/clase-02-TT/live/internal/domain"

type Service interface {
	Create(product *domain.Product) error
	GetAll() ([]domain.Product, error)
}
