package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang-poc/oauth2/middleware"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{"*"},
		},
	))
	r.Use(middleware.EnsureValidToken())
	r.GET("/api/external", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"msg": "Your access token was successfully validated!",
		})
	})
	r.Run(":3001")
}
