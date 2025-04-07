package controllers

import (
	"second/circuito/application"
	"second/circuito/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveCircuitoCtrl struct {
	uc *application.SaveCircuito
}

func NewSaveCircuitoCtrl(uc *application.SaveCircuito) *SaveCircuitoCtrl {
	return &SaveCircuitoCtrl{uc:uc}
}

func (ctrl *SaveCircuitoCtrl) Run(c *gin.Context) {
	var circuito domain.Circuito

	if err := c.ShouldBindJSON(&circuito); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	err := ctrl.uc.Run(circuito)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "Circuito",
				"id del circuito": circuito.IdCircuito,
				"attributes": gin.H{
					"id de la mesa": circuito.IdMesa,
					"color": circuito.Color,
				},
			},
		})
	}
}