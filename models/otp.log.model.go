package models

type OtpLog struct {
	Base
	Type                 string `gorm:"not null"`
	Requester            string `gorm:"not null"`
	OtpLifespanInSeconds int    `gorm:"column:otp_lifespan_in_seconds;not null"`
	ExpiryInUnixTime     int64  `gorm:"column:expiry_in_unix_time;not null"`
	Otp                  *string
	ReferenceToken       *string `gorm:"column:reference_token" sql:"index"`
	IsConsumed           bool    `gorm:"column:is_consumed;not null;default:false"`
}
