package products

import "github.com/bootcamp-go/bcgo-w8/02-go-web/clase-02-TT/live/internal/domain"

type SliceBasedRepository struct {
	data []domain.Product
}

func (repository *SliceBasedRepository) Create(product *domain.Product) (err error) {
	product.ID = len(repository.data) + 1
	repository.data = append(repository.data, *product)
	return
}

func (repository *SliceBasedRepository) GetAll() (result []domain.Product, err error) {
	result = repository.data
	return
}
