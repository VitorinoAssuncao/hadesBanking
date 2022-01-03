package validations

import "errors"

var (
	errorAccountCPFRequired    = errors.New("campo CPF é obrigatório no cadastro de nova conta")
	errorAccountCPFInvalid     = errors.New("o CPF adicionado é inválido")
	errorAccountCPFWrongSize   = errors.New("o CPF informado é do tamanho incorreto, favor validar novamente")
	errorAccountCPFTestNumber  = errors.New("o CPF informado é inválido, em vista de ser considerado um cpf de teste pelo governo federal, favor informar um cpf valido")
	errorAccountNameRequired   = errors.New("campo Nome é obrigatório no cadastro de nova conta")
	errorAccountSecretRequired = errors.New("campo Senha é obrigatorio no cadastro de novas contas")
	errorAccountBalanceInvalid = errors.New("campo Saldo, deve ser igual ou maior que zero (0)")
)
