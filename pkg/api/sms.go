package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	APIKey              = os.Getenv("AFRICASTALKING_API_KEY_SETTINGS_LABEL")
	UserName            = os.Getenv("AFRICASTALKING_USERNAME_SETTINGS_LABEL")
	BaseLiveEndpoint    = os.Getenv("AFRICASTALKING_BASELIVE_ENDPOINT")
	BaseSandboxEndpoint = os.Getenv("AFRICASTALKING_SANDBOX_ENDPOINT")
)

type (
	// BulkSMSInput is passed to SendBulkSMS as a parameter.
	BulkSMSInput struct {
		To      []string `json:"to"`
		Message string   `json:"message"`
		From    string   `json:"from"`
	}

	// BulkSMSRecipient is returned as part of the BulkSMSResponse.
	BulkSMSRecipient struct {
		StatusCode uint   `json:"statusCode"`
		Number     string `json:"number"`
		Status     string `json:"status"`
		Cost       string `json:"cost"`
		MessageID  string `json:"messageId"`
	}

	// BulkSMSResponse is returned by SendBulkSMS as a response.
	BulkSMSResponse struct {
		SMSMessageData struct {
			Message    string             `json:"message"`
			Recipients []BulkSMSRecipient `json:"recipients"`
		} `json:"SMSMessageData"`
	}

	AfricasTalkingSettings struct {
		ApiKey   string `json:"api_key"`
		Username string `json:"username"`
		Endpoint string `json:"endpoint"`
	}
)

// SendBulkSMSViaAfricasTalking makes a POST request to send bulk SMS's the Africa's Talking and returns a response.
// It uses opinionated defaults.
func SendBulkSMSViaAfricasTalking(message string, phoneNumbers []string) error {
	bulkMessage := []map[string]string{}
	for _, phoneNumber := range phoneNumbers {
		bulkMessage = append(bulkMessage, map[string]string{
			"to":       phoneNumber,
			"meassage": message,
		})
	}

	requestBody := BulkSMSRecipient{}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
		return err
	}

	responce, err := http.Post(BaseSandboxEndpoint, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer responce.Body.Close()

	var res map[string]interface{}

	json.NewDecoder(responce.Body).Decode(&res)
	fmt.Println(res)
	return nil
}

func SendNotificationSMSAfterOrder() {

}
