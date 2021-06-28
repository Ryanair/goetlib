package fixtures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsOk(t *testing.T) {
	//Given
	stringSet := []string{"a", "b", "c"}
	message := "b"

	//When
	result := Contains(stringSet, message)

	//Then
	assert.True(t, result)

}

func TestContainsError(t *testing.T) {
	stringSet := []string{"a", "b", "c"}
	message := "m"

	//When
	result := Contains(stringSet, message)

	//Then
	assert.False(t, result)
}
