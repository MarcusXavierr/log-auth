package stringgenerator_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/MarcusXavierr/log-auth/internal/enum"
	"github.com/MarcusXavierr/log-auth/internal/login"
	"github.com/MarcusXavierr/log-auth/internal/stringgenerator"
)

const (
	shipmentIntelDataset = "shipment"
	freightIntelDataset  = "freight"
	fakeAccessToken      = "xyz.cringe.whatever"
	fakeEmail            = "leonardo.poliva@bol.com.br"
)

type mockLogin struct {
}

func (f mockLogin) GetResponseData() (login.Response, error) {
	return login.Response{Auth: login.Auth{AccessToken: fakeAccessToken}}, nil
}

func (f mockLogin) SendLoginRequest() (*http.Response, error) {
	panic("not implemented")
}

func (f mockLogin) GetCredentials() login.Credentials {
	return login.Credentials{Email: fakeEmail, Password: "secret"}
}

func (f mockLogin) GetDatasets() login.Datasets {
	return login.Datasets{ShipmentIntel: shipmentIntelDataset, FreightIntel: freightIntelDataset}
}

func TestGeneratePBILocalUrl(t *testing.T) {
	fakeLogin := mockLogin{}
	t.Run("Generate string for shipmentIntel", func(t *testing.T) {
		got, _ := stringgenerator.GeneratePBILocalUrl(fakeLogin, enum.ShipmentIntel)
		want := fmt.Sprintf("%s?token=%s&email=%s&cookiesHasBeenAccepted=true&datasetId=%s",
			stringgenerator.BasePBIUrl,
			fakeAccessToken,
			fakeEmail,
			shipmentIntelDataset,
		)

		validateString(got, want, t)
	})

	t.Run("Generate string for freightIntel", func(t *testing.T) {
		got, _ := stringgenerator.GeneratePBILocalUrl(fakeLogin, enum.FreightIntel)
		want := fmt.Sprintf("%s?token=%s&email=%s&cookiesHasBeenAccepted=true&datasetId=%s",
			stringgenerator.BasePBIUrl,
			fakeAccessToken,
			fakeEmail,
			freightIntelDataset,
		)

		validateString(got, want, t)
	})
}

func validateString(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("expected %s but got %s", want, got)
	}
}
