package config

import (
	"generic-otp-service/config/types"
	"generic-otp-service/constants/enums"
	"generic-otp-service/utilities"
)

var AppConfig *types.ApplicationConfig

var environmentVariableUtilities = utilities.EnvironmentVariableUtilities{}

func init() {
	AppConfig = &types.ApplicationConfig{
		Otp: setupOtpConfig(),
	}
}

func setupOtpConfig() types.OtpConfig {
	return types.OtpConfig{
		Secret: environmentVariableUtilities.GetEnvironmentVariableAsString(enums.OtpSecret.ToString()),
	}
}
