package api

import (
	"github.com/go-resty/resty/v2"
	"github.com/infisical/go-sdk/packages/errors"
)

const callKubernetesAuthLoginOperation = "CallKubernetesAuthLogin"

func CallKubernetesAuthLogin(httpClient *resty.Client, request KubernetesAuthLoginRequest) (credential MachineIdentityAuthLoginResponse, e error) {
	var responseData MachineIdentityAuthLoginResponse

	response, err := httpClient.R().
		SetResult(&responseData).
		SetBody(request).
		Post("/v1/auth/kubernetes-auth/login")

	if err != nil {
		return MachineIdentityAuthLoginResponse{}, errors.NewRequestError(callKubernetesAuthLoginOperation, err)
	}

	if response.IsError() {
		return MachineIdentityAuthLoginResponse{}, errors.NewAPIError(callKubernetesAuthLoginOperation, response)
	}

	return responseData, nil
}
