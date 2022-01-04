package account

import "errors"

var (
	ErrorAccountCPFExists       = errors.New("já existe uma conta cadastrada com este CPF")
	ErrorCreateAccount          = errors.New("erro ao criar nova conta")
	ErrorAccountIDNotFound      = errors.New("conta não localizada, favor validar o ID informado")
	ErrorAccountLogin           = errors.New("login ou senha inválidos, favor validar")
	ErrorAccountTokenGeneration = errors.New("erro ao gerar o token de acesso")
	ErrorAccountsListing        = errors.New("erro ao carregar a lista com todas as contas")
)
