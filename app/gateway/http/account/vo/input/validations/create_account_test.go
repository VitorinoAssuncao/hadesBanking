package input

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"stoneBanking/app/gateway/http/account/vo/input"
)

func Test_ValidateAccountInput(t *testing.T) {
	testCases := []struct {
		name    string
		input   input.CreateAccountVO
		want    input.CreateAccountVO
		wantErr bool
	}{
		{
			name: "right input data, gonna pass all the validations",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "10545395020",
				Secret:  "123456",
				Balance: 0,
			},
			want: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "10545395020",
				Secret:  "123456",
				Balance: 0,
			},
			wantErr: false,
		},
		{
			name: "input data as cpf field void, in this case causing a error in return",
			input: input.CreateAccountVO{
				Name:    "",
				CPF:     "10545395020",
				Secret:  "123456",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "input data as cpf field void, in this case causing a error in return",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "",
				Secret:  "123456",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "input data as cpf field is not valid, in this case causing a error in return",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "10545395021",
				Secret:  "123456",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "input data as cpf as the wrong size, in this case causing a error in return",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "105453950",
				Secret:  "123456",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "input data as test cpf value, in this case causing a error in return",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "11111111111",
				Secret:  "123456",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "input data as secret field void, in this case causing a error in return",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "10545395020",
				Secret:  "",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "input data as a negative balance, in this case causing a error in return",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "10545395020",
				Secret:  "123456",
				Balance: -50,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			got, err := ValidateAccountInput(test.input)

			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
