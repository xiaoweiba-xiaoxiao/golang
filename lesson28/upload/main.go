package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func loadFile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hollo",
	})
}

func upFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("get file failed,error:%v\n", err)
		return
	}
	path := fmt.Sprintf("/home/golang/log/%s", file.Filename)
	c.SaveUploadedFile(file, path)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s is load sucessfull", file.Filename),
	})
	// file := c.Formfile("file")
	// path := "/home/golang/logs/"
}

func main() {
	r := gin.Default()
	r.GET("/", loadFile)
	r.POST("/", upFile)
	r.Run()
}
