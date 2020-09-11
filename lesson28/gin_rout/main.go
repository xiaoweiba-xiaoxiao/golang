package main

import (
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welecom to myheart",
	})
}

func logindo(c *gin.Context) {
	user := c.PostForm("user")
	password := c.PostForm("password")
	c.JSON(200, gin.H{
		"user":     user,
		"password": password,
	})
}

func loginPut(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is put",
	})
}

func loginDel(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is delete",
	})
}

func main() {
	r := gin.Default()
	r.GET("/login", login)
	r.POST("/login", logindo)
	r.PUT("/login", loginPut)
	r.DELETE("/login", loginDel)
	r.Run()
}
