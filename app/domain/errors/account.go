package customError

import "errors"

var (
	ErrorAccountCPFExists         = errors.New("an account with this cpf already exist")
	ErrorCreateAccount            = errors.New("error when creating a new account")
	ErrorAccountIDNotFound        = errors.New("account not found, please validate the ID informed")
	ErrorAccountCPFNotFound       = errors.New("account not found, please validate the CPF informed")
	ErrorAccountIDSearching       = errors.New("error when searching for the account")
	ErrorAccountLogin             = errors.New("cpf or secret invalid, please validate then")
	ErrorAccountTokenGeneration   = errors.New("error when generating the authorization token")
	ErrorAccountsListing          = errors.New("error when listing all accounts")
	ErrorAccountCPFRequired       = errors.New("the field 'cpf' is required")
	ErrorAccountCPFInvalid        = errors.New("the value of 'cpf' is not valid")
	ErrorAccountCPFWrongSize      = errors.New("the value of 'cpf' is in the wrong size, please validate")
	ErrorAccountCPFTestNumber     = errors.New("the value of 'cpf' is a test-value, in this case invalid")
	ErrorAccountNameRequired      = errors.New("the field 'Name' is required")
	ErrorAccountSecretRequired    = errors.New("the field 'Secret' is required")
	ErrorAccountBalanceInvalid    = errors.New("the field 'Balance' need to by equal or major than 0(zero)")
	ErrorAccountAcessUnauthorized = errors.New("acess unauthorized, please validate the token and id informed")
	ErrorServerTokenNotFound      = errors.New("authorization token invalid")
	ErrorServerExtractToken       = errors.New("error when extracting the data from token, please validate him")
)
