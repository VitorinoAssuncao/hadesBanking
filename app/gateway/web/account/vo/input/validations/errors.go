package input

import "errors"

var (
	ErrorAccountCPFRequired    = errors.New("campo CPF é obrigatório no cadastro de nova conta")
	ErrorAccountCPFInvalid     = errors.New("o CPF adicionado é inválido")
	ErrorAccountCPFWrongSize   = errors.New("o CPF informado é do tamanho incorreto, favor validar novamente")
	ErrorAccountCPFTestNumber  = errors.New("o CPF informado é inválido, em vista de ser considerado um cpf de teste pelo governo federal, favor informar um cpf valido")
	ErrorAccountNameRequired   = errors.New("campo Nome é obrigatório no cadastro de nova conta")
	ErrorAccountSecretRequired = errors.New("campo Senha é obrigatorio no cadastro de novas contas")
	ErrorAccountBalanceInvalid = errors.New("campo Saldo, deve ser igual ou maior que zero (0)")
)
