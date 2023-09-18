package route

import (
	"net/http"

	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/app"
	"github.com/gin-gonic/gin"
)

// SetupRoutes -> route setup
func SetupRoutes(address string) error {

	customerHandler := app.NewCustomerHandler()
	productHandler := app.NewProductHandler()
	orderHandler := app.NewOrderHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Savannah Informatics Customer Order Service")
	})

	apiRoutes := r.Group("/api")
	customerRoutes := apiRoutes.Group("/customer")

	{
		customerRoutes.POST("/register", customerHandler.AddCustomer)
		customerRoutes.POST("/signin", customerHandler.SignInCustomer)
	}

	customerProtectedRoutes := apiRoutes.Group("/customers", app.AuthorizeJWT())
	{
		customerProtectedRoutes.GET("/", customerHandler.GetAllCustomer)
		customerProtectedRoutes.GET("/:customer", customerHandler.GetCustomer)
		customerProtectedRoutes.GET("/:customer/products", customerHandler.GetProductOrdered)
		customerProtectedRoutes.PUT("/:customer", customerHandler.UpdateCustomer)
		customerProtectedRoutes.DELETE("/:customer", customerHandler.DeleteCustomer)
	}

	productRoutes := apiRoutes.Group("/products", app.AuthorizeJWT())
	{
		productRoutes.GET("/", productHandler.GetAllProduct)
		productRoutes.GET("/:product", productHandler.GetProduct)
		productRoutes.POST("/", productHandler.AddProduct)
		productRoutes.PUT("/:product", productHandler.UpdateProduct)
		productRoutes.DELETE("/:product", productHandler.DeleteProduct)
	}

	orderRoutes := apiRoutes.Group("/order", app.AuthorizeJWT())
	{
		orderRoutes.POST("/product/:product/quantity/:quantity", orderHandler.OrderProduct)
	}

	messageRoutes := apiRoutes.Group("/messages", app.AuthorizeJWT())
	{
		messageRoutes.POST("/send-bulk-sms", app.SendBulkSMSHandler)
		messageRoutes.GET("/get-africas-talking-settings", app.GetAfricasTalkingSettingsHandler)
	}

	return r.Run(address)

}
