package main

import (
	"boilerplate/db"
	"boilerplate/docs"
	"boilerplate/forms"
	"fmt"
	"log"
	"os"

	"boilerplate/routes"

	"github.com/gin-contrib/gzip"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
	"github.com/twinj/uuid"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

// @title           Boilerplate API
// @version         1.0
// @description     API Documentation for your Go Gin boilerplate
// @termsOfService  http://swagger.io/terms/

// @contact.name   Afif Musyayyidin
// @contact.email  musyayyidinafif32@gmail.com

// @license.name  FIFA
// @license.url   https://opensource.org/licenses/FIFA

// @host      localhost:8000
// @BasePath  /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error : failed to load the env file")
	}

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	binding.Validator = new(forms.DefaultValidator)
	db.Init()
	r.Use(CORSMiddleware())
	r.Use(RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	routes.Routes(r)

	log.Fatal(r.Run(":8000"))

}
