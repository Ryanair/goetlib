package http

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestContextInit(t *testing.T) {
	//Given
	headers := map[string]string{
		"header1": "one",
	}
	payload := "jsonBody"
	principal := "username"
	InitRequestContext(headers, payload, principal)

	//When
	context := GetContext()

	//Then
	assert.Equal(t, headers, context.GetHeaders())
	assert.Equal(t, payload, context.GetPayload())
	assert.Equal(t, principal, context.GetPrincipal())
	assert.True(t, context.IsInit())
}

func TestContextInitFromEvent(t *testing.T) {
	//Given
	headers := map[string]string{
		"authorization":"token",
		"x-b3-traceid":"123",
		"x-b3-spanid":"123",
		"signid": "sign",
	}
	payload := "jsonBody"
	principal := "username"

	request := events.APIGatewayV2HTTPRequest{
		RouteKey: "ANY /api/resource",
		Headers: headers,
		Body: payload,
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			Authorizer: &events.APIGatewayV2HTTPRequestContextAuthorizerDescription{
				Lambda: map[string]interface{}{
					"sub":         principal,
					"authorities": []string{"resource:read"}},
			},
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "POST",
				Path:   "/api/resource"}}}

	InitRequestContextFromEvent(request)

	//When
	context := GetContext()

	//Then
	assert.Equal(t, headers, context.GetHeaders())
	assert.Equal(t, payload, context.GetPayload())
	assert.Equal(t, principal, context.GetPrincipal())
	assert.True(t, context.IsInit())
}

func TestContextInitFromEventWithHeaders(t *testing.T) {
	//Given
	headers := map[string]string{
		"authorization":"token",
		"x-b3-traceid":"123",
		"x-b3-spanid":"123",
		"signid": "sign",
	}
	keyHeaders := []string{"authorization"}
	payload := "jsonBody"
	principal := "username"
	expectedHeaders := map[string]string{"authorization":"token"}

	request := events.APIGatewayV2HTTPRequest{
		RouteKey: "ANY /api/resource",
		Headers: headers,
		Body: payload,
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			Authorizer: &events.APIGatewayV2HTTPRequestContextAuthorizerDescription{
				Lambda: map[string]interface{}{
					"sub":         principal,
					"authorities": []string{"resource:read"}},
			},
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "POST",
				Path:   "/api/resource"}}}

	InitRequestContextFromEventWithHeaders(request, keyHeaders)

	//When
	context := GetContext()

	//Then
	assert.Equal(t, expectedHeaders, context.GetHeaders())
	assert.Equal(t, payload, context.GetPayload())
	assert.Equal(t, principal, context.GetPrincipal())
	assert.True(t, context.IsInit())
}

func TestContextHeadersToRequest(t *testing.T) {
	//Given
	headers := map[string]string{
		"authorization":"token",
		"x-b3-traceid":"123",
		"x-b3-spanid":"123",
		"signid": "sign",
	}
	payload := "jsonBody"
	principal := "username"

	event := events.APIGatewayV2HTTPRequest{
		RouteKey: "ANY /api/resource",
		Headers: headers,
		Body: payload,
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			Authorizer: &events.APIGatewayV2HTTPRequestContextAuthorizerDescription{
				Lambda: map[string]interface{}{
					"sub":         principal,
					"authorities": []string{"resource:read"}},
			},
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "POST",
				Path:   "/api/resource"}}}

	InitRequestContextFromEvent(event)

	request, _ := http.NewRequest(http.MethodPost, "endpoint", nil)

	//When
	ctx := GetContext()
	ctx.AddHeadersToRequest(request)

	//Then
	for k,v := range headers {
		assert.Equal(t, v, request.Header.Get(k))
	}
}

func TestContextNotInit(t *testing.T) {
	//Given
	flushRequestContext()
	//When
	context := GetContext()

	//Then
	assert.Empty(t, context.GetHeaders())
	assert.Empty(t, context.GetPayload())
	assert.Empty(t, context.GetPrincipal())
	assert.False(t, context.IsInit())
}
