package controllers

import (
	"generic-otp-service/dto"
	"generic-otp-service/models"
	"generic-otp-service/repositories"
	"generic-otp-service/services"
	"generic-otp-service/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type InternalOtpController struct {
}

// GenerateOtpNumber godoc
// @Summary Generates T.O.T.P. number.
// @Description Generates T.O.T.P. number. Read more: https://en.wikipedia.org/wiki/Time-based_One-time_Password_algorithm
// @Tags internal-otp
// @Accept  json
// @Produce  json
// @Param payload body dto.ApiInputBasicOtp true "Payload to generate T.O.T.P."
// @Success 200 {object} dto.OtpRepositoryTimeBasedOtpResult
// @Failure 500 {object} dto.HttpError
// @Router /api/internal/v1/acquire [post]
func (instance InternalOtpController) GenerateOtpNumber(context *gin.Context) {
	var input dto.ApiInputBasicOtp

	if bodyParserError := context.ShouldBindJSON(&input); bodyParserError != nil {
		utilities.HttpErrorUtils{}.NewHttpError(context, 400, bodyParserError)
		return
	}

	dbConnection := models.DbPrimary
	otpLogDbRepository := repositories.GetDbOtpLogRepository(dbConnection)
	service := services.InternalOtpService{
		OtpLogDbRepository: otpLogDbRepository,
	}
	result, exception := service.GenerateOtpForApi(strings.TrimSpace(input.Requester), input.Length, input.OtpLifespanInSeconds)

	if exception != nil {
		utilities.HttpErrorUtils{}.NewHttpError(context, exception.HttpStatus, exception.Error)
		return
	}

	context.JSON(http.StatusOK, result)
}

// GenerateOtpNumber godoc
// @Summary Validates T.O.T.P. number.
// @Description Validates T.O.T.P. number. Read more: https://en.wikipedia.org/wiki/Time-based_One-time_Password_algorithm
// @Tags internal-otp
// @Accept  json
// @Produce  json
// @Param payload body dto.ApiInputValidateBasicOtp true "Payload to validate T.O.T.P."
// @Success 200 {object} dto.ApiResultValidateBasicOtp
// @Failure 404 {object} dto.HttpError This denotes that the validation request cannot be validated against a request log entry. Only applicable if acquire request logging is enabled.
// @Failure 400 {object} dto.HttpError
// @Failure 409 {object} dto.HttpError
// @Failure 500 {object} dto.HttpError
// @Router /api/internal/v1/validate [post]
func (instance InternalOtpController) ValidateOtpNumber(context *gin.Context) {
	var input dto.ApiInputValidateBasicOtp

	if bodyParserError := context.ShouldBindJSON(&input); bodyParserError != nil {
		utilities.HttpErrorUtils{}.NewHttpError(context, 400, bodyParserError)
		return
	}

	dbConnection := models.DbPrimary
	otpLogDbRepository := repositories.GetDbOtpLogRepository(dbConnection)
	service := services.InternalOtpService{
		OtpLogDbRepository: otpLogDbRepository,
	}
	isValid, exception := service.ValidateOtpForApi(input.Requester, input.Length, input.OtpLifespanInSeconds, input.Otp, input.ReferenceToken)

	if exception != nil {
		utilities.HttpErrorUtils{}.NewHttpError(context, exception.HttpStatus, exception.Error)
		return
	}

	payload := dto.ApiResultValidateBasicOtp{
		IsValid: isValid,
		Input:   input,
	}

	context.JSON(http.StatusOK, payload)
}
