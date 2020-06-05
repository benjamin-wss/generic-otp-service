package enums

type OtpEnvironmentVariableKeyEnum int

const (
	OtpSecret                     OtpEnvironmentVariableKeyEnum = 1
	OtpRequestLoggingEnabled      OtpEnvironmentVariableKeyEnum = 2
	OtpVerificationLoggingEnabled OtpEnvironmentVariableKeyEnum = 3
)

func (keys OtpEnvironmentVariableKeyEnum) ToString() string {
	switch keys {
	case OtpSecret:
		return "OTP_SECRET"
	case OtpRequestLoggingEnabled:
		return "OTP_REQUEST_LOGGING_ENABLED"
	case OtpVerificationLoggingEnabled:
		return "OTP_VERIFICATION_LOGGING_ENABLED"
	default:
		panic("OtpEnvironmentVariableKeyEnum : Unexpected value requested/assigned.")
	}
}
