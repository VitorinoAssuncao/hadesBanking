package account

import (
	"context"
	"stoneBanking/app/application/vo/input"
	"stoneBanking/app/application/vo/output"
)

func (usecase *usecase) Create(ctx context.Context, accountData input.CreateAccountVO) (*output.AccountOutputVO, error) {
	var accountOutput output.AccountOutputVO
	var err error
	return &accountOutput, err
}
