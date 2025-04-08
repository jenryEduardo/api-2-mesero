package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"second/ultrasonico/domain"
	"github.com/gin-gonic/gin"
)

func SaveIn(c *gin.Context) {
	var stats domain.Sensor

	err := c.ShouldBindJSON(&stats)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(stats)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al convertir datos"})
		return
	}

	// Hacer la petición POST
	resp, err := http.Post("http://54.81.41.160:3010/statusSensor1", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar estado"})
		return
	}
	defer resp.Body.Close()

	// Responder al cliente original
	c.JSON(http.StatusOK, gin.H{"message": "Estado enviado correctamente", "status": stats.Status})
}
