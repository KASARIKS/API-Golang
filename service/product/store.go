package product

import (
	"database/sql"
	"fmt"
	"strings"

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

func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = :name, price = :price, image = :image, description = :description, quantity = :quantity"+
		"WHERER id = :id",
		sql.Named("name", product.Name),
		sql.Named("price", product.Price),
		sql.Named("image", product.Image),
		sql.Named("description", product.Description),
		sql.Named("quantity", product.Quantity),
		sql.Named("id", product.ID))
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetProductsByIDs(productsIDs []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(productsIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	// Have to convert []int to []interface{}
	args := make([]interface{}, len(productsIDs))
	for i, v := range productsIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		p, err := ScanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
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
