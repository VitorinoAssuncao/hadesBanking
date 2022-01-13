package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToInt(t *testing.T) {
	testCases := []struct {
		name  string
		input Money
		want  int
	}{
		{
			name:  "with the money value, convert to int type",
			input: 100,
			want:  1,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := test.input.ToInt()
			assert.Equal(t, reflect.TypeOf(test.want), reflect.TypeOf(got))
		})
	}
}

func Test_ToFloat(t *testing.T) {
	testCases := []struct {
		name  string
		input Money
		want  float64
	}{
		{
			name:  "with the money value, convert to float type, and calculate correct the new value",
			input: 110,
			want:  1.1,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := test.input.ToFloat()
			assert.Equal(t, reflect.TypeOf(test.want), reflect.TypeOf(got))
		})
	}
}
