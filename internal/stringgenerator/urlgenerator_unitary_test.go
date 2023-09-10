package stringgenerator

import (
	"fmt"
	"testing"

	"github.com/MarcusXavierr/log-auth/internal/constants"
	"github.com/MarcusXavierr/log-auth/internal/enum"
	"github.com/MarcusXavierr/log-auth/internal/login"
	"github.com/spf13/viper"
)

const (
	shipmentIntelDataset = "shipment"
	freightIntelDataset  = "freight"
	fakeAccessToken      = "xyz.cringe.whatever"
	fakeEmail            = "leonardo.poliva@bol.com.br"
)

var urlGeneratorMock = mountUrlGeneratorMock()

func TestMountQueryString(t *testing.T) {

	t.Run("Mount query String for Shipment Intel", func(t *testing.T) {
		urlGeneratorMock.desiredPlatform = enum.DesiredPlatform(enum.ShipmentIntel)
		got := urlGeneratorMock.mountQueryString()
		want := fmt.Sprintf(
			"?token=%s&email=%s&cookiesHasBeenAccepted=true&datasetId=%s",
			fakeAccessToken,
			fakeEmail,
			shipmentIntelDataset,
		)
		validateString(t, got, want)
	})

	t.Run("Mount query String for FreightIntel", func(t *testing.T) {
		urlGeneratorMock.desiredPlatform = enum.DesiredPlatform(enum.FreightIntel)
		got := urlGeneratorMock.mountQueryString()
		want := fmt.Sprintf(
			"?token=%s&email=%s&cookiesHasBeenAccepted=true&datasetId=%s",
			fakeAccessToken,
			fakeEmail,
			freightIntelDataset,
		)

		validateString(t, got, want)
	})

	t.Run("Mount query string for a product without dataset id", func(t *testing.T) {
		urlGeneratorMock.desiredPlatform = enum.Others
		got := urlGeneratorMock.mountQueryString()
		want := fmt.Sprintf(
			"?token=%s&email=%s&cookiesHasBeenAccepted=true",
			fakeAccessToken,
			fakeEmail,
		)

		validateString(t, got, want)
	})
}

func validateString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %s but got %s", want, got)
	}
}

func mountUrlGeneratorMock() urlGeneratorData {
	viper.SetDefault(constants.EMAIL, fakeEmail)
	viper.SetDefault(constants.SHIPMENT_INTEL_DATASET_ID, shipmentIntelDataset)
	viper.SetDefault(constants.FREIGHT_INTEL_DATASET_ID, freightIntelDataset)

	return urlGeneratorData{
		login:         login.MountLoginData(),
		loginResponse: login.Response{Auth: login.Auth{AccessToken: fakeAccessToken}},
	}
}
