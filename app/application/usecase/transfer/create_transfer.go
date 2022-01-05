package transfer

import (
	"context"
	validations "stoneBanking/app/application/validations/transfer"
	"stoneBanking/app/application/vo/input"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) Create(ctx context.Context, transferData input.CreateTransferVO) (*output.TransferOutputVO, error) {
	transferData, err := validations.ValidateTransferData(transferData)

	if err != nil {
		return &output.TransferOutputVO{}, err
	}

	accountOrigin, err := usecase.accountRepository.GetByID(ctx, types.AccountExternalID(transferData.AccountOriginID))
	if err != nil {
		return &output.TransferOutputVO{}, ErrorTransferCreateOriginError
	}

	accountDestiny, err := usecase.accountRepository.GetByID(ctx, types.AccountExternalID(transferData.AccountDestinyID))
	if err != nil {
		return &output.TransferOutputVO{}, ErrorTransferCreateOriginError
	}

	if accountOrigin.Balance < types.Money(transferData.Amount) {
		return &output.TransferOutputVO{}, ErrorTransferCreateInsufficientFunds
	}

	transfer := transfer.Transfer{
		AccountOriginID:      types.AccountExternalID(transferData.AccountOriginID),
		AccountDestinationID: types.AccountExternalID(transferData.AccountDestinyID),
		Amount:               types.Money(transferData.Amount),
	}

	accountOrigin.Balance -= types.Money(transferData.Amount)
	accountDestiny.Balance += types.Money(transferData.Amount)

	usecase.accountRepository.UpdateBalance(ctx, accountOrigin.Balance.ToInt(), accountOrigin.ExternalID)
	usecase.accountRepository.UpdateBalance(ctx, accountDestiny.Balance.ToInt(), accountDestiny.ExternalID)
	transfer, err = usecase.transferRepository.Create(ctx, transfer)
	if err != nil {
		return &output.TransferOutputVO{}, err
	}

	transferOutput := output.TransferToTransferOutput(transfer)

	return &transferOutput, nil
}