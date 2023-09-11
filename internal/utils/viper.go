package utils

import (
	"net/http"
	"os"

	"github.com/MarcusXavierr/log-auth/internal/constants"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func InitializeViper() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "Cannot find UserHomeDir")
	}

	viper.SetConfigName(".log-auth")
	viper.AddConfigPath(home)
	viper.SetDefault(constants.ENDPOINT, "http://api-auth.homol.logcomex.io/api/login")
	viper.SetDefault(constants.REQUEST_METHOD, http.MethodPost)
	viper.SetDefault(constants.SHIPMENT_INTEL_DATASET_ID, "62c3d669-c1ed-499e-bcd5-32ff1108a814")
	viper.SetDefault(constants.FREIGHT_INTEL_DATASET_ID, "4765cb7a-b9b3-4f04-a378-a34045e31836")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.Wrap(err, "ConfigFileNotFoundError")
		}
		return err
	}

	return nil
}
