package server

import (
	"NETSCANNER/hosts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHost(c *gin.Context) {
	var json hosts.Address
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hosts.Adresses = append(hosts.Adresses, &json)
	c.JSON(http.StatusOK, gin.H{"status": "адрес добавлен"})
}
