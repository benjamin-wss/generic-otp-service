package dto

import "time"

type HealthCheckGreeting struct {
	Greeting string        `json:"greeting" example:"Ah, la vache! Ze service is working !"`
	Date     time.Time     `json:"date" example:"2020-06-04T00:00:16.2963059+08:00"`
	Uptime   time.Duration `json:"uptime" example:"10178631900"`
}
