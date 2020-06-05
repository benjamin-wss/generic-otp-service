package repositories

import (
	"generic-otp-service/models"
	"github.com/jinzhu/gorm"
)

type IDbOtpLogRepository interface {
	Create(input models.OtpLog) (models.OtpLog, error)
}

func GetDbOtpLogRepository(dbConnection *gorm.DB) IDbOtpLogRepository {
	return &DbOtpLogRepository{dbConnection}
}

type DbOtpLogRepository struct {
	dbConnection *gorm.DB
}

func (instance DbOtpLogRepository) Create(input models.OtpLog) (models.OtpLog, error) {
	clonedInput := input

	if result := instance.dbConnection.Create(&clonedInput); result.Error != nil {
		return clonedInput, result.Error
	}

	return clonedInput, nil
}
