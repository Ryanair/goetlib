package logger

import (
	"github.com/Ryanair/gofrlib/log"
)

func NewLogConfiguration(logLevel string, application string, project string, projectGroup string) log.Configuration {
	logConfiguration := log.NewConfiguration(
		logLevel,
		application,
		project,
		projectGroup,
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
		log.With("Body.context.trace.spanId", traceId)
		log.With("SpanId", traceId)
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

func SetEvent(source string, body map[string]interface{}, params map[string]interface{}) {
	log.With("Body.context.origin.event.eventSource", source)
	log.With("Body.context.origin.event.eventBody", body)
	log.With("Body.context.origin.event.eventParams", params)
}
