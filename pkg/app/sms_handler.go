package app

import (
	"context"
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

func SendBulkSMSHandler(c *gin.Context) {
	var input api.BulkSMSInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	atClient := api.GetAfricasTalkingSettings(c.Query("apiKey"), c.Query("username"), false)

	bulkSMSResponse, err := atClient.SendBulkSMS(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bulkSMSResponse)
}
