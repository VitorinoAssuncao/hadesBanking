package token

type RepositoryMock struct {
	GenerateTokenFunc             func(accountExternalID string) (signedToken string, err error)
	ExtractAccountIDFromTokenFunc func(token string) (accountExternalID string, err error)
}

func (r *RepositoryMock) GenerateToken(accountExternalID string) (signedToken string, err error) {
	return r.GenerateTokenFunc(accountExternalID)
}
func (r *RepositoryMock) ExtractAccountIDFromToken(token string) (accountExternalID string, err error) {
	return r.ExtractAccountIDFromTokenFunc(token)
}
