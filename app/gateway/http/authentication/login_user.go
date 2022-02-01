package authentication

import (
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
//@Failure	400 {object} response.OutputError
//@Failure 500 {object} response.OutputError
//@Router /login [POST]
func (c *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Authorization.Login"
	resp := response.NewResponse(w)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.LogWarn(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}
	defer r.Body.Close()

	log.LogInfo(operation, "unmarshal the data to a internal object")
	var loginData input.LoginVO
	err = json.Unmarshal(reqBody, &loginData)
	if err != nil {
		log.LogWarn(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	log.LogInfo(operation, "validating the input data")
	err = validations.ValidateLoginInputData(loginData)
	if err != nil {
		log.LogWarn(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	account := account.Account{
		CPF:    loginData.CPF.TrimCPF(),
		Secret: types.Password(loginData.Secret),
	}

	log.LogInfo(operation, "trying to log in the system")
	token, err := c.usecase.LoginUser(r.Context(), account)
	if err != nil {
		if errors.Is(err, customError.ErrorAccountTokenGeneration) {
			log.LogWarn(operation, err.Error())
			resp.InternalError(response.NewError(err))
			return
		}

		log.LogError(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	loginOutput := output.LoginOutputVO{
		Authorization: token,
	}

	resp.Ok(loginOutput)
}
