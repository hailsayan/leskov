package types

import "time"

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}

type CartCheckoutPayload struct {
	Items []CartCheckoutItem `json:"items" validate:"required"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"orderID"`
	ProductID int       `json:"productID"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

type CartCheckoutItem struct {
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}
