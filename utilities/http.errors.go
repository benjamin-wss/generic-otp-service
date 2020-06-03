package utilities

import (
	"generic-otp-service/dto"
	"github.com/gin-gonic/gin"
)

type HttpErrorUtils struct{}

func (instance HttpErrorUtils) NewHttpError(context *gin.Context, httpErrorCode int, exception error) {
	er := dto.HttpError{
		Code:    httpErrorCode,
		Message: exception.Error(),
	}
	context.JSON(httpErrorCode, er)
}
