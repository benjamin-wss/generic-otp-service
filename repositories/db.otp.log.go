package repositories

import (
	"errors"
	"generic-otp-service/constants"
	"generic-otp-service/models"
	"github.com/jinzhu/gorm"
)

type IDbOtpLogRepository interface {
	Create(input models.OtpLog) (models.OtpLog, error)
	GetExistingEntry(requestType string, requester string, otp string, referenceToken string) (*models.OtpLog, error)
	CreateConsumedEntry(input models.OtpLog) (*models.OtpLog, error)
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

func (instance DbOtpLogRepository) GetExistingEntry(requestType string, requester string, otp string, referenceToken string) (*models.OtpLog, error) {
	var payload models.OtpLog

	dbResponse := instance.dbConnection.Where(&models.OtpLog{
		Type:           requestType,
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

func (instance DbOtpLogRepository) CreateConsumedEntry(input models.OtpLog) (*models.OtpLog, error) {
	dbConnection := instance.dbConnection
	transaction := dbConnection.Begin()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	if transactionInitError := transaction.Error; transactionInitError != nil {
		return nil, transactionInitError
	}

	payload := input
	newRecordResponse := transaction.Create(&payload)

	if newRecordResponse.Error != nil {
		transaction.Rollback()
		return nil, newRecordResponse.Error
	}

	var numberOfRows int
	dbCountResult := transaction.Model(&models.OtpLog{}).Where(&models.OtpLog{
		Type:           payload.Type,
		Requester:      payload.Requester,
		Otp:            payload.Otp,
		ReferenceToken: payload.ReferenceToken,
		IsConsumed:     payload.IsConsumed,
	}).Count(&numberOfRows)

	if dbCountResult.Error != nil && gorm.IsRecordNotFoundError(dbCountResult.Error) == false {
		transaction.Rollback()
		return nil, dbCountResult.Error
	}

	if numberOfRows > 1 {
		transaction.Rollback()
		return nil, errors.New(constants.GenericDbErrorMessages.ConflictingEntryDetected)
	}

	transactionResponse := transaction.Commit()

	return &payload, transactionResponse.Error
}
