package logger

import (
	"encoding/json"
	"github.com/Ryanair/gofrlib/log"
	"github.com/aws/aws-lambda-go/events"
)

func NewLogConfiguration(logLevel string, application string, project string, projectGroup string, version string) log.Configuration {
	logConfiguration := log.NewConfiguration(
		logLevel,
		application,
		project,
		projectGroup,
		version,
		"")
	return logConfiguration
}

func InitLogger(logConfiguration log.Configuration) {
	log.Init(logConfiguration)
}

func InitLoggerTest() {
	logConfiguration := log.NewConfiguration(
		"debug",
		"test",
		"test-project",
		"test-project-group",
		"version",
		"")

	InitLogger(logConfiguration)
}

func InitialLambdaConfiguration(functionName string, logGroup string, logStream string, vpc string, env string) {
	log.With("function_name", functionName)
	log.With("@log_group", logGroup)
	log.With("@log_stream", logStream)
	log.With("@vpc", vpc)
	log.With("@env", env)
}

func SetTraceId(traceId string) {
	if traceId != "" {
		log.With("Body.context.trace.traceId", traceId)
		log.With("TraceId", traceId)
	}
}

func SetSpanId(spanId string) {
	if spanId != "" {
		log.With("Body.context.trace.spanId", spanId)
		log.With("SpanId", spanId)
	}
}

func SetRequestInfo(method string, url string, route string, query string, userAgent string) {
	log.With("Body.context.origin.request.method", method)
	log.With("Body.context.origin.request.url", url)
	log.With("Body.context.origin.request.route", route)
	log.With("Body.context.origin.request.query", query)
	log.With("Body.context.origin.request.userAgent", userAgent)
}

func SetEvent(source string, body string, params map[string]string) {
	log.With("Body.context.origin.event.eventSource", source)
	log.With("Body.context.origin.event.eventBody", body)
	log.With("Body.context.origin.event.eventParams", params)
}

func SetSQSEvent(event events.SQSMessage) {
	log.With("Body.context.origin.event.eventSource", event.EventSource)
	log.With("Body.context.origin.event.eventBody", event.Body)

	raw, _ := json.Marshal(event.Attributes)
	log.With("Body.context.origin.event.eventParams", string(raw))
}

func SetApigwRequest(event events.APIGatewayV2HTTPRequest) {
	log.With("Body.context.origin.request.method", event.RequestContext.HTTP.Method)
	log.With("Body.context.origin.request.url", event.RequestContext.DomainName)
	log.With("Body.context.origin.request.route", event.RequestContext.RouteKey)
	log.With("Body.context.origin.request.query", event.RawQueryString)
	log.With("Body.context.origin.request.userAgent", event.RequestContext.HTTP.UserAgent)
}