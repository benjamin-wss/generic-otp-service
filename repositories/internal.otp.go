package repositories

import (
	"errors"
	"fmt"
	"generic-otp-service/config"
	"generic-otp-service/constants"
	"generic-otp-service/dto"
	"github.com/xlzd/gotp"
	"regexp"
	"strings"
	"time"
)

type InternalOtp struct{}

func (instance InternalOtp) GenerateTimeBasedOtp(requester string, length int, interval int) (dto.OtpRepositoryTimeBasedOtpResult, error) {
	_, exception := checkIsRequesterInputValid(requester)
	var payload dto.OtpRepositoryTimeBasedOtpResult

	if exception != nil {
		return payload, exception
	}

	randomSecret := gotp.RandomSecret(16)
	requesterInUpperCase := getSecret(randomSecret, requester)
	lib := gotp.NewTOTP(requesterInUpperCase, length, interval, nil)

	otpValue, otpTimeoutInSeconds := lib.NowWithExpiration()

	payload = dto.OtpRepositoryTimeBasedOtpResult{
		ReferenceToken:  randomSecret,
		Otp:             otpValue,
		ExpiryInSeconds: otpTimeoutInSeconds,
	}

	return payload, nil
}

func (instance InternalOtp) ValidateTimeBasedOtp(requester string, length int, interval int, otp, referenceToken string) (bool, error) {
	_, exception := checkIsRequesterInputValid(requester)

	if exception != nil {
		return false, exception
	}

	requesterInUpperCase := getSecret(referenceToken, requester)
	lib := gotp.NewTOTP(requesterInUpperCase, length, interval, nil)
	isValid := lib.Verify(otp, int(time.Now().Unix()))

	return isValid, nil
}

func checkIsRequesterInputValid(requester string) (bool, error) {
	regularExpressionString := constants.RequesterRegularExpression
	regularExpressionProcessor, err := regexp.Compile(regularExpressionString)

	if err != nil {
		regExCompileError := errors.New("problem compiling regular expression for validation")
		return false, regExCompileError
	}

	processedString := regularExpressionProcessor.ReplaceAllString(requester, "")

	if len(processedString) < len(requester) {
		requesterMismatchError := errors.New(fmt.Sprintf("requester string has issues, it must conform to the %s regular expression", regularExpressionString))
		return false, requesterMismatchError
	}

	return true, nil
}

func getSecret(prefix, postfix string) string {
	baseResult := fmt.Sprintf("%s%s%s", prefix, config.AppConfig.Otp.Secret, postfix)
	upperCaseResult := strings.ToUpper(baseResult)

	return upperCaseResult
}
