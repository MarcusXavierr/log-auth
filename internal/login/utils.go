package login

import (
	"github.com/MarcusXavierr/log-auth/internal/constants"
	"github.com/spf13/viper"
)

func MountLoginData() loginData {
	email := viper.GetString(constants.EMAIL)
	password := viper.GetString(constants.PASSWORD)
	endpoint := viper.GetString(constants.ENDPOINT)
	requestMethod := viper.GetString(constants.REQUEST_METHOD)
	shipmentIntel := viper.GetString(constants.SHIPMENT_INTEL_DATASET_ID)
	freightIntel := viper.GetString(constants.FREIGHT_INTEL_DATASET_ID)

	data := loginData{
		credentials: Credentials{Email: email, Password: password},
		datasets:    Datasets{ShipmentIntel: shipmentIntel, FreightIntel: freightIntel},
		endpoint:    endpoint,
		method:      requestMethod,
	}

	return data
}
