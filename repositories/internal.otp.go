package repositories

import (
	"errors"
	"fmt"
	"generic-otp-service/constants"
	"github.com/xlzd/gotp"
	"regexp"
	"strings"
	"time"
)

type InternalOtp struct{}

func (instance InternalOtp) GenerateTimeBasedOtp(requester string, length int, interval int) (string, int64, error) {
	_, exception := checkIsRequesterInputValid(requester)

	if exception != nil {
		return "", 0, exception
	}

	requesterInUpperCase := getSecret(requester)
	otp := gotp.NewTOTP(requesterInUpperCase, length, interval, nil)

	otpValue, otpTimeoutInSeconds := otp.NowWithExpiration()

	return otpValue, otpTimeoutInSeconds, nil
}

func (instance InternalOtp) ValidateTimeBasedOtp(requester string, length int, interval int, otp string) (bool, error) {
	_, exception := checkIsRequesterInputValid(requester)

	if exception != nil {
		return false, exception
	}

	requesterInUpperCase := getSecret(requester)
	otpUtil := gotp.NewTOTP(requesterInUpperCase, length, interval, nil)
	isValid := otpUtil.Verify(otp, int(time.Now().Unix()))

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

func getSecret(postfix string) string {
	baseResult := fmt.Sprintf("%s%s", "4S62BZNFXXSZLCRO", postfix)
	upperCaseResult := strings.ToUpper(baseResult)

	return upperCaseResult
}
