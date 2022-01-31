package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/http/authentication/vo/input"
	validations "stoneBanking/app/gateway/http/authentication/vo/input/validations"
	"stoneBanking/app/gateway/http/authentication/vo/output"
	"stoneBanking/app/gateway/http/response"
)

//@Summary Log in the account
//@Description With the data received, validate if is correct, and log the user, returning a token of authorization
//@Accept json
//@Produce json
//@Param account body input.LoginVO true "Account Login Data"
//@Success 200 {object} output.LoginOutputVO
//@Failure	400 {object} output.OutputError
//@Failure 500 {object} output.OutputError
//@Router /login [POST]
func (controller *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Account.GetBalance"
	resp := response.NewResponse(w)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.BadRequest(output.OutputError{Error: err.Error()})
		return
	}
	defer r.Body.Close()

	controller.log.LogInfo(operation, "unmarshal the data to a internal object")
	var loginData input.LoginVO
	err = json.Unmarshal(reqBody, &loginData)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.BadRequest(output.OutputError{Error: err.Error()})
		return
	}

	controller.log.LogInfo(operation, "validating the input data")
	err = validations.ValidateLoginInputData(loginData)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.BadRequest(output.OutputError{Error: err.Error()})
		return
	}

	account := account.Account{
		CPF:    loginData.CPF.TrimCPF(),
		Secret: types.Password(loginData.Secret),
	}

	controller.log.LogInfo(operation, "trying to log in the system")
	token, err := controller.usecase.LoginUser(context.Background(), account)
	if err != nil {
		if errors.Is(err, customError.ErrorAccountTokenGeneration) {
			controller.log.LogError(operation, err.Error())
			resp.InternalError(output.OutputError{Error: err.Error()})
			return
		}

		controller.log.LogError(operation, err.Error())
		resp.BadRequest(output.OutputError{Error: err.Error()})
		return
	}

	loginOutput := output.LoginOutputVO{
		Authorization: token,
	}

	resp.Ok(loginOutput)
}