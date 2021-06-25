package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBodyCurrentCase(t *testing.T) {
	//Given
	request := "{\n\t\"id\": \"1\",\n\t\"firstname\": \"john\",\n\t\"lastname\": \"doe\"\n}"
	expected := "{\"id\": \"1\",\"firstname\": \"john\",\"lastname\": \"doe\"}"

	//When
	result := Clean(request)

	//Then
	assert.Equal(t, expected, result)
}
