package products

import (
	"errors"

	"github.com/bootcamp-go/bcgo-w8/02-go-web/clase-02-TT/live/internal/domain"
)

var (
	ErrProductAlreadyExists = errors.New("product already exists")
)

type Repository interface {
	Create(product *domain.Product) error
	GetAll() ([]domain.Product, error)
}
