package transfer

import (
	"context"
	"errors"
	"net/http"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/http/middleware"
	"stoneBanking/app/gateway/http/response"
	"stoneBanking/app/gateway/http/transfer/vo/output"
)

//@Sumary Get all transfers
//@Description With a valid Authorization Token, get all the transfers that has made or received by the account
//@Accept json
//@Produce json
//@Param authorization header string true "Authorization Token"
//@Success 200 {object} []output.TransferOutputVO
//@Failure	400 {object} output.OutputError
//@Failure 500 {object} output.OutputError
//@Router /transfers [GET]
func (controller Controller) GetAllByAccountID(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Transfer.GetAllByAccountID"
	resp := response.NewResponse(w)

	accountID, err := middleware.GetAccountIDFromContext(r.Context())
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.BadRequest([]output.OutputError{{Error: err.Error()}})
		return
	}
	transfers, err := controller.usecase.GetAllByAccountID(context.Background(), types.ExternalID(accountID))

	controller.log.LogInfo(operation, "creating the objects to by listed")
	transfersOutput := output.ToTransfersOutput(transfers)
	if err != nil {
		if errors.Is(err, customError.ErrorTransferAccountNotFound) {
			resp.BadRequest(output.OutputError{Error: err.Error()})
			return
		}

		resp.InternalError(output.OutputError{Error: err.Error()})
		return
	}

	controller.log.LogInfo(operation, "transfers listed successfully")
	resp.Ok(transfersOutput)
}
