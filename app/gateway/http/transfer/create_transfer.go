package transfer

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/http/middleware"
	"stoneBanking/app/gateway/http/response"
	"stoneBanking/app/gateway/http/transfer/vo/input"
	validations "stoneBanking/app/gateway/http/transfer/vo/input/validations/transfer"
	"stoneBanking/app/gateway/http/transfer/vo/output"
)

//@Summary Create a transfer
//@Description With the data received, validate her and if all is correct, create a new transfer, and update the balance of accounts
//@Accept json
//@Produce json
//@Param authorization header string true "Authorization Token"
//@Param transfer body input.CreateTransferVO true "Transfer Creation Data"
//@Success 200 {object} output.TransferOutputVO
//@Failure	400 {object} response.OutputError
//@Failure 500 {object} response.OutputError
//@Router /transfers [POST]
func (c Controller) Create(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Transfer.Create"
	c.log.SetRequestIDFromContext(r.Context())
	resp := response.NewResponse(w)

	c.log.LogDebug(operation, "getting the account id from the token in the header")
	accountID, err := middleware.GetAccountIDFromContext(r.Context())
	if err != nil {
		c.log.LogError(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}
	defer r.Body.Close()

	c.log.LogInfo(operation, "receive the body and unmarshal the data")
	var transferInput input.CreateTransferVO
	if err = json.NewDecoder(r.Body).Decode(&transferInput); err != nil {
		c.log.LogError(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	transferInput.AccountOriginID = accountID

	c.log.LogInfo(operation, "validating the data")

	if transferInput, err = validations.ValidateTransferData(transferInput); err != nil {
		c.log.LogError(operation, err.Error())
		resp.BadRequest(output.OutputError{Error: err.Error()})
		return
	}

	c.log.LogInfo(operation, "transforming in a internal object")
	transfer := transferInput.ToEntity()
	newTransfer, err := c.usecase.Create(context.Background(), transfer)
	if err != nil {
		if !errors.Is(err, customError.ErrorTransferCreate) {
			c.log.LogWarn(operation, err.Error())
			resp.BadRequest(response.NewError(err))
			return
		}

		c.log.LogError(operation, err.Error())
		resp.InternalError(response.NewError(err))
		return
	}

	transferOutput := output.ToTransferOutput(newTransfer)
	c.log.LogDebug(operation, "transfer created successfully")
	resp.Created(transferOutput)
}
