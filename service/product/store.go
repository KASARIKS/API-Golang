package product

import (
	"database/sql"

	"github.com/kasariks/api_golang/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	products := []types.Product{}

	rows, err := s.db.Query("SELECT * FROM products;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product, err := ScanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	return products, nil
}

func ScanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
