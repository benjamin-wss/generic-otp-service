package dto

type OtpRepositoryTimeBasedOtpResult struct {
	ReferenceToken  string `json:"referenceToken"`
	Otp             string `json:"otp"`
	ExpiryInSeconds int64  `json:"expiryInSeconds"`
}
