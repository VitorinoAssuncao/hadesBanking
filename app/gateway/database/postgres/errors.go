package postgres

import "errors"

var (
	errorCreateAccount      = errors.New("erro ao criar nova conta")
	errorGetByCPF           = errors.New("erro ao procurar o cpf, favor tentar novamente")
	errorAccountCPFNotFound = errors.New("CPF não localizado, favor validar o cpf informado")
)
