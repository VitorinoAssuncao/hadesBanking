package account

import "errors"

var (
	errorAccountCPFExists       = errors.New("já existe uma conta cadastrada com este CPF")
	errorCreateAccount          = errors.New("erro ao criar nova conta")
	errorAccountIDNotFound      = errors.New("conta não localizada, favor validar o ID informado")
	errorAccountLogin           = errors.New("login ou senha inválidos, favor validar")
	errorAccountTokenGeneration = errors.New("erro ao gerar o token de acesso")
)
