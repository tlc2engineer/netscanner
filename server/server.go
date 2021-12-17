package server

import (
	"net/http"
	"netscan/hosts"

	"github.com/gin-gonic/gin"
)

func AddHost(c *gin.Context) {
	var json hosts.Host
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := hosts.AddHost(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "адрес добавлен"})
}

func RemoveHost(c *gin.Context) {
	var json hosts.Host
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := hosts.RemoveHost(json.IP)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "адрес удален"})

}

func UpdateHost(c *gin.Context) {
	var json hosts.Host
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := hosts.UpdateHost(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "адрес изменен"})
}
