package services

import (
	"errors"
	"generic-otp-service/constants"
	"generic-otp-service/repositories"
	"regexp"
)

type InternalOtpService struct{}

func (instance InternalOtpService) GenerateOtp(requester string, length int, interval int) (string, int64, error) {
	cleanedRequester, regExpError := cleanRequesterString(requester)

	if regExpError != nil {
		return "", 0, regExpError
	}

	otpNumber, otpTimeoutInSeconds, exception := repositories.InternalOtp{}.GenerateTimeBasedOtp(cleanedRequester, length, interval)

	if exception != nil {
		return "", 0, exception
	}

	return otpNumber, otpTimeoutInSeconds, nil
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

	return processedString, nil
}
