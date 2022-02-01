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

//@Summary Get all transfers
//@Description With a valid Authorization Token, get all the transfers that has made or received by the account
//@Accept json
//@Produce json
//@Param authorization header string true "Authorization Token"
//@Success 200 {object} []output.TransferOutputVO
//@Failure	400 {object} response.OutputError
//@Failure 500 {object} response.OutputError
//@Router /transfers [GET]
func (c Controller) GetAllByAccountID(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Transfer.GetAllByAccountID"
	resp := response.NewResponse(w)

	accountID, err := middleware.GetAccountIDFromContext(r.Context())
	if err != nil {
		if errors.Is(err, customError.ErrorTransferAccountNotFound) {
			c.log.LogError(operation, err.Error())
			resp.BadRequest(response.NewError(err))
		}

		c.log.LogError(operation, err.Error())
		resp.InternalError(response.NewError(err))
		return
	}
	transfers, err := c.usecase.GetAllByAccountID(context.Background(), types.ExternalID(accountID))

	c.log.LogInfo(operation, "creating the objects to by listed")
	transfersOutput := output.ToTransfersOutput(transfers)
	if err != nil {
		if errors.Is(err, customError.ErrorTransferAccountNotFound) {
			resp.BadRequest(response.NewError(err))
			return
		}

		resp.InternalError(response.NewError(err))
		return
	}

	c.log.LogInfo(operation, "transfers listed successfully")
	resp.Ok(transfersOutput)
}
