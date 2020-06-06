package types

type OtpConfig struct {
	OtpSecret                     string
	OtpRequestLoggingEnabled      bool
	OtpVerificationLoggingEnabled bool
}
