package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TrimCPF(t *testing.T) {
	testCases := []struct {
		name  string
		input Document
		want  string
	}{
		{
			name:  "with a string with special characters return just the numbers",
			input: "324.330.820-80",
			want:  "32433082080",
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := test.input.TrimCPF()
			assert.Equal(t, test.want, string(got))
		})
	}
}

func Test_ToString(t *testing.T) {
	testCases := []struct {
		name  string
		input Document
		want  string
	}{
		{
			name:  "with a document, convert him to string",
			input: "324.330.820-80",
			want:  "32433082080",
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := test.input.TrimCPF().ToString()
			assert.Equal(t, reflect.TypeOf(test.want), reflect.TypeOf(got))
		})
	}
}
