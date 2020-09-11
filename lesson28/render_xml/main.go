package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/xml", func(c *gin.Context) {
		type People struct {
			Name string
			Age  int
			Sex  string
		}
		var data People
		data.Name = "hehe"
		data.Age = 18
		data.Sex = "girl"
		c.XML(http.StatusOK, data)
	})
	router.Run()
}
