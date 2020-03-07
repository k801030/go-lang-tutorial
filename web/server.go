package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ts   string `json:"ts,omitempty"`
	Data Data   `json:"data,omitempty"`
}

type Data struct {
	Name string `json:"name,omitempty"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health-check", HealthCheckHandler)
	r.GET("/user/:name", GetUserNameHandler)
	return r
}

func HealthCheckHandler(c *gin.Context) {
	d := Data{
		Name: "unknown",
	}
	r := Response{
		Ts:   getCurrentDate(),
		Data: d,
	}
	c.JSON(http.StatusOK, r)
}

func GetUserNameHandler(c *gin.Context) {
	name := c.Param("name")
	d := Data{
		Name: name,
	}
	r := Response{
		Ts:   getCurrentDate(),
		Data: d,
	}
	c.JSON(http.StatusOK, r)
}

func main() {
	router := setupRouter()
	router.Run(":80")
}

func getCurrentDate() string {
	return time.Now().Format("01/02/2006 15:04:05")
}
