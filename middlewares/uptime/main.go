package uptime

import (
	"github.com/gin-gonic/gin"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func uptime() time.Duration {
	return time.Since(startTime)
}

func CalculateUptime(context *gin.Context) {
	currentUptime := uptime()

	context.Set("serverUptime", currentUptime)
	context.Next()
}
