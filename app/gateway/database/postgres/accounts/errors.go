package postgres_account

import "errors"

var (
	errorCreateAccount      = errors.New("erro ao criar nova conta")
	errorAccountIDNotFound  = errors.New("conta não localizada, favor validar o ID informado")
	errorAccountCPFNotFound = errors.New("não foi possível localizar este cpf, favor validar o cpf informado")
)
