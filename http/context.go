package http

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type RequestContext struct {
	headers map[string]string
	payload string
	principal string
	init bool
}

var requestContext RequestContext

func InitRequestContext(headers map[string]string, payload string, principal string) {
	requestContext = RequestContext{headers: headers, payload: payload, principal: principal, init: true}
}

func InitRequestContextFromEvent(event events.APIGatewayV2HTTPRequest) {
	requestContext = RequestContext{
		headers: Propagate(event.Headers, ToPropagate),
		payload: Clean(event.Body),
		principal: getPrincipalFromContext(event),
		init: true,
	}
}

func InitRequestContextFromEventWithHeaders(event events.APIGatewayV2HTTPRequest, propagableEvents []string) {
	requestContext = RequestContext{
		headers: Propagate(event.Headers, propagableEvents),
		payload: Clean(event.Body),
		principal: getPrincipalFromContext(event),
		init: true,
	}
}

func getPrincipalFromContext(event events.APIGatewayV2HTTPRequest) string {
	return event.RequestContext.Authorizer.Lambda["sub"].(string)
}

func flushRequestContext() {
	requestContext = RequestContext{}
}

func GetContext() RequestContext {
	return requestContext
}

func(c *RequestContext) GetHeaders() map[string]string {
	return c.headers
}

func(c *RequestContext) GetPrincipal() string {
	return c.principal
}

func(c *RequestContext) GetPayload() string {
	return c.payload
}

func(c *RequestContext) AddHeadersToRequest(request *http.Request)  {
	for key, val := range c.headers {
		request.Header.Set(key, val)
	}
}

func (c *RequestContext) IsInit() bool {
	return c.init
}