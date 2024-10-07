package validator_rules

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
)

type mockFieldLevel struct {
	validator.FieldLevel
	data string
}

func (m mockFieldLevel) Field() reflect.Value {
	return reflect.ValueOf(m.data)
}

func Test_Password(t *testing.T) {
	tests := []struct {
		name     string
		password string
		error    bool
	}{
		{
			name:     "Valid strong password",
			password: "StrongP@ssw0rd123!",
			error:    false,
		},
		{
			name:     "Invalid weak password",
			password: "weak",
			error:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mockFieldLevel{data: tt.password}

			assert.Equal(t, Password(mock), !tt.error)
		})
	}
}
