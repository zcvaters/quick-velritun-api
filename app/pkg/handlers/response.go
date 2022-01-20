package handlers

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}, err error) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status
	log.Printf("API response: %v", body)
	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, err
}