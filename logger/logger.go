package logger

import (
	"github.com/Ryanair/gofrlib/log"
	"go.uber.org/zap"
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
		log.With(zap.String("Body.context.trace.spanId", traceId))
		log.With(zap.String("SpanId", traceId))
	}
}

func SetSpanId(spanId string) {
	if spanId != "" {
		log.With(zap.String("Body.context.trace.spanId", spanId))
		log.With(zap.String("SpanId", spanId))
	}
}

func SetRequestInfo(method string, url string, route string, query string, userAgent string) {
	log.WithCustomAttr("Body.context.origin.request.method", method)
	log.WithCustomAttr("Body.context.origin.request.url", url)
	log.WithCustomAttr("Body.context.origin.request.route", route)
	log.WithCustomAttr("Body.context.origin.request.query", query)
	log.WithCustomAttr("Body.context.origin.request.userAgent", userAgent)
}
