package account

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Account struct {

	//ID único da conta cadastrada.
	ID types.AccountID `json:"account_id"`

	//Nome do Próprietario da Conta
	Name string `json:"name"`

	//CPF do Próprietario da Conta
	Cpf string `json:"cpf"`

	//Senha Protegida (Por hash) da conta.
	Secret string `json:"secret"`

	//Saldo atual da conta, em centavos brasileiros (R$)
	Balance types.Money `json:"balance"`

	//Data de criação da conta
	Created_at time.Time `json:"created_at"`
}
