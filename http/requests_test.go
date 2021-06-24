package http

import (
	"github.com/Ryanair/goetlib/logger"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToRequestWhenGet(t *testing.T) {
	//Given
	logger.InitLoggerTest()

	authorities := map[string]interface{}{
		"sub":         "PRINCIPAL",
		"authorities": []string{"resource:read", "resource:write"}}

	event := events.APIGatewayV2HTTPRequest{
		RouteKey: "ANY /api/resource",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			Authorizer: &events.APIGatewayV2HTTPRequestContextAuthorizerDescription{
				Lambda: authorities,
			},
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
				Path:   "/api/resource"}}}

	expected := Request{
		Method: "GET",
		Path:   "/api/resource"}

	//When
	result := ToRequest(event)

	//Then
	assert.Equal(t, result, expected)
}

func TestToRequestWhenGetWithSubResource(t *testing.T) {
	//Given
	logger.InitLoggerTest()

	authorities := map[string]interface{}{
		"sub":         "PRINCIPAL",
		"authorities": []string{"resource:read", "resource:write"}}

	event := events.APIGatewayV2HTTPRequest{
		RouteKey: "ANY /api/resource/{proxy+}",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			Authorizer: &events.APIGatewayV2HTTPRequestContextAuthorizerDescription{
				Lambda: authorities,
			},
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
				Path:   "/api/resource/123456"}}}

	expected := Request{
		Method: "GET",
		Path:   "/api/resource/{proxy+}"}

	//When
	result := ToRequest(event)

	//Then
	assert.Equal(t, result, expected)
}
