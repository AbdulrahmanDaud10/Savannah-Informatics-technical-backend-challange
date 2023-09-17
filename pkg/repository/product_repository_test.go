package repository

import (
	"reflect"
	"testing"

	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/api"
	"gorm.io/gorm"
)

func Test_productRepository_Getproduct(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantProduct api.Product
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &productRepository{
				connection: tt.fields.connection,
			}
			gotProduct, err := db.Getproduct(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("productRepository.Getproduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProduct, tt.wantProduct) {
				t.Errorf("productRepository.Getproduct() = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}

func Test_productRepository_GetAllproduct(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	tests := []struct {
		name         string
		fields       fields
		wantProducts []api.Product
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &productRepository{
				connection: tt.fields.connection,
			}
			gotProducts, err := db.GetAllproduct()
			if (err != nil) != tt.wantErr {
				t.Errorf("productRepository.GetAllproduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProducts, tt.wantProducts) {
				t.Errorf("productRepository.GetAllproduct() = %v, want %v", gotProducts, tt.wantProducts)
			}
		})
	}
}

func Test_productRepository_AddProduct(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	type args struct {
		product api.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    api.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &productRepository{
				connection: tt.fields.connection,
			}
			got, err := db.AddProduct(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("productRepository.AddProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productRepository.AddProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productRepository_UpdateProduct(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	type args struct {
		product api.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    api.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &productRepository{
				connection: tt.fields.connection,
			}
			got, err := db.UpdateProduct(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("productRepository.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productRepository.UpdateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productRepository_DeleteProduct(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	type args struct {
		product api.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    api.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &productRepository{
				connection: tt.fields.connection,
			}
			got, err := db.DeleteProduct(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("productRepository.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productRepository.DeleteProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
