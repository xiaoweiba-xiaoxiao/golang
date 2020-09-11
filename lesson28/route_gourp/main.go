package main

import (
	"github.com/gin-gonic/gin"
)

func loginGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "this login get",
	})
}

func welecome(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "this is welecome",
	})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginGet)
		v1.GET("/welecome", welecome)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/ha", loginGet)
		v2.GET("/wel", welecome)
	}
	router.Run()
}
