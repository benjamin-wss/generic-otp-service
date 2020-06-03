package main

import (
	"fmt"
	"generic-otp-service/controllers"
	_ "generic-otp-service/docs"
	"generic-otp-service/middlewares/uptime"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Generic OTP Service API
// @version 1.0
// @description Generic OTP API Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Benjamin Wong
// @contact.url http://www.swagger.io/support
// @contact.email do-not-mail-this@pintu.dk

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /

func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.BestSpeed))

	router.GET("/", uptime.CalculateUptime, controllers.IndexController{}.Get)

	v1 := router.Group("/api/internal/v1")
	{
		v1.POST("/acquire", controllers.InternalOtpController{}.GenerateOtpNumber)
		v1.POST("/validate", controllers.InternalOtpController{}.ValidateOtpNumber)
	}

	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	ginHttpPortNumber := fmt.Sprintf(":%d", 3000)

	_ = router.Run(ginHttpPortNumber)
}

//package main
//
//import (
//	"fmt"
//	"github.com/xlzd/gotp"
//	"strings"
//	"time"
//)
//
//func main() {
//	otp := generateOtp()
//	time.Sleep(2 * time.Second)
//
//	res := validateOtp(otp)
//
//	fmt.Printf("Piki: %t", res)
//}
//
//var sec = strings.ToUpper("4S62BZNFXXSZLCRObenpintudk")
//
//func generateOtp() string {
//	otp := gotp.NewTOTP(sec, 6, 1, nil)
//	otpValue, _ := otp.NowWithExpiration()
//
//	return otpValue
//}
//
//func validateOtp(input string) bool {
//	otp := gotp.NewTOTP(sec, 6, 300, nil)
//	res := otp.Verify(input, int(time.Now().Unix()))
//
//	return res
//}
