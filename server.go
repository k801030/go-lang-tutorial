package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		date := time.Now().Format("01/02/2006 15:04:05")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
			"date":    date,
		})
	})
	router.Run(":8080")
}
