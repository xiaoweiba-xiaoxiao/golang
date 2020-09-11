package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	router.GET("/loginForm", func(c *gin.Context) {
		a := struct {
			Name     string
			Password string
			Mobile   string
		}{}
		a.Name = "wanchao"
		a.Password = "hehe"
		a.Mobile = "13579246810"
		c.JSON(200, a)
	})
	router.Run()
}
