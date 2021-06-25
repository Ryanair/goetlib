package http

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestToResponse(t *testing.T) {
	//Given
	payload := map[string]interface{}{
		"id": "1234",
		"firstname": "john",
		"lastname":	"doe",
	}
	jsonBytes,_ := json.Marshal(payload)
	jsonPayload := string(jsonBytes)
	status := http.StatusOK

	//When
	result := ToResponse(payload, status)

	//Then
	assert.Equal(t, result.StatusCode, status)
	assert.Equal(t, result.Body, jsonPayload)
}
