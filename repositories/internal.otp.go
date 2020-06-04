package repositories

import (
	"fmt"
	"generic-otp-service/config"
	"generic-otp-service/dto"
	"github.com/xlzd/gotp"
	"strings"
	"time"
)

type InternalOtp struct{}

func (instance InternalOtp) GenerateTimeBasedOtp(length int, interval int) (dto.OtpRepositoryTimeBasedOtpResult, error) {
	var payload dto.OtpRepositoryTimeBasedOtpResult

	randomSecret := gotp.RandomSecret(16)
	requesterInUpperCase := getSecret(randomSecret)
	lib := gotp.NewTOTP(requesterInUpperCase, length, interval, nil)

	otpValue, otpTimeoutInSeconds := lib.NowWithExpiration()

	payload = dto.OtpRepositoryTimeBasedOtpResult{
		ReferenceToken:  randomSecret,
		Otp:             otpValue,
		ExpiryInSeconds: otpTimeoutInSeconds,
	}

	return payload, nil
}

func (instance InternalOtp) ValidateTimeBasedOtp(length, interval int, otp, referenceToken string) (bool, error) {
	requesterInUpperCase := getSecret(referenceToken)
	lib := gotp.NewTOTP(requesterInUpperCase, length, interval, nil)
	isValid := lib.Verify(otp, int(time.Now().Unix()))

	return isValid, nil
}

func getSecret(prefix string) string {
	baseResult := fmt.Sprintf("%s%s", prefix, config.AppConfig.Otp.Secret)
	upperCaseResult := strings.ToUpper(baseResult)

	return upperCaseResult
}
