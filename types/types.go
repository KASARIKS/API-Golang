package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	GetProductsByIDs(ps []int) ([]Product, error)
	UpdateProduct(product Product) error
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CreatedAt   string  `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type Order struct {
	ID        int     `json:"id"`
	UserID    int     `json:"userID"`
	Total     float64 `json:"total"`
	Status    string  `json:"status"`
	Address   string  `json:"address"`
	CreatedAt string  `json:"createdAt"`
}

type OrderItem struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"orderID"`
	ProductID int     `json:"productID"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"createdAt"`
}

type CartItem struct {
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

type CartCheckoutPayload struct {
	Items []CartItem `json:"items"`
}
