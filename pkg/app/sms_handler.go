package app

import (
	"log"
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

// SendAfricasTalkingBulkSMSHandler takes expected input from the request and sends the message
func SendAfricasTalkingBulkSMSHandler(c *gin.Context) {
	expectedInput := struct {
		Message    string   `json:"message"`
		Recipients []string `json:"recipients"`
	}{}
	c.Bind(&expectedInput)

	var sandbox api.AfricasTalkingSettings
	africasTalkingSettings, err := api.GetAfricasTalkingSettings(apiKey, username, sandbox)
	if err != nil {
		log.Fatal("error getting africa's Talking settings")
	}

	Err := api.SendAfricastalkingBulkSMS(africasTalkingSettings, expectedInput.Message, expectedInput.Recipients)
	if Err != nil {
		log.Fatal("error sending sms via africa's talking")

	}
	// TODO: save message in the DB
}
