package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/api"
	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CustomerHandler -> interface to Customer entity
type CustomerHandler interface {
	AddCustomer(*gin.Context)
	GetCustomer(*gin.Context)
	GetAllCustomer(*gin.Context)
	SignInCustomer(*gin.Context)
	UpdateCustomer(*gin.Context)
	DeleteCustomer(*gin.Context)
	GetProductOrdered(*gin.Context)
}

type customerHandler struct {
	repo repository.CustomerRepository
}

// NewCustomerHandler --> returns new customer handler
func NewCustomerHandler() CustomerHandler {
	return &customerHandler{
		repo: repository.NewCustomerRepository(),
	}
}

func hashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func (h *customerHandler) GetAllCustomer(ctx *gin.Context) {
	fmt.Println(ctx.Get("customerID"))
	customer, err := h.repo.GetAllCustomer()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, customer)

}

func (h *customerHandler) GetCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err := h.repo.GetCustomer(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, customer)

}

func (h *customerHandler) SignInCustomer(ctx *gin.Context) {
	var customer api.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	dbCustomer, err := h.repo.GetByEmail(customer.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "No Such Customer Found"})
		return

	}
	if isTrue := comparePassword(dbCustomer.Password, customer.Password); isTrue {
		fmt.Println("customer before", dbCustomer.ID)
		token := api.GenerateToken(dbCustomer.ID)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully SignedIN", "token": token})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Password not matched"})

}

func (h *customerHandler) AddCustomer(ctx *gin.Context) {
	var customer api.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashPassword(&customer.Password)
	customer, err := h.repo.AddCustomer(customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	customer.Password = ""
	ctx.JSON(http.StatusOK, customer)

}

func (h *customerHandler) UpdateCustomer(ctx *gin.Context) {
	var customer api.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("uscustomerer")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	customer.ID = uint(intID)
	customer, err = h.repo.UpdateCustomer(customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, customer)

}

func (h *customerHandler) DeleteCustomer(ctx *gin.Context) {
	var customer api.Customer
	id := ctx.Param("customer")
	intID, _ := strconv.Atoi(id)
	customer.ID = uint(intID)
	customer, err := h.repo.DeleteCustomer(customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, customer)

}

func (h *customerHandler) GetProductOrdered(ctx *gin.Context) {
	customerStr := ctx.Param("customer")
	customerID, _ := strconv.Atoi(customerStr)
	if products, err := h.repo.GetProductOrdered(customerID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, products)
	}
}
