package store

import (
	"database/sql"
	"time"

	"github.com/hailsayan/leskov/internal/types"
)

var QueryTimeoutDuration = time.Second * 5

type IUsers interface {
	GetUserByEmail(email string) (*types.User, error)
	Create(types.User) error
	GetUserByID(id int) (*types.User, error)
}
type IProduct interface {
	GetProducts() ([]*types.Product, error)
	GetProductsByID(ids []int) ([]types.Product, error)
	CreateProduct(types.CreateProductPayload) error
	UpdateProduct(types.Product) error
	GetProductByID(int) (*types.Product, error)
}
type IOrder interface {
	CreateOrder(types.Order) (int, error)
	CreateOrderItem(types.OrderItem) error
}

type Storage struct {
	Users   IUsers
	Product IProduct
	Order   IOrder
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users:   &UserStore{db},
		Product: &ProductStore{db},
		Order:   &OrderStore{db},
	}
}
