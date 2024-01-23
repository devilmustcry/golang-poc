package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PanicErrorHandler(c *gin.Context, err any) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
		"Messager": "Internal Server error",
		"Code":     "500-000-000",
	})
}

func main() {
	gin.SetMode("release")
	r := gin.New()
	r.Use(gin.CustomRecovery(PanicErrorHandler))
	r.GET("/hello", func(ctx *gin.Context) {
		//file, _ := os.Open("./a.json")
		var p *int
		a := *p
		fmt.Println(a)
	})
	r.Run(":8080")

}
