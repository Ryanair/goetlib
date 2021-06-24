package http

import (
	"encoding/json"
	"github.com/Ryanair/gofrlib/log"
	"github.com/aws/aws-lambda-go/events"
)

func ToResponse(rawData interface{}, status int) events.APIGatewayV2HTTPResponse {
	log.Debug("RawData: ", rawData)
	var response string
	jsonBytes, err := json.Marshal(rawData)
	if err != nil {
		log.Error("Error converting to json ", rawData)
		response = "[]"
	} else {
		log.Debug("JSON:", string(jsonBytes))
		response = string(jsonBytes)
		if response == "null" {
			response = "[]"
		}
	}
	return events.APIGatewayV2HTTPResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:       response,
		StatusCode: status}
}
