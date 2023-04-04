package products

import "github.com/bootcamp-go/bcgo-w8/02-go-web/clase-02-TT/live/internal/domain"

type DefaultService struct {
	Storage Repository
}

func (s DefaultService) Create(product *domain.Product) (err error) {
	// Si el precio es mayor a $100, se aplica un recargo de $10.
	if product.Price > 100 {
		product.Price += 10
	}

	// Crear el producto en la base de datos.
	err = s.Storage.Create(product)
	if err != nil {
		// TODO: Discriminar errores del repositorio.
	}

	return
}

func (s DefaultService) GetAll() (products []domain.Product, err error) {
	products, err = s.Storage.GetAll()
	if err != nil {
		// TODO: Discriminar errores del repositorio.
	}
	return
}
