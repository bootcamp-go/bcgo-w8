package product

import "errors"

// Model
// Product is an schema for a generic product
type Product struct {
	ID          int		`json:"id"`
	Name        string	`json:"name"`
	Type		string	`json:"type"`
	Count		int		`json:"count"`
	Price		float64	`json:"price"`
	WarehouseId int		`json:"warehouse_id"`	// foreign key
}

type ProductFull struct {
	Product
	WarehouseName string	`json:"warehouse_name"`
	WarehouseAddress string	`json:"warehouse_address"`
}

// Repository interface for product
type Repository interface {
	// GetAll returns all products
	GetAll() (pr []*Product, err error)
	// GetFullData returns all products with warehouse name and address
	GetFullData(id int) (pf *ProductFull, err error)
}
var (
	ErrRepositoryInternal = errors.New("internal repository error")
	ErrRepositoryNotFound = errors.New("product not found")
)