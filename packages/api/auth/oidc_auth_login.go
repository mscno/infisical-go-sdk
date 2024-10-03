package api

import (
	"github.com/go-resty/resty/v2"
	"github.com/mscno/infisical-go-sdk/packages/errors"
)

const callOidcAuthLoginOperation = "CallOidcAuthLogin"

func CallOidcAuthLogin(httpClient *resty.Client, request OidcAuthLoginRequest) (credential MachineIdentityAuthLoginResponse, e error) {
	var responseData MachineIdentityAuthLoginResponse

	response, err := httpClient.R().
		SetResult(&responseData).
		SetBody(request).
		Post("/v1/auth/oidc-auth/login")

	if err != nil {
		return MachineIdentityAuthLoginResponse{}, errors.NewRequestError(callOidcAuthLoginOperation, err)
	}

	if response.IsError() {
		return MachineIdentityAuthLoginResponse{}, errors.NewAPIError(callOidcAuthLoginOperation, response)
	}

	return responseData, nil
}
