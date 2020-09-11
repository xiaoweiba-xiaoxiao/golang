package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welecom !",
		})
	})
	r.POST("/", func(c *gin.Context) {
		user := c.GetHeader("user")
		password := c.GetHeader("password")
		c.JSON(200, gin.H{
			"user":     user,
			"password": password,
		})
	})
	r.Run(":9090")
}
