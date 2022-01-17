package input

import "stoneBanking/app/domain/types"

type LoginVO struct {
	CPF    types.Document `json:"cpf" example:"600.246.058-67"`
	Secret types.Password `json:"secret" example:"123456"`
}
