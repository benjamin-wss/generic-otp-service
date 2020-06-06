package services

import (
	"errors"
	"generic-otp-service/config"
	"generic-otp-service/constants"
	"generic-otp-service/dto"
	"generic-otp-service/models"
	"generic-otp-service/repositories"
	"regexp"
	"strings"
)

type InternalOtpService struct {
	OtpLogDbRepository repositories.IDbOtpLogRepository
}

const (
	otpRequestRowType                   = `ACQUIRE`
	otpRequestNoLoggedErrorMessage      = `otp request does not exist`
	otpRequestConsumedErrorMessage      = `otp has been invalidated by admin`
	otpVerificationRowType              = `VERIFY`
	otpVerificationConsumedErrorMessage = `otp was consumed earlier`
)

func (instance InternalOtpService) GenerateOtpForApi(requester string, length int, interval int) (*dto.OtpRepositoryTimeBasedOtpResult, *dto.ApiErrorGeneric) {
	inputIssue := guardOtpSetupParameters(length)

	if inputIssue != nil {
		return nil, inputIssue
	}

	result, exception := instance.GenerateOtp(requester, length, interval)

	if exception != nil {
		return nil, &dto.ApiErrorGeneric{
			HttpStatus: 500,
			Error:      exception,
		}
	}

	return result, nil
}

func guardOtpSetupParameters(length int) *dto.ApiErrorGeneric {
	if length > 10 {
		return &dto.ApiErrorGeneric{
			HttpStatus: 400,
			Error:      errors.New("the length cannot be greater than 10"),
		}
	}

	return nil
}

func (instance InternalOtpService) GenerateOtp(requester string, length int, otpLifespanInSecondsinterval int) (*dto.OtpRepositoryTimeBasedOtpResult, error) {
	computationResult, exception := repositories.InternalOtp{}.GenerateTimeBasedOtp(length, otpLifespanInSecondsinterval)

	if exception != nil {
		return nil, exception
	}

	if config.AppConfig.Otp.OtpRequestLoggingEnabled == true {
		_, dbError := instance.OtpLogDbRepository.Create(models.OtpLog{
			Type:                 "ACQUIRE",
			Requester:            requester,
			OtpLifespanInSeconds: otpLifespanInSecondsinterval,
			ExpiryInUnixTime:     computationResult.ExpiryInSeconds,
			Otp:                  &computationResult.Otp,
			ReferenceToken:       &computationResult.ReferenceToken,
			IsConsumed:           false,
		})

		if dbError != nil {
			return &computationResult, dbError
		}
	}

	return &computationResult, nil
}

func (instance InternalOtpService) ValidateOtpForApi(requester string, length int, interval int, otp, referenceToken string) (bool, *dto.ApiErrorGeneric) {
	inputIssue := guardOtpSetupParameters(length)

	if inputIssue != nil {
		return false, inputIssue
	}

	const invalidSecretKeySegmentString = "string"

	if strings.ToUpper(referenceToken) == strings.ToUpper(invalidSecretKeySegmentString) {
		return false, &dto.ApiErrorGeneric{
			HttpStatus: 400,
			Error:      errors.New("string is not a valid value for referenceToken"),
		}
	}

	isValid, exception := instance.ValidateOtp(requester, length, interval, otp, referenceToken)

	if exception != nil {
		customException := dto.ApiErrorGeneric{
			HttpStatus: 500,
			Error:      exception,
		}

		if exception.Error() == otpRequestNoLoggedErrorMessage {
			customException.HttpStatus = 404
			return false, &customException
		}

		if exception.Error() == otpRequestConsumedErrorMessage || exception.Error() == otpVerificationConsumedErrorMessage {
			customException.HttpStatus = 400
			return false, &customException
		}

		if exception.Error() == constants.GenericDbErrorMessages.ConflictingEntryDetected {
			customException.HttpStatus = 409
			return false, &customException
		}

		return false, &customException
	}

	return isValid, nil
}

func (instance InternalOtpService) ValidateOtp(requester string, length int, interval int, otp, referenceToken string) (bool, error) {
	cleanedReferenceToken, referenceTokenRegExpError := cleanSecretSectionKey(referenceToken)

	if config.AppConfig.Otp.OtpVerificationLoggingEnabled == true {
		verificationEntryRequest, verificationEntryRequestError := instance.OtpLogDbRepository.GetExistingEntry(
			otpVerificationRowType,
			requester,
			otp,
			cleanedReferenceToken,
		)

		if verificationEntryRequestError != nil {
			return false, verificationEntryRequestError
		}

		if verificationEntryRequest != nil && verificationEntryRequest.IsConsumed == true {
			return false, errors.New(otpVerificationConsumedErrorMessage)
		}
	}

	var existingEntry *models.OtpLog
	if config.AppConfig.Otp.OtpRequestLoggingEnabled == true {
		acquireEntryResult, acquireEntryError := instance.OtpLogDbRepository.GetExistingEntry(
			otpRequestRowType,
			requester,
			otp,
			cleanedReferenceToken,
		)

		if acquireEntryError != nil {
			return false, acquireEntryError
		}

		if acquireEntryResult == nil {
			return false, errors.New(otpRequestNoLoggedErrorMessage)
		}

		if acquireEntryResult.IsConsumed == true {
			return false, errors.New(otpRequestConsumedErrorMessage)
		}

		existingEntry = acquireEntryResult
	}

	if referenceTokenRegExpError != nil {
		return false, referenceTokenRegExpError
	}

	isValid, exception := repositories.InternalOtp{}.ValidateTimeBasedOtp(length, interval, otp, cleanedReferenceToken)

	if exception != nil {
		return false, exception
	}

	if isValid == false {
		return false, nil
	}

	if config.AppConfig.Otp.OtpRequestLoggingEnabled {
		logEntry := models.OtpLog{
			Type:                 otpVerificationRowType,
			Requester:            requester,
			OtpLifespanInSeconds: interval,
			ExpiryInUnixTime:     0,
			Otp:                  &otp,
			ReferenceToken:       &referenceToken,
			IsConsumed:           true,
		}

		if existingEntry != nil && existingEntry.ExpiryInUnixTime > 0 {
			logEntry.ExpiryInUnixTime = existingEntry.ExpiryInUnixTime
		}

		_, loggingError := instance.OtpLogDbRepository.CreateConsumedEntry(logEntry)

		if loggingError != nil {
			return false, loggingError
		}
	}

	return isValid, nil
}

func cleanSecretSectionKey(secretSection string) (string, error) {
	regularExpressionString := constants.RequesterRegularExpression
	regularExpressionProcessor, err := regexp.Compile(regularExpressionString)

	if err != nil {
		regExCompileError := errors.New("problem compiling regular expression for validation")
		return "", regExCompileError
	}

	processedString := regularExpressionProcessor.ReplaceAllString(secretSection, "")
	upperCaseString := strings.ToUpper(processedString)

	return upperCaseString, nil
}
