package api

import (
	"time"

	"gorm.io/gorm"
)

// Customer -> model for customers table
type Customer struct {
	gorm.Model
	Name     string `gorm:"name" json:"name"`
	Code     string `gorm:"code" json:"code"`
	Email    string `gorm:"email,unique" json:"email"`
	Password string `gorm:"password" json:"password"`
}

// TableName --> Table for Customer Model
func (Customer) TableName() string {
	return "customers"
}

// Product --> Model for Product table
type Product struct {
	gorm.Model
	Name        string `gorm:"name" json:"name"`
	Quantity    int    `gorm:"quantity" json:"quantity"`
	Description string `gorm:"description" json:"description"`
}

// TableName --> Table for Product Model
func (Product) TableName() string {
	return "products"
}

// Order --> Model to entity Order
type Order struct {
	gorm.Model
	CustomerID uint
	ProductID  uint
	// Customer   Customer  `gorm:"foreignkey:customerID"`
	// Product    Product   `gorm:"foreignkey:ProductID"`
	Quantity int       `gorm:"quantity" json:"quantity"`
	Amount   int       `gorm:"amount" jsoon:"amount"`
	Time     time.Time `gorm:"time" json:"time"`
}

// TableName --> Table for Order Model
func (Order) TableName() string {
	return "orders"
}
