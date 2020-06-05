package config

import (
	"generic-otp-service/config/types"
	"generic-otp-service/constants"
	"generic-otp-service/constants/enums"
	"generic-otp-service/utilities"
)

var AppConfig *types.ApplicationConfig

var environmentVariableUtilities = utilities.EnvironmentVariableUtilities{}

func init() {
	AppConfig = &types.ApplicationConfig{
		Otp: setupOtpConfig(),
		Db:  setupDbConfig(),
	}
}

func setupOtpConfig() types.OtpConfig {
	return types.OtpConfig{
		Secret: environmentVariableUtilities.GetEnvironmentVariableAsString(enums.OtpSecret.ToString()),
	}
}

func setupDbConfig() types.DbConfig {
	config := types.DbConfig{
		DbType:               environmentVariableUtilities.GetEnvironmentVariableAsString(enums.DbType.ToString()),
		DbHost:               environmentVariableUtilities.GetEnvironmentVariableAsString(enums.DbHost.ToString()),
		DbPort:               environmentVariableUtilities.GetEnvironmentVariableAsInteger(enums.DbPort.ToString()),
		DbName:               environmentVariableUtilities.GetEnvironmentVariableAsString(enums.DbName.ToString()),
		DbUser:               environmentVariableUtilities.GetEnvironmentVariableAsString(enums.DbUser.ToString()),
		DbPassword:           environmentVariableUtilities.GetEnvironmentVariableAsString(enums.DbPassword.ToString()),
		DbPostgresSslSetting: "",
	}

	if config.DbUser == constants.SupportedDbTypes.Postgres {
		config.DbPostgresSslSetting = environmentVariableUtilities.GetEnvironmentVariableAsString(enums.DbPostgresSslSetting.ToString())
	}

	return config
}
