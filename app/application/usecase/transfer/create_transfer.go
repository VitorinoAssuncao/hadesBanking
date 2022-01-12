package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) Create(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {

	accountOrigin, err := usecase.accountRepository.GetByID(ctx, types.ExternalID(transferData.AccountOriginExternalID))
	if err != nil {
		return transfer.Transfer{}, customError.ErrorTransferCreateOriginError
	}

	accountDestiny, err := usecase.accountRepository.GetByID(ctx, types.ExternalID(transferData.AccountDestinationExternalID))
	if err != nil {
		return transfer.Transfer{}, customError.ErrorTransferCreateDestinyError
	}

	if accountOrigin.Balance < types.Money(transferData.Amount) {
		return transfer.Transfer{}, customError.ErrorTransferCreateInsufficientFunds
	}

	newTransfer := transfer.Transfer{
		AccountOriginID:              types.InternalID(accountOrigin.ID),
		AccountOriginExternalID:      types.ExternalID(accountOrigin.ExternalID),
		AccountDestinationID:         types.InternalID(accountDestiny.ID),
		AccountDestinationExternalID: types.ExternalID(accountDestiny.ExternalID),
		Amount:                       types.Money(transferData.Amount),
	}

	newTransfer, err = usecase.transferRepository.Create(ctx, newTransfer)
	if err != nil {
		return transfer.Transfer{}, err
	}

	return newTransfer, nil
}
