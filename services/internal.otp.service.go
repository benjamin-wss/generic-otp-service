package services

import (
	"errors"
	"generic-otp-service/constants"
	"generic-otp-service/dto"
	"generic-otp-service/repositories"
	"regexp"
	"strings"
)

type InternalOtpService struct{}

func (instance InternalOtpService) GenerateOtpForApi(requester string, length int, interval int) (*dto.OtpRepositoryTimeBasedOtpResult, *dto.ApiErrorGeneric) {
	if strings.ToUpper(requester) == "STRING" {
		return nil, &dto.ApiErrorGeneric{
			HttpStatus: 400,
			Error:      errors.New("string is not a valid value for requester"),
		}
	}

	if length > 10 {
		return nil, &dto.ApiErrorGeneric{
			HttpStatus: 400,
			Error:      errors.New("the length cannot be greater than 10"),
		}
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

func (instance InternalOtpService) GenerateOtp(requester string, length int, interval int) (*dto.OtpRepositoryTimeBasedOtpResult, error) {
	cleanedRequester, regExpError := cleanRequesterString(requester)

	if regExpError != nil {
		return nil, regExpError
	}

	repoResult, exception := repositories.InternalOtp{}.GenerateTimeBasedOtp(cleanedRequester, length, interval)

	if exception != nil {
		return nil, exception
	}

	return &repoResult, nil
}

func (instance InternalOtpService) ValidateOtp(requester string, length int, interval int, otp string) (bool, error) {
	cleanedRequester, regExpError := cleanRequesterString(requester)

	if regExpError != nil {
		return false, regExpError
	}

	isValid, exception := repositories.InternalOtp{}.ValidateTimeBasedOtp(cleanedRequester, length, interval, otp)

	if exception != nil {
		return false, exception
	}

	return isValid, nil
}

func cleanRequesterString(requester string) (string, error) {
	regularExpressionString := constants.RequesterRegularExpression
	regularExpressionProcessor, err := regexp.Compile(regularExpressionString)

	if err != nil {
		regExCompileError := errors.New("problem compiling regular expression for validation")
		return "", regExCompileError
	}

	processedString := regularExpressionProcessor.ReplaceAllString(requester, "")
	upperCaseString := strings.ToUpper(processedString)

	return upperCaseString, nil
}
