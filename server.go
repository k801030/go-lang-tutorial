package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		date := time.Now().Format("01/02/2006 15:04:05")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
			"date":    date,
		})
	})
	r.Run(":8080")
}
