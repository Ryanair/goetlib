package errors

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type HttpError struct {
	Timestamp 	string	`json:"timestamp"`
	Path		string	`json:"path"`
	StatusCode	int		`json:"status"`
	StatusText	string	`json:"error"`
	Message		string	`json:"message"`
	RequestId	string	`json:"requestId"`
}

func ToBadRequestResponse(event events.APIGatewayV2HTTPRequest, message string) events.APIGatewayV2HTTPResponse {
	return ToResponseWithBody(event, http.StatusBadRequest, message)
}

func ToUnauthorizedResponse(event events.APIGatewayV2HTTPRequest) events.APIGatewayV2HTTPResponse {
	return ToResponseVoid(http.StatusUnauthorized)
}

func ToForbiddenResponse(event events.APIGatewayV2HTTPRequest) events.APIGatewayV2HTTPResponse {
	return ToResponseVoid(http.StatusForbidden)
}

func ToNotFoundResponse(event events.APIGatewayV2HTTPRequest, message string) events.APIGatewayV2HTTPResponse {
	return ToResponseWithBody(event, http.StatusNotFound, message)
}

func ToResponseWithBody(event events.APIGatewayV2HTTPRequest, httpCode int, message string) events.APIGatewayV2HTTPResponse {
	httpError := HttpError{
		fmt.Sprint(event.RequestContext.TimeEpoch),
		event.RawPath,
		httpCode,
		http.StatusText(httpCode),
		message,
		event.RequestContext.RequestID,
	}
	httpErrorMessage,_ := json.Marshal(httpError)

	return events.APIGatewayV2HTTPResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(httpErrorMessage),
		StatusCode: httpCode}

}

func ToResponseVoid(httpCode int) events.APIGatewayV2HTTPResponse {

	return events.APIGatewayV2HTTPResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: httpCode}

}
