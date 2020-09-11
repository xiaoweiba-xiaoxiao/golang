package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/**/*")
	r.GET("/post/index", func(c *gin.Context) {
		c.HTML(200, "post/index.html", gin.H{
			"title": "post",
		})
	})
	r.GET("/user/index", func(c *gin.Context) {
		c.HTML(200, "user/user.html", gin.H{
			"title": "user",
		})
	})
	r.Run()
}
