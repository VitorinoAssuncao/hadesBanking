package customError

import "errors"

var (
	ErrorAccountCPFExists       = errors.New("já existe uma conta cadastrada com este CPF")
	ErrorCreateAccount          = errors.New("erro ao criar nova conta")
	ErrorAccountIDNotFound      = errors.New("conta não localizada, favor validar o ID informado")
	ErrorAccountLogin           = errors.New("login ou senha inválidos, favor validar")
	ErrorAccountTokenGeneration = errors.New("erro ao gerar o token de acesso")
	ErrorAccountsListing        = errors.New("erro ao carregar a lista com todas as contas")
	ErrorAccountCPFRequired     = errors.New("campo CPF é obrigatório no cadastro de nova conta")
	ErrorAccountCPFInvalid      = errors.New("o CPF adicionado é inválido")
	ErrorAccountCPFWrongSize    = errors.New("o CPF informado é do tamanho incorreto, favor validar novamente")
	ErrorAccountCPFTestNumber   = errors.New("o CPF informado é inválido, em vista de ser considerado um cpf de teste pelo governo federal, favor informar um cpf valido")
	ErrorAccountNameRequired    = errors.New("campo Nome é obrigatório no cadastro de nova conta")
	ErrorAccountSecretRequired  = errors.New("campo Senha é obrigatorio no cadastro de novas contas")
	ErrorAccountBalanceInvalid  = errors.New("campo Saldo, deve ser igual ou maior que zero (0)")

	ErrorServerTokenNotFound = errors.New("não foi possível localizar o token de autenticação, favor logar no sistema novamente")
	ErrorServerExtractToken  = errors.New("não foi possível extrair os dados do token")
)
