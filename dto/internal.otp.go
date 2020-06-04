package dto

type ApiInputBasicOtp struct {
	Requester            string `json:"requester" example:"jim@starfleet.com"`
	Length               int    `json:"length"`
	OtpLifespanInSeconds int    `json:"otpLifespanInSeconds"`
}

type ApiInputValidateBasicOtp struct {
	Requester string `json:"requester"`
	Length    int    `json:"length"`
	Interval  int    `json:"interval"`
	Otp       string `json:"otp"`
}

type ApiResultValidateBasicOtp struct {
	IsValid bool                     `json:"isValid"`
	Input   ApiInputValidateBasicOtp `json:"input"`
}
