package repository

import (
	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/api"
	"gorm.io/gorm"
)

// ProductRepository --> Interface to ProductRepository
type ProductRepository interface {
	Getproduct(int) (api.Product, error)
	GetAllproduct() ([]api.Product, error)
	AddProduct(api.Product) (api.Product, error)
	UpdateProduct(api.Product) (api.Product, error)
	DeleteProduct(api.Product) (api.Product, error)
}

type productRepository struct {
	connection *gorm.DB
}

// NewProductRepository --> returns new product repository
func NewProductRepository() ProductRepository {
	return &productRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *productRepository) Getproduct(id int) (product api.Product, err error) {
	return product, db.connection.First(&product, id).Error
}

func (db *productRepository) GetAllproduct() (products []api.Product, err error) {
	return products, db.connection.Find(&products).Error
}

func (db *productRepository) AddProduct(product api.Product) (api.Product, error) {
	return product, db.connection.Create(&product).Error
}

func (db *productRepository) UpdateProduct(product api.Product) (api.Product, error) {
	if err := db.connection.First(&product, product.ID).Error; err != nil {
		return product, err
	}
	return product, db.connection.Model(&product).Updates(&product).Error
}

func (db *productRepository) DeleteProduct(product api.Product) (api.Product, error) {
	if err := db.connection.First(&product, product.ID).Error; err != nil {
		return product, err
	}
	return product, db.connection.Delete(&product).Error
}
