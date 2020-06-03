package controllers

import (
	"generic-otp-service/dto"
	"generic-otp-service/services"
	"generic-otp-service/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InternalOtpController struct {
}

// GenerateOtpNumber godoc
// @Summary Generates T.O.T.P. number.
// @Description Generates T.O.T.P. number. Read more: https://en.wikipedia.org/wiki/Time-based_One-time_Password_algorithm
// @Tags Bambu
// @Accept  json
// @Produce  json
// @Param payload body dto.ApiInputBasicOtp true "Payload to generate T.O.T.P."
// @Success 200 {object} dto.ApiResultBasicOtp
// @Failure 500 {object} dto.HttpError
// @Router /api/internal/v1/acquire [post]
func (instance InternalOtpController) GenerateOtpNumber(context *gin.Context) {
	var input dto.ApiInputBasicOtp

	if bodyParserError := context.ShouldBindJSON(&input); bodyParserError != nil {
		utilities.HttpErrorUtils{}.NewHttpError(context, 400, bodyParserError)
		return
	}

	service := services.InternalOtpService{}
	otp, expiryInSeconds, exception := service.GenerateOtp(input.Requester, input.Length, input.OtpLifespanInSeconds)

	if exception != nil {
		utilities.HttpErrorUtils{}.NewHttpError(context, 500, exception)
		return
	}

	response := dto.ApiResultBasicOtp{Otp: otp, ExpiryInSeconds: expiryInSeconds}

	context.JSON(http.StatusOK, response)
}

// GenerateOtpNumber godoc
// @Summary Validates T.O.T.P. number.
// @Description Generates T.O.T.P. number. Read more: https://en.wikipedia.org/wiki/Time-based_One-time_Password_algorithm
// @Tags Bambu
// @Accept  json
// @Produce  json
// @Param payload body dto.ApiInputValidateBasicOtp true "Payload to validate T.O.T.P."
// @Success 200 {object} dto.ApiResultValidateBasicOtp
// @Failure 500 {object} dto.HttpError
// @Router /api/internal/v1/validate [post]
func (instance InternalOtpController) ValidateOtpNumber(context *gin.Context) {
	var input dto.ApiInputValidateBasicOtp

	if bodyParserError := context.ShouldBindJSON(&input); bodyParserError != nil {
		utilities.HttpErrorUtils{}.NewHttpError(context, 400, bodyParserError)
		return
	}

	service := services.InternalOtpService{}
	isValid, exception := service.ValidateOtp(input.Requester, input.Length, input.Interval, input.Otp)

	if exception != nil {
		utilities.HttpErrorUtils{}.NewHttpError(context, 500, exception)
		return
	}

	payload := dto.ApiResultValidateBasicOtp{
		IsValid: isValid,
		Input:   input,
	}

	context.JSON(http.StatusOK, payload)
}
