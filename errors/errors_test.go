package errors

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestToBadRequestResponse(t *testing.T) {
	//Given
	event := events.APIGatewayV2HTTPRequest{}
	message := "BAD REQUEST"

	//When
	result := ToBadRequestResponse(event, message)

	//Then
	assert.Equal(t, result.StatusCode, http.StatusBadRequest)
	assert.True(t, strings.Contains(result.Body, message))

}

func TestToUnauthorizedResponse(t *testing.T) {
	//Given
	event := events.APIGatewayV2HTTPRequest{}

	//When
	result := ToUnauthorizedResponse(event)

	//Then
	assert.Equal(t, result.StatusCode, http.StatusUnauthorized)
	assert.Empty(t, result.Body)
}
