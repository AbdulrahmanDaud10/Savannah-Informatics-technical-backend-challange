package app

import (
	"fmt"
	"net/http"

	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/api"
	"github.com/gin-gonic/gin"
)

func GetAfricasTalkingSettingsHandler(c *gin.Context) {
	apiKey := c.Query("apiKey")
	username := c.Query("username")
	sandbox := c.Query("sandbox") == "true"

	if apiKey == "" || username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	atClient := api.GetAfricasTalkingSettings(apiKey, username, sandbox)

	c.JSON(http.StatusOK, atClient)
}

func SendBulkSMSViaAfricasTalkingHandler(ctx *gin.Context) {
	type ExpectedInput struct {
		Message      string   `json:"message"`
		PhoneNumbers []string `json:"phone_numbers"`
	}

	var expectedInput ExpectedInput
	ctx.Bind(&expectedInput)

	checkErr := api.SendBulkSMSViaAfricasTalking(expectedInput.Message, expectedInput.PhoneNumbers)
	fmt.Printf("could not send the notifications message to customer:%v", checkErr)
}
