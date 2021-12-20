package postgres

import "errors"

var (
	errorCreateAccount      = errors.New("erro ao criar nova conta")
	errorAccountCPFNotFound = errors.New("CPF não localizado, favor validar o cpf informado")
)
