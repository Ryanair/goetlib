package http

import (
	"github.com/aws/aws-lambda-go/events"
	"strings"
)

type Request struct {
	Method string
	Path   string
}

func ToRequest(event events.APIGatewayV2HTTPRequest) Request {
	return Request{
		Method: event.RequestContext.HTTP.Method,
		Path:	strings.Replace(event.RouteKey, "ANY ", "", 1),
	}
}