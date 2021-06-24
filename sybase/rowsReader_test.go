package sybase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type rowsMock struct {
	mock.Mock
}

func (m *rowsMock) Columns() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *rowsMock) Next() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *rowsMock) Scan(dest ...interface{}) error {
	args := m.Called(dest)
	d := dest[0].(*interface{})
	if *d == nil {
		*d = 5740147
	} else {
		*d = 5740148
	}
	return args.Error(0)
}

func TestConvertToJson(t *testing.T) {
	//Given
	rows := new(rowsMock)
	columns := []string{"column_id"}
	rows.On("Columns").Return(columns, nil)
	rows.On("Next").Return(true).Once()
	rows.On("Next").Return(true).Once()
	rows.On("Next").Return(false)
	rows.On("Scan", mock.Anything).Return(nil)
	expected := ([]map[string]interface{}{{"column_id": 5740147}, {"column_id": 5740148}})

	//When
	result, _ := ToMap(rows)

	//Then
	assert.Equal(t, expected, result)
}
