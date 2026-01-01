package order

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

func (s *Store) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec("INSERT INTO orders (userId, total, status, address, createdAt) VALUES (:userId, :total, :status, :address, :createdAt)",
		sql.Named("userId", order.UserID),
		sql.Named("total", order.Total),
		sql.Named("status", order.Status),
		sql.Named("address", order.Status),
		sql.Named("createdAt", order.CreatedAt))
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (orderId, productId, quantity, price) VALUES (:orderId, :productId, :quantity, :price)",
		sql.Named("orderId", orderItem.OrderID),
		sql.Named("productId", orderItem.ProductID),
		sql.Named("quantity", orderItem.Quantity),
		sql.Named("price", orderItem.Price))

	return err
}
