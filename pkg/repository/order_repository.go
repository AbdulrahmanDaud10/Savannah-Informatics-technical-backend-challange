package repository

import (
	"time"

	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/api"
	"gorm.io/gorm"
)

// OrderRepository --> Repository for Order Model
type OrderRepository interface {
	OrderProduct(int, int, int) error
}

type orderRepository struct {
	connection *gorm.DB
}

// NewOrderRepository --> returns new order repository
func NewOrderRepository() OrderRepository {
	return &orderRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *orderRepository) OrderProduct(userID int, productID int, quantity int) error {
	return db.connection.Create(&api.Order{
		Model:      gorm.Model{},
		CustomerID: 0,
		ProductID:  uint(productID),
		Customer:   api.Customer{},
		Product:    api.Product{},
		Quantity:   quantity,
		Amount:     0,
		Time:       time.Time{},
	}).Error

}
