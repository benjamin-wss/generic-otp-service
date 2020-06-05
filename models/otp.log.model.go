package models

type OtpLog struct {
	Base
	Requester            string `gorm:"not null"`
	OtpLifespanInSeconds int    `gorm:"column:otp_lifespan_in_seconds;not null"`
	Otp                  string
	ReferenceToken       string `gorm:"column:reference_token"`
}
