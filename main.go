// Version: 1.0
package main

import (
	"fiet-backend/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.InitDB()
}

func main() {
	config.LoadEnv()
	config.InitDB()

	gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
}
