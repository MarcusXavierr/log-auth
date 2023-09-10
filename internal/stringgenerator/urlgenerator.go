package stringgenerator

import (
	"fmt"

	"github.com/MarcusXavierr/log-auth/internal/enum"

	"github.com/MarcusXavierr/log-auth/internal/login"
)

const (
	BasePBIUrl = "http://localhost:3000/bi/dashboard/"
)

// urlGeneratorData aims to store data from user login, its desiredPlatform and also the login response which stores the access token
type urlGeneratorData struct {
	login           login.Login
	loginResponse   login.Response
	desiredPlatform enum.DesiredPlatform
}

func GeneratePBILocalUrl(loginData login.Login, desiredPlatform enum.DesiredPlatform) (string, error) {
	response, err := loginData.GetResponseData()

	if err != nil {
		return "", err
	}

	urlData := urlGeneratorData{login: loginData, loginResponse: response, desiredPlatform: desiredPlatform}

	return BasePBIUrl + urlData.mountQueryString(), nil
}

func (u urlGeneratorData) mountQueryString() string {
	queryString := fmt.Sprintf(
		"?token=%s&email=%s&cookiesHasBeenAccepted=true",
		u.loginResponse.Auth.AccessToken,
		u.login.GetCredentials().Email,
	)

	return queryString + u.getRightDataset()
}

func (u urlGeneratorData) getRightDataset() string {
	baseQueryString := "&datasetId="

	switch u.desiredPlatform {
	case enum.ShipmentIntel:
		return baseQueryString + u.login.GetDatasets().ShipmentIntel
	case enum.FreightIntel:
		return baseQueryString + u.login.GetDatasets().FreightIntel
	}

	return ""
}
