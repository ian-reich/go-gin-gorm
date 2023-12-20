package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	route := app

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hai"})
	})
}
