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

const invalidSecretKeySegmentString = "string"

func (instance InternalOtpService) GenerateOtpForApi(requester string, length int, interval int) (*dto.OtpRepositoryTimeBasedOtpResult, *dto.ApiErrorGeneric) {
	inputIssue := guardOtpSetupParameters(requester, length)

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

func guardOtpSetupParameters(requester string, length int) *dto.ApiErrorGeneric {
	if strings.ToUpper(requester) == strings.ToUpper(invalidSecretKeySegmentString) {
		return &dto.ApiErrorGeneric{
			HttpStatus: 400,
			Error:      errors.New("string is not a valid value for requester"),
		}
	}

	if length > 10 {
		return &dto.ApiErrorGeneric{
			HttpStatus: 400,
			Error:      errors.New("the length cannot be greater than 10"),
		}
	}

	return nil
}

func (instance InternalOtpService) GenerateOtp(requester string, length int, interval int) (*dto.OtpRepositoryTimeBasedOtpResult, error) {
	cleanedRequester, regExpError := cleanSecretSectionKey(requester)

	if regExpError != nil {
		return nil, regExpError
	}

	repoResult, exception := repositories.InternalOtp{}.GenerateTimeBasedOtp(cleanedRequester, length, interval)

	if exception != nil {
		return nil, exception
	}

	return &repoResult, nil
}

func (instance InternalOtpService) ValidateOtpForApi(requester string, length int, interval int, otp, referenceToken string) (bool, *dto.ApiErrorGeneric) {
	inputIssue := guardOtpSetupParameters(requester, length)

	if inputIssue != nil {
		return false, inputIssue
	}

	if strings.ToUpper(referenceToken) == strings.ToUpper(invalidSecretKeySegmentString) {
		return false, &dto.ApiErrorGeneric{
			HttpStatus: 400,
			Error:      errors.New("string is not a valid value for referenceToken"),
		}
	}

	isValid, exception := instance.ValidateOtp(requester, length, interval, otp, referenceToken)

	if exception != nil {
		return false, &dto.ApiErrorGeneric{
			HttpStatus: 500,
			Error:      exception,
		}
	}

	return isValid, nil
}

func (instance InternalOtpService) ValidateOtp(requester string, length int, interval int, otp, referenceToken string) (bool, error) {
	cleanedRequester, regExpError := cleanSecretSectionKey(requester)

	if regExpError != nil {
		return false, regExpError
	}

	cleanedReferenceToken, referenceTokenRegExpError := cleanSecretSectionKey(referenceToken)

	if referenceTokenRegExpError != nil {
		return false, referenceTokenRegExpError
	}

	isValid, exception := repositories.InternalOtp{}.ValidateTimeBasedOtp(cleanedRequester, length, interval, otp, cleanedReferenceToken)

	if exception != nil {
		return false, exception
	}

	return isValid, nil
}

func cleanSecretSectionKey(requester string) (string, error) {
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
