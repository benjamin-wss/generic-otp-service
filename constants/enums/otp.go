package enums

type OtpEnvironmentVariableKeyEnum string

const (
	OtpSecret OtpEnvironmentVariableKeyEnum = "OTP_SECRET"
)

func (keys OtpEnvironmentVariableKeyEnum) ToString() string {
	switch keys {
	case OtpSecret:
		return "OTP_SECRET"
	default:
		panic("OtpEnvironmentVariableKeyEnum : Unexpected value requested/assigned.")
	}
}
