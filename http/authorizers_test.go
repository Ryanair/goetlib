package http

import (
	"github.com/Ryanair/goetlib/logger"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToAuthorizers(t *testing.T) {
	//Given
	logger.InitLoggerTest()

	authorities := map[string]interface{}{
		"sub":         "PRINCIPAL",
		"authorities": []string{"resource:read", "resource:write"}}

	request := events.APIGatewayV2HTTPRequest{
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			Authorizer: &events.APIGatewayV2HTTPRequestContextAuthorizerDescription{
				Lambda: authorities,
			},
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
				Path:   "/api/resource"}}}

	expected := []string{"resource:read", "resource:write"}

	//When
	result := ToAuthorizers(request)

	//Then
	assert.Equal(t, result, expected)
}
