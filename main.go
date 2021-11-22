package main

import (
	"NETSCANNER/controller"
	"NETSCANNER/hosts"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("GO!")
	go controller.Process()
	r := gin.Default()
	r.GET("/stat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hosts": hosts.Adresses,
		})
	})
	r.POST("/stop", func(c *gin.Context) {
		controller.Run = false
	})
	r.POST("/start", func(c *gin.Context) {
		if !controller.Run {
			controller.Run = true
			go controller.Process()
		}
	})
	r.Run()

}
