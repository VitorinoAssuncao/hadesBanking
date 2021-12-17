package errors

//TODO Separar os erros em suas camadas aprópriadas
import "errors"

var (
	ErrorLoginCPFRequired    = errors.New("Campo CPF é obrigatório no login")
	ErrorLoginSecretRequired = errors.New("Campo Senha é obrigatório no login")
	ErrorLoginCPFNotFound    = errors.New("CPF Não localizado, favor validar o cadastro")
	ErrorLoginSecretWrong    = errors.New("Senha incorreta, favor validar novamente")
	ErrorLoginTokenCreation  = errors.New("Erro na geração do token")

	ErrorAccountCPFRequired    = errors.New("Campo CPF é obrigatório no cadastro de nova conta")
	ErrorAccountCPFExists      = errors.New("Já existe uma conta cadastrada com este CPF")
	ErrorAccountCPFNotFound    = errors.New("CPF não localizado, favor validar o cpf informado")
	ErrorAccountCPFNotNumbers  = errors.New("Deve-se informar apenas números no campo de CPF")
	ErrorAccountCPFInvalid     = errors.New("CPF adicionado é inválido")
	ErrorAccountNameRequired   = errors.New("Campo Nome é obrigatório no cadastro de nova conta")
	ErrorAccountSecretRequired = errors.New("Campo Senha é obrigatorio no cadastro de novas contas")
	ErrorAccountBalanceInvalid = errors.New("Campo Saldo, deve ser igual ou maior que zero (0)")
	ErrorCreateAccount         = errors.New("Erro ao criar nova conta")

	ErrorAccountNotFound  = errors.New("Conta não localizada, favor validar o cpf")
	ErrorGetAllAccounts   = errors.New("Erro ao buscar todas as contas")
	ErrorAccountInvalidID = errors.New("ID de Conta informado inválido")

	ErrorTransferCreation         = errors.New("Erro ao criar nova transação")
	ErrorOriginAccountIDRequired  = errors.New("ID de conta de Origem obrigatório para transação")
	ErrorDestinyAccountIDRequired = errors.New("ID de conta de Destino obrigatório para transação")
	ErrorTransferValueIncorrect   = errors.New("Valor de transfêrencia deve ser maior que zero(0)")
	ErrorInsufficientFunds        = errors.New("Saldo da conta insuficiente")
	ErrorAccountOriginNotFound    = errors.New("Conta de origem não localizada")
	ErrorAccountDestinyNotFound   = errors.New("Conta de destino não localizada")
	ErrorTransferSameAccount      = errors.New("Não é possível transferir para a mesma conta")

	ErrorReadingRequest          = errors.New("Erro ao ler o corpo da requisição")
	ErrorDatabaseConnection      = errors.New("Erro ao conectar ao banco de dados")
	ErrorDatabaseCreatingAccount = errors.New("Erro ao criar a conta")
	ErrorDatabaseGetByCPF        = errors.New("Erro ao procurar conta pelo cpf")
)
