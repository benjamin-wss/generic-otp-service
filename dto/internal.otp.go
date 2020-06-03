package dto

type ApiInputBasicOtp struct {
	Requester string `json:"requester"`
	Length    int    `json:"length"`
	Interval  int    `json:"interval"`
}

type ApiResultBasicOtp struct {
	Otp             string `json:"otp"`
	ExpiryInSeconds int64  `json:"expiryInSeconds"`
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
