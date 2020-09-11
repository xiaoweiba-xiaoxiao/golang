package main

import (
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welecome to index!",
	})
}

func indexPost(c *gin.Context) {
	user := c.PostForm("user")
	passwd := c.PostForm("password")
	data := gin.H{
		"user":     user,
		"password": passwd,
	}
	c.JSON(200, data)
}

func main() {
	r := gin.Default()
	r.GET("/", index)
	r.POST("/", indexPost)
	r.Run()
}
