package product

import (
	"database/sql"
)

const (
	QueryGetAll = "SELECT id, name, type, count, price, warehouse_id FROM product"
	QueryGetFull = "SELECT pr.id, pr.name, pr.type, pr.count, pr.price, pr.warehouse_id, wh.name, wh.address FROM product pr " +
					"INNER JOIN warehouse wh ON pr.warehouse_id = wh.id " +
					"WHERE pr.id = ?"
)

func NewRepositoryMySQL(db *sql.DB) Repository {
	return &RepositoryMySQL{db: db}
}

type RepositoryMySQL struct {
	db *sql.DB
}
func (rp *RepositoryMySQL) GetAll() (pr []*Product, err error) {
	// prepare statement
	var stmt *sql.Stmt
	stmt, err = rp.db.Prepare(QueryGetAll)
	if err != nil {
		err = ErrRepositoryInternal
		return
	}
	defer stmt.Close()

	var rows *sql.Rows
	rows, err = stmt.Query()
	if err != nil {
		err = ErrRepositoryInternal
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price, &p.WarehouseId)
		if err != nil {
			err = ErrRepositoryInternal
			return
		}

		pr = append(pr, &p)
	}

	return
}

func (rp *RepositoryMySQL) GetFullData(id int) (pf *ProductFull, err error) {
	// prepare statement
	row := rp.db.QueryRow(QueryGetFull, id)

	// scan
	pf = new(ProductFull)
	err = row.Scan(&pf.ID, &pf.Name, &pf.Type, &pf.Count, &pf.Price, &pf.WarehouseId, &pf.WarehouseName, &pf.WarehouseAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrRepositoryNotFound
			return
		}
		err = ErrRepositoryInternal
		return
	}

	return
}