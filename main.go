package main

import (
	"bytes"
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

// IoTEvent IoT buttom event
type IoTEvent struct {
	ClickType    string `json:"clickType"`
	SerialNumber string `json:"serialNumber"`
}

func main() {
	// Disable ssl verification if you dont have valid certificate
	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(apiCall)
}

func apiCall(ctx context.Context, name IoTEvent) {

	// Replace with URL of stackstorm webhook
	url := "https://your.stackstorm.url/api/webhooks/your-webhook-name"

	// Build payload sent to webhook
	var jsonStr = []byte(`{"type":"` + name.ClickType + `", "name":"` + name.SerialNumber + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// Use either API key or token generated from StackStorm for authentification
	req.Header.Set("St2-Api-Key", "YOURSTACKSTORMAPIKEYGOESHERE")
	req.Header.Set("Content-Type", "application/json")

	// Make call
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}
