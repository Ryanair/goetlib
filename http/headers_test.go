package http

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPropagate(t *testing.T) {
	//Given
	expected := map[string]string{
		"authorization": "token",
		"x-b3-traceid":  "16",
	}
	request := events.APIGatewayV2HTTPResponse{
		Headers:    map[string]string{
			"authorization": "token",
			"x-b3-traceid":  "16",
			"anotherHeader": "notPropagate"},
		}

	//When
	result := Propagate(request.Headers, ToPropagate)

	//Then
	assert.Equal(t, expected, result)
}

func TestPropagateWhenExist(t *testing.T) {
	//Given
	expected := map[string]string{
		"authorization": "token",
	}
	request := events.APIGatewayV2HTTPResponse{
		Headers:    map[string]string{
			"authorization": "token",
			"anotherHeader": "notPropagate"},
		StatusCode: 401}

	//When
	result := Propagate(request.Headers, ToPropagate)

	//Then
	assert.Equal(t, expected, result)
}
