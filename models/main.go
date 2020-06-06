package models

import (
	"fmt"
	"generic-otp-service/config"
	configTypes "generic-otp-service/config/types"
	"generic-otp-service/constants"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DbPrimary *gorm.DB

func ConnectPrimaryDatabase() {
	dbConfig := config.AppConfig.Db
	connectionString := deriveConnectionStringFromDbConfig(&dbConfig)

	database, err := gorm.Open(dbConfig.DbType, connectionString)

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&OtpLog{})

	if config.AppConfig.Gin.Mode == nil {
		database.LogMode(true)
	}

	DbPrimary = database
}

func deriveConnectionStringFromDbConfig(dbConfig *configTypes.DbConfig) string {
	if dbConfig.DbType != constants.SupportedDbTypes.Postgres {
		panic("Only postgres db supported for now")
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		dbConfig.DbHost,
		dbConfig.DbPort,
		dbConfig.DbUser,
		dbConfig.DbName,
		dbConfig.DbPassword,
		dbConfig.DbPostgresSslSetting)

	return connectionString
}
