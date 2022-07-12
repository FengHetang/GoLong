package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/index", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"data": "main",
		})
	})
	r.Run(":1013")
}
