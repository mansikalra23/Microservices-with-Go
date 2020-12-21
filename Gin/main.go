package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/mansikalra23/Microservices/GinGonic/Gin/middlewares"
)

func main() {
	r := gin.New()

	r.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello KloudOne!",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		body := c.Request.Body // to read the body
		value, err := ioutil.ReadAll(body)

		if err != nil {
			fmt.Println(err.Error())
		}

		c.JSON(200, gin.H{
			"message": string(value), // entering the message
		})
	})

	r.Run(":8000")
}
