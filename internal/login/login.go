// Package login provides functionalities for authenticating users on the
// platform and managing login-related data.
//
// This package allows users to log in to the platform by interacting with an external API to obtain authentication tokens.
// It also offers the ability to store and retrieve user login information, such as access tokens and profile data.

package login

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// ILogin represents the interface for the struct that stores user information required for login and also performs the login.
type ILogin interface {
	GetResponseData() (Response, error)
	SendLoginRequest() (*http.Response, error)
	GetCredentials() Credentials
	GetDatasets() Datasets
}

// Login This struct stores the essential information required for user login and
// their corresponding settings
type Login struct {
	credentials Credentials
	datasets    Datasets
	endpoint    string
	method      string
}

type Datasets struct {
	ShipmentIntel string
	FreightIntel  string
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Auth Auth `json:"auth"`
}

type Auth struct {
	AccessToken string `json:"access_token"`
}

// GetResponseData This function sends a request to the endpoint specified in the
// Login struct and parses the response into the Response struct
func (l Login) GetResponseData() (Response, error) {
	httpResponse, err := l.SendLoginRequest()
	if err != nil {
		return Response{}, err
	}
	defer httpResponse.Body.Close()

	buffer, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return Response{}, errors.Wrap(err, "ioutil.ReadAll")
	}

	response := Response{}
	if err := json.Unmarshal(buffer, &response); err != nil {
		return Response{}, errors.Wrap(err, "json.Unmarshal")
	}

	return response, nil
}

// SendLoginRequest This function sends a request to the endpoint specified in
// the struct along with the credentials it contains, and then returns an HTTP
// response
func (l Login) SendLoginRequest() (*http.Response, error) {
	payload, err := json.Marshal(l.credentials)

	if err != nil {
		return nil, errors.Wrap(err, "create json")
	}

	req, err := http.NewRequest(l.method, l.endpoint, bytes.NewBuffer(payload))

	if err != nil {
		return nil, errors.Wrap(err, "create request")
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	response, err := httpClient.Do(req)

	if err != nil {
		return nil, errors.Wrap(err, "error on request")
	}

	return response, nil
}

// GetCredentials returns the user Credentials that are used to log in
func (l Login) GetCredentials() Credentials {
	return l.credentials
}

func (l Login) GetDatasets() Datasets {
	return l.datasets
}
