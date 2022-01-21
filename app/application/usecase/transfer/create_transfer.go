package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) Create(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
	const operation = "Usecase.Transfer.Create"

	usecase.logRepository.LogInfo(operation, "searching for account of origin")
	accountOrigin, err := usecase.accountRepository.GetByID(ctx, types.ExternalID(transferData.AccountOriginExternalID))
	if err != nil {
		usecase.logRepository.LogError(operation, err.Error())
		return transfer.Transfer{}, customError.ErrorTransferCreateOriginError
	}

	usecase.logRepository.LogInfo(operation, "searching for account of destiny")
	accountDestiny, err := usecase.accountRepository.GetByID(ctx, types.ExternalID(transferData.AccountDestinationExternalID))
	if err != nil {
		usecase.logRepository.LogError(operation, err.Error())
		return transfer.Transfer{}, customError.ErrorTransferCreateDestinyError
	}

	usecase.logRepository.LogInfo(operation, "validating if the funds is sufficient")
	if accountOrigin.Balance < types.Money(transferData.Amount) {
		usecase.logRepository.LogError(operation, customError.ErrorTransferCreateInsufficientFunds.Error())
		return transfer.Transfer{}, customError.ErrorTransferCreateInsufficientFunds
	}

	newTransfer := transfer.Transfer{
		AccountOriginID:              types.InternalID(accountOrigin.ID),
		AccountOriginExternalID:      types.ExternalID(accountOrigin.ExternalID),
		AccountDestinationID:         types.InternalID(accountDestiny.ID),
		AccountDestinationExternalID: types.ExternalID(accountDestiny.ExternalID),
		Amount:                       types.Money(transferData.Amount),
	}

	usecase.logRepository.LogInfo(operation, "creating the transfer and updating the balances")
	newTransfer, err = usecase.transferRepository.Create(ctx, newTransfer)
	if err != nil {
		usecase.logRepository.LogError(operation, err.Error())
		return transfer.Transfer{}, customError.ErrorTransferCreate
	}

	usecase.logRepository.LogInfo(operation, "transfer created sucessfully")
	return newTransfer, nil
}
