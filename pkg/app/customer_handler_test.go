package app

import (
	"testing"

	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/repository"
	"github.com/gin-gonic/gin"
)

func Test_hashPassword(t *testing.T) {
	type args struct {
		pass *string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			hashPassword(tt.args.pass)
		})
	}
}

func Test_comparePassword(t *testing.T) {
	type args struct {
		dbPass string
		pass   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comparePassword(tt.args.dbPass, tt.args.pass); got != tt.want {
				t.Errorf("comparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customerHandler_GetAllCustomer(t *testing.T) {
	type fields struct {
		repo repository.CustomerRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			h := &customerHandler{
				repo: tt.fields.repo,
			}
			h.GetAllCustomer(tt.args.ctx)
		})
	}
}

func Test_customerHandler_GetCustomer(t *testing.T) {
	type fields struct {
		repo repository.CustomerRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			h := &customerHandler{
				repo: tt.fields.repo,
			}
			h.GetCustomer(tt.args.ctx)
		})
	}
}

func Test_customerHandler_SignInCustomer(t *testing.T) {
	type fields struct {
		repo repository.CustomerRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			h := &customerHandler{
				repo: tt.fields.repo,
			}
			h.SignInCustomer(tt.args.ctx)
		})
	}
}

func Test_customerHandler_AddCustomer(t *testing.T) {
	type fields struct {
		repo repository.CustomerRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			h := &customerHandler{
				repo: tt.fields.repo,
			}
			h.AddCustomer(tt.args.ctx)
		})
	}
}

func Test_customerHandler_UpdateCustomer(t *testing.T) {
	type fields struct {
		repo repository.CustomerRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			h := &customerHandler{
				repo: tt.fields.repo,
			}
			h.UpdateCustomer(tt.args.ctx)
		})
	}
}

func Test_customerHandler_DeleteCustomer(t *testing.T) {
	type fields struct {
		repo repository.CustomerRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			h := &customerHandler{
				repo: tt.fields.repo,
			}
			h.DeleteCustomer(tt.args.ctx)
		})
	}
}

func Test_customerHandler_GetProductOrdered(t *testing.T) {
	type fields struct {
		repo repository.CustomerRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			h := &customerHandler{
				repo: tt.fields.repo,
			}
			h.GetProductOrdered(tt.args.ctx)
		})
	}
}
