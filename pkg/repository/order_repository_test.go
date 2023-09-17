package repository

import (
	"testing"

	"gorm.io/gorm"
)

func Test_orderRepository_OrderProduct(t *testing.T) {
	type fields struct {
		connection *gorm.DB
	}
	type args struct {
		customerID    int
		productID int
		quantity  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &orderRepository{
				connection: tt.fields.connection,
			}
			if err := db.OrderProduct(tt.args.customerID, tt.args.productID, tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("orderRepository.OrderProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
