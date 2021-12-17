package input

type LoginVO struct {
	CPF    string `json:"cpf" example:"600.246.058-67"`
	Secret string `json:"secret" example:"123456"`
}
