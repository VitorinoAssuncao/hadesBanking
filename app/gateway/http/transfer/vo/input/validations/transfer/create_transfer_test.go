package transfer

import (
	"stoneBanking/app/gateway/http/transfer/vo/input"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateTransferData(t *testing.T) {
	testCases := []struct {
		name    string
		input   input.CreateTransferVO
		want    input.CreateTransferVO
		wantErr bool
	}{
		{
			name: "com os dados corretos, deve passar por todas as validações sem apresentar erro",
			input: input.CreateTransferVO{
				AccountOriginID:  "dd809821-9ecf-411d-bfd8-eca5230439f7",
				AccountDestinyID: "f65fc7d3-26e0-411a-ac3c-785ec328912f",
				Amount:           100,
			},
			want: input.CreateTransferVO{
				AccountOriginID:  "dd809821-9ecf-411d-bfd8-eca5230439f7",
				AccountDestinyID: "f65fc7d3-26e0-411a-ac3c-785ec328912f",
				Amount:           100,
			},
			wantErr: false,
		},
		{
			name: "com o ID repetido entre origem e destino, deverá apresentar erro",
			input: input.CreateTransferVO{
				AccountOriginID:  "dd809821-9ecf-411d-bfd8-eca5230439f7",
				AccountDestinyID: "dd809821-9ecf-411d-bfd8-eca5230439f7",
				Amount:           100,
			},
			want:    input.CreateTransferVO{},
			wantErr: true,
		},
		{
			name: "sem o dado de conta de origem, deve apresentar erro",
			input: input.CreateTransferVO{
				AccountOriginID:  "",
				AccountDestinyID: "f65fc7d3-26e0-411a-ac3c-785ec328912f",
				Amount:           100,
			},
			want:    input.CreateTransferVO{},
			wantErr: true,
		},
		{
			name: "sem o dado de conta de destino, deve apresentar erro",
			input: input.CreateTransferVO{
				AccountOriginID:  "f65fc7d3-26e0-411a-ac3c-785ec328912f",
				AccountDestinyID: "",
				Amount:           100,
			},
			want:    input.CreateTransferVO{},
			wantErr: true,
		},
		{
			name: "com o valor de montante menor ou igual a zero, deve apresentar erro",
			input: input.CreateTransferVO{
				AccountOriginID:  "dd809821-9ecf-411d-bfd8-eca5230439f7",
				AccountDestinyID: "f65fc7d3-26e0-411a-ac3c-785ec328912f",
				Amount:           0,
			},
			want:    input.CreateTransferVO{},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		got, err := ValidateTransferData(test.input)

		assert.Equal(t, (err != nil), test.wantErr)
		assert.Equal(t, test.want, got)
	}
}
