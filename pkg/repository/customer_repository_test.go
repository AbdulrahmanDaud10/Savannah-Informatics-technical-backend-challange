package repository

import (
	"reflect"
	"testing"

	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/api"
	"gorm.io/gorm"
)

func Test_customerRepository_GetCustomer(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantCustomer api.Customer
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &customerRepository{
				connection: tt.fields.connection,
			}
			gotCustomer, err := db.GetCustomer(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("customerRepository.GetCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomer, tt.wantCustomer) {
				t.Errorf("customerRepository.GetCustomer() = %v, want %v", gotCustomer, tt.wantCustomer)
			}
		})
	}
}

func Test_customerRepository_GetByEmail(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantCustomer api.Customer
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &customerRepository{
				connection: tt.fields.connection,
			}
			gotCustomer, err := db.GetByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("customerRepository.GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomer, tt.wantCustomer) {
				t.Errorf("customerRepository.GetByEmail() = %v, want %v", gotCustomer, tt.wantCustomer)
			}
		})
	}
}

func Test_customerRepository_GetAllCustomer(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	tests := []struct {
		name          string
		fields        fields
		wantCustomers []api.Customer
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &customerRepository{
				connection: tt.fields.connection,
			}
			gotCustomers, err := db.GetAllCustomer()
			if (err != nil) != tt.wantErr {
				t.Errorf("customerRepository.GetAllCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomers, tt.wantCustomers) {
				t.Errorf("customerRepository.GetAllCustomer() = %v, want %v", gotCustomers, tt.wantCustomers)
			}
		})
	}
}

func Test_customerRepository_AddCustomer(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	type args struct {
		Customer api.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    api.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &customerRepository{
				connection: tt.fields.connection,
			}
			got, err := db.AddCustomer(tt.args.Customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("customerRepository.AddCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customerRepository.AddCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customerRepository_UpdateCustomer(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	type args struct {
		Customer api.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    api.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &customerRepository{
				connection: tt.fields.connection,
			}
			got, err := db.UpdateCustomer(tt.args.Customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("customerRepository.UpdateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customerRepository.UpdateCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customerRepository_DeleteCustomer(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	type args struct {
		Customer api.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    api.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &customerRepository{
				connection: tt.fields.connection,
			}
			got, err := db.DeleteCustomer(tt.args.Customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("customerRepository.DeleteCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customerRepository.DeleteCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customerRepository_GetProductOrdered(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	type args struct {
		CustomerID int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOrders []api.Order
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &customerRepository{
				connection: tt.fields.connection,
			}
			gotOrders, err := db.GetProductOrdered(tt.args.CustomerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("customerRepository.GetProductOrdered() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("customerRepository.GetProductOrdered() = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}
