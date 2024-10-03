package infisical

import (
	"os"

	api "github.com/infisical/go-sdk/packages/api/auth"
	"github.com/infisical/go-sdk/packages/util"
)

type KubernetesAuthLoginOptions struct {
	IdentityID              string
	ServiceAccountTokenPath string
}

// func epochTime() time.Time { return time.Unix(0, 0) }

type AuthInterface interface {
	SetAccessToken(accessToken string)
	UniversalAuthLogin(clientID string, clientSecret string) (credential MachineIdentityCredential, err error)
	OidcAuthLogin(identityId string, jwt string) (credential MachineIdentityCredential, err error)
}

type Auth struct {
	client *InfisicalClient
}

func (a *Auth) SetAccessToken(accessToken string) {
	a.client.setAccessToken(accessToken, util.ACCESS_TOKEN)
}

func (a *Auth) UniversalAuthLogin(clientID string, clientSecret string) (credential MachineIdentityCredential, err error) {

	if clientID == "" {
		clientID = os.Getenv(util.INFISICAL_UNIVERSAL_AUTH_CLIENT_ID_ENV_NAME)
	}
	if clientSecret == "" {
		clientSecret = os.Getenv(util.INFISICAL_UNIVERSAL_AUTH_CLIENT_SECRET_ENV_NAME)
	}

	credential, err = api.CallUniversalAuthLogin(a.client.httpClient, api.UniversalAuthLoginRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	})

	if err != nil {
		return MachineIdentityCredential{}, err
	}

	a.client.setAccessToken(credential.AccessToken, util.UNIVERSAL_AUTH)
	return credential, nil

}

func (a *Auth) OidcAuthLogin(identityId string, jwt string) (credential MachineIdentityCredential, err error) {
	if identityId == "" {
		identityId = os.Getenv(util.INFISICAL_OIDC_AUTH_IDENTITY_ID_ENV_NAME)
	}

	credential, err = api.CallOidcAuthLogin(a.client.httpClient, api.OidcAuthLoginRequest{
		IdentityID: identityId,
		JWT:        jwt,
	})

	if err != nil {
		return MachineIdentityCredential{}, err
	}

	a.client.setAccessToken(credential.AccessToken, util.OIDC_AUTH)
	return credential, nil

}

func NewAuth(client *InfisicalClient) AuthInterface {
	return &Auth{client: client}
}
