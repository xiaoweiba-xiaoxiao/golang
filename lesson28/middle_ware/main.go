package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func DeliteTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("hello", "123456")
		c.Next()
		latecy := time.Since(t)
		log.Print(latecy)
	}
}

func main() {
	r := gin.Default()
	r.Use(DeliteTime())
	r.GET("/test", func(c *gin.Context) {
		hello := c.MustGet("hello").(string)
		c.JSON(200, gin.H{
			"message": hello,
		})
	})
	r.Run()
}
