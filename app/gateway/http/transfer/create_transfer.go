package transfer

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
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

	var transferData = input.CreateTransferVO{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.log.LogError(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	c.log.LogDebug(operation, "unmarshalling the data to a input object")
	json.Unmarshal(reqBody, &transferData) //nolint: errorlint

	transferData.AccountOriginID = accountID

	c.log.LogDebug(operation, "validating the data")
	transferData, err = validations.ValidateTransferData(transferData)
	if err != nil {
		c.log.LogWarn(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	c.log.LogDebug(operation, "transforming in a internal object")
	transfer := transferData.ToEntity()
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
