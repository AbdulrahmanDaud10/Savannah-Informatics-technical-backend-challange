package repository

import (
	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/api"
	"gorm.io/gorm"
)

// CustomerRepository -> Customer CRUD
type CustomerRepository interface {
	AddCustomer(api.Customer) (api.Customer, error)
	GetCustomer(int) (api.Customer, error)
	GetByEmail(string) (api.Customer, error)
	GetAllCustomer() ([]api.Customer, error)
	UpdateCustomer(api.Customer) (api.Customer, error)
	DeleteCustomer(api.Customer) (api.Customer, error)
	GetProductOrdered(int) ([]api.Order, error)
}

type customerRepository struct {
	connection *gorm.DB
}

// NewCustomerRepository --> returns new Customer repository
func NewCustomerRepository() CustomerRepository {
	return &customerRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *customerRepository) GetCustomer(id int) (Customer api.Customer, err error) {
	return Customer, db.connection.First(&Customer, id).Error
}

func (db *customerRepository) GetByEmail(email string) (Customer api.Customer, err error) {
	return Customer, db.connection.First(&Customer, "email=?", email).Error
}

func (db *customerRepository) GetAllCustomer() (Customers []api.Customer, err error) {
	return Customers, db.connection.Find(&Customers).Error
}

func (db *customerRepository) AddCustomer(Customer api.Customer) (api.Customer, error) {
	return Customer, db.connection.Create(&Customer).Error
}

func (db *customerRepository) UpdateCustomer(Customer api.Customer) (api.Customer, error) {
	if err := db.connection.First(&Customer, Customer.ID).Error; err != nil {
		return Customer, err
	}
	return Customer, db.connection.Model(&Customer).Updates(&Customer).Error
}

func (db *customerRepository) DeleteCustomer(Customer api.Customer) (api.Customer, error) {
	if err := db.connection.First(&Customer, Customer.ID).Error; err != nil {
		return Customer, err
	}
	return Customer, db.connection.Delete(&Customer).Error
}

func (db *customerRepository) GetProductOrdered(CustomerID int) (orders []api.Order, err error) {
	return orders, db.connection.Where("Customer_id = ?", CustomerID).Set("gorm:auto_preload", true).Find(&orders).Error
}
