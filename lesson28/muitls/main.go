package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func multiLoad(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}
	files := form.File["file"]
	var data []interface{}
	for index, val := range files {
		file := fmt.Sprintf("/home/golang/log/%s%d", val.Filename, index)
		err := c.SaveUploadedFile(val, file)
		data = append(data, fmt.Sprintf("%s is load successful!", val.Filename))
		if err != nil {
			continue
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func main() {
	r := gin.Default()
	r.POST("/", multiLoad)
	r.Run()
}
