package main

import (
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is index",
	})
}

func indexPost(c *gin.Context) {
	user := c.Param("user")
	password := c.Param("password")
	// iteams := c.Params()
	c.JSON(200, gin.H{
		"user":     user,
		"password": password,
	})
}

func main() {
	r := gin.Default()
	r.GET("/", index)
	r.POST("/:user/:password", indexPost)
	r.Run()
}
