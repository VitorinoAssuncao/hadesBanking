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

	accountOrigin, err := usecase.accountRepository.GetByID(ctx, types.ExternalID(transferData.AccountOriginID))
	if err != nil {
		return &output.TransferOutputVO{}, ErrorTransferCreateOriginError
	}

	accountDestiny, err := usecase.accountRepository.GetByID(ctx, types.ExternalID(transferData.AccountDestinyID))
	if err != nil {
		return &output.TransferOutputVO{}, ErrorTransferCreateOriginError
	}

	if accountOrigin.Balance < types.Money(transferData.Amount) {
		return &output.TransferOutputVO{}, ErrorTransferCreateInsufficientFunds
	}

	transfer := transfer.Transfer{
		AccountOriginID:              types.InternalID(accountOrigin.ID),
		AccountOriginExternalID:      types.ExternalID(accountOrigin.ExternalID),
		AccountOriginName:            accountOrigin.Name,
		AccountDestinationID:         types.InternalID(accountDestiny.ID),
		AccountDestinationExternalID: types.ExternalID(accountDestiny.ExternalID),
		AccountDestinationName:       accountDestiny.Name,
		Amount:                       types.Money(transferData.Amount),
	}

	transfer, err = usecase.transferRepository.Create(ctx, transfer)
	if err != nil {
		return &output.TransferOutputVO{}, err
	}

	transferOutput := output.TransferToTransferOutput(transfer)

	return &transferOutput, nil
}
