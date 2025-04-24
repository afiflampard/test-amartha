package main

import (
	"boilerplate/db"
	"boilerplate/forms"
	"fmt"
	"log"
	"os"

	"boilerplate/routes"

	"github.com/gin-contrib/gzip"

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

func main() {
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error : failed to load the env file")
	}

	r := gin.Default()
	binding.Validator = new(forms.DefaultValidator)
	db.Init()
	r.Use(CORSMiddleware())
	r.Use(RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	routes.Routes(r)

	log.Fatal(r.Run(":8000"))

}
