package repositories

import (
	"generic-otp-service/models"
	"github.com/jinzhu/gorm"
)

type IDbOtpLogRepository interface {
	Create(input models.OtpLog) (models.OtpLog, error)
	GetExistingEntry(requester string, otp string, referenceToken string) (*models.OtpLog, error)
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

func (instance DbOtpLogRepository) GetExistingEntry(requester string, otp string, referenceToken string) (*models.OtpLog, error) {
	var payload models.OtpLog

	dbResponse := instance.dbConnection.Where(&models.OtpLog{
		Requester:      requester,
		Otp:            &otp,
		ReferenceToken: &referenceToken,
	}).First(&payload)

	if dbResponse.Error != nil && gorm.IsRecordNotFoundError(dbResponse.Error) == true {
		return nil, nil
	}

	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}

	return &payload, nil
}
