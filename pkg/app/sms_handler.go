package app

import (
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
// TODO:
// func SendBulkSMSViaAfricasTalkingHandler(ctx *gin.Context) {

// 	expectedInput := struct {
// 		Message      string   `json:"message"`
// 		PhoneNumbers []string `json:"phone_numbers"`
// 	}{}

// 	ctx.Bind(&expectedInput)

// 	checkErr := api.SendBulkSMSViaAfricasTalking(expectedInput.Message, expectedInput.PhoneNumbers, api.AfricasTalkingSettings)
// 	fmt.Printf("could not send the notifications message to customer:%v", checkErr)
// }
