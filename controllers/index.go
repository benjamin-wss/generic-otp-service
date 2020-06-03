package controllers

import (
	"generic-otp-service/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type IndexController struct {
}

// Get godoc
// @Summary Gets status of current server
// @Description Returns values regarding sever uptime and caller HTTP request metadata
// @Tags server-health-check
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.HealthCheckGreeting
// @Router / [get]
func (instance IndexController) Get(context *gin.Context) {
	serverUptime := context.MustGet("serverUptime").(time.Duration)

	greeting := dto.HealthCheckGreeting{
		Greeting: "Ah, la vache! Ze service is working !",
		Date:     time.Now(),
		Uptime:   serverUptime,
	}

	context.JSON(http.StatusOK, greeting)
}
