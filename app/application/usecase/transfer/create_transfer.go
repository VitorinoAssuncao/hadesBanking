package transfer

import (
	"context"

	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (u *usecase) Create(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
	const operation = "Usecase.Transfer.Create"
	u.m.Lock()

	u.logger.LogDebug(operation, "searching for account of origin")
	accountOrigin, err := u.accountRepository.GetByID(ctx, types.ExternalID(transferData.AccountOriginExternalID))
	if err != nil {
		u.logger.LogError(operation, err.Error())
		return transfer.Transfer{}, customError.ErrorTransferCreateOriginError
	}

	u.logger.LogDebug(operation, "searching for account of destiny")
	accountDestiny, err := u.accountRepository.GetByID(ctx, types.ExternalID(transferData.AccountDestinationExternalID))
	if err != nil {
		u.logger.LogError(operation, err.Error())
		return transfer.Transfer{}, customError.ErrorTransferCreateDestinyError
	}

	u.logger.LogDebug(operation, "validating if the funds is sufficient")
	if accountOrigin.Balance < types.Money(transferData.Amount) {
		u.logger.LogError(operation, customError.ErrorTransferCreateInsufficientFunds.Error())
		return transfer.Transfer{}, customError.ErrorTransferCreateInsufficientFunds
	}

	newTransfer := transfer.Transfer{
		AccountOriginID:              types.InternalID(accountOrigin.ID),
		AccountOriginExternalID:      types.ExternalID(accountOrigin.ExternalID),
		AccountDestinationID:         types.InternalID(accountDestiny.ID),
		AccountDestinationExternalID: types.ExternalID(accountDestiny.ExternalID),
		Amount:                       types.Money(transferData.Amount),
	}

	u.logger.LogDebug(operation, "creating the transfer and updating the balances")
	newTransfer, err = u.transferRepository.Create(ctx, newTransfer)
	if err != nil {
		u.logger.LogError(operation, err.Error())
		return transfer.Transfer{}, customError.ErrorTransferCreate
	}
	u.m.Unlock()
	u.logger.LogDebug(operation, "transfer created successfully")
	return newTransfer, nil
}
