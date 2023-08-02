package main

import (
	"GINWeb_project/Controller"
	"GINWeb_project/Repository"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Repository.Init("./data/"); err != nil {
		os.Exit(-1)
	}

	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := Controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})

	err := r.Run()
	if err != nil {
		return
	}
}
