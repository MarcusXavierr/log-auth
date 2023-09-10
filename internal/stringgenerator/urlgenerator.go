package stringgenerator

import (
	"fmt"

	"github.com/MarcusXavierr/log-auth/internal/enum"

	"github.com/MarcusXavierr/log-auth/internal/login"
)

const (
	basePBIUrl = "http://localhost:3000/bi/dashboard/"
)

type Struct struct {
	login           login.ILogin
	loginResponse   login.Response
	desiredPlatform enum.DesiredPlatform
}

func GeneratePBILocalUrl(loginData login.ILogin, desiredPlatform enum.DesiredPlatform) (string, error) {
	response, err := loginData.GetResponseData()

	if err != nil {
		return "", err
	}

	tmpName := Struct{login: loginData, loginResponse: response, desiredPlatform: desiredPlatform}

	return basePBIUrl + tmpName.mountQueryString(), nil
}

// func GeneratePBILocalUrl(struc Struct) string {
// 	return basePBIUrl + struc.mountQueryString()
// }

func (s Struct) mountQueryString() string {
	queryString := fmt.Sprintf(
		"?token=%s&email=%s&cookiesHasBeenAccepted=true",
		s.loginResponse.Auth.AccessToken,
		s.login.GetCredentials().Email,
	)

	return queryString + s.getRightDataset()
}

func (s Struct) getRightDataset() string {
	baseQueryString := "&datasetId="

	switch s.desiredPlatform {
	case enum.ShipmentIntel:
		return baseQueryString + s.login.GetDatasets().ShipmentIntel
	case enum.FreightIntel:
		return baseQueryString + s.login.GetDatasets().FreightIntel
	}

	return ""
}
