package http

import (
	"encoding/json"
	"github.com/Ryanair/gofrlib/log"
	"github.com/aws/aws-lambda-go/events"
)

func ToAuthorizers(event events.APIGatewayV2HTTPRequest) []string {
	if event.RequestContext.Authorizer == nil || len(event.RequestContext.Authorizer.Lambda) == 0 {
		log.Debug("Error reading Authorizer response, response is not present")
		return nil
	}
	var lambdaAuth UserInfo
	jsonString, err := json.Marshal(event.RequestContext.Authorizer.Lambda)
	if err != nil {
		log.Debug("Error reading Authorizer response", err)
		return nil
	}
	err = json.Unmarshal(jsonString, &lambdaAuth)
	if err != nil {
		log.Debug("Error reading Authorizer response", err)
		return nil
	}

	return lambdaAuth.Authorities
}
