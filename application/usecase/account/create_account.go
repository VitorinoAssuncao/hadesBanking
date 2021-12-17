package account

import (
	"context"
	"hades/adapters/vo/input"
	"stoneBanking/domain/vo/output"
)

func (usecase *usecase) Create(ctx context.Context, accountData input.CreateAccountVO) (*output.AccountOutputVO, error) {
	var accountOutput output.AccountOutputVO
	var err error
	return &accountOutput, err
}
