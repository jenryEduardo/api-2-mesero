package controllers

import (
	"second/circuito/application"
	"second/circuito/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateCircuitoCtrl struct {
	uc *application.UpdateCircuito
}

func NewUpdateCircuitoCtrl (uc *application.UpdateCircuito) *UpdateCircuitoCtrl {
	return &UpdateCircuitoCtrl{uc:uc}
}

func (ctrl *UpdateCircuitoCtrl) Run(c *gin.Context) {
	var circuito domain.Circuito

	if err := c.ShouldBindJSON(&circuito); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":"los campos son requeridos"})
		return
	}

	id := c.Param("idCircuito")
	idCircuito, err := strconv.Atoi(id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"El id debe ser númerico",
		})
		return
	}

	if err := ctrl.uc.Run(idCircuito, circuito); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": gin.H{
				"type": "Circuito",
				"attributes": gin.H{
					"idCircuito": circuito.IdCircuito,
					"Id de la mesa": circuito.IdMesa,
					"color del circuito": circuito.Color,
				},
			},
		})
	}
	return
}