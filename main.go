package main

import (
	"fmt"
	"netscan/controller"

	"netscan/hosts"
	"netscan/server"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("GO!")
	go controller.Process()
	r := gin.Default()
	r.GET("/stat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hosts": hosts.Hosts,
			"alarm": controller.Alarm,
		})
	})
	r.POST("/stop", func(c *gin.Context) {
		controller.Run = false
	})
	r.POST("/add", server.AddHost)
	r.POST("/remove", server.RemoveHost)
	r.POST("/update", server.UpdateHost)
	r.POST("/start", func(c *gin.Context) {
		if !controller.Run {
			controller.Run = true
			go controller.Process()
		}
	})
	r.Run()

}
