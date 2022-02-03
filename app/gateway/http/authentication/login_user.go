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
	c.log.SetRequestIDFromContext(r.Context())

	resp := response.NewResponse(w)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.log.LogWarn(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}
	defer r.Body.Close()

	c.log.LogDebug(operation, "unmarshal the data to a internal object")
	var loginData input.LoginVO
	err = json.Unmarshal(reqBody, &loginData)
	if err != nil {
		c.log.LogWarn(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	c.log.LogDebug(operation, "validating the input data")
	err = validations.ValidateLoginInputData(loginData)
	if err != nil {
		c.log.LogWarn(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	account := account.Account{
		CPF:    loginData.CPF.TrimCPF(),
		Secret: types.Password(loginData.Secret),
	}

	c.log.LogDebug(operation, "trying to log in the system")
	token, err := c.usecase.LoginUser(r.Context(), account)
	if err != nil {
		if errors.Is(err, customError.ErrorAccountTokenGeneration) {
			c.log.LogWarn(operation, err.Error())
			resp.InternalError(response.NewError(err))
			return
		}

		c.log.LogError(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	loginOutput := output.LoginOutputVO{
		Authorization: token,
	}
	c.log.LogInfo(operation, "account logged successfully")
	resp.Ok(loginOutput)
}
