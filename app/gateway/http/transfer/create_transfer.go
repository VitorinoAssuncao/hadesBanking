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

//@Sumary Create a transfer
//@Description With the data received, validate her and if all is correct, create a new transfer, and update the balance of accounts
//@Accept json
//@Produce json
//@Param authorization header string true "Authorization Token"
//@Param transfer body input.CreateTransferVO true "Transfer Creation Data"
//@Success 200 {object} output.TransferOutputVO
//@Failure	400 {object} output.OutputError
//@Failure 500 {object} output.OutputError
//@Router /transfers [POST]
func (controller Controller) Create(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Transfer.Create"
	resp := response.NewResponse(w)

	controller.log.LogInfo(operation, "getting the account id from the token in the header")
	accountID, err := middleware.GetAccountIDFromToken(r, controller.token)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.BadRequest(output.OutputError{Error: err.Error()})
		return
	}
	defer r.Body.Close()

	var transferData = input.CreateTransferVO{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.BadRequest(output.OutputError{Error: err.Error()})
		return
	}

	controller.log.LogInfo(operation, "unmarshalling the data to a input object")
	json.Unmarshal(reqBody, &transferData) //nolint: errorlint

	transferData.AccountOriginID = accountID

	controller.log.LogInfo(operation, "validating the data")
	transferData, err = validations.ValidateTransferData(transferData)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.BadRequest(output.OutputError{Error: err.Error()})
		return
	}

	controller.log.LogInfo(operation, "transforming in a internal object")
	transfer := transferData.ToEntity()
	newTransfer, err := controller.usecase.Create(context.Background(), transfer)
	if err != nil {
		if !errors.Is(err, customError.ErrorTransferCreate) {
			controller.log.LogError(operation, err.Error())
			resp.BadRequest(output.OutputError{Error: err.Error()})
			return
		}

		controller.log.LogError(operation, err.Error())
		resp.InternalError(output.OutputError{Error: err.Error()})
		return
	}

	transferOutput := output.ToTransferOutput(newTransfer)
	controller.log.LogInfo(operation, "transfer created sucessfully")
	resp.Created(transferOutput)
}
