package controllers

import (
	"second/historial-de-entregas/application"
	"second/historial-de-entregas/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveHistorialCtrl struct {
	uc *application.SaveHistorial
}

func NewSaveHistorialCtrl(uc *application.SaveHistorial) *SaveHistorialCtrl {
	return &SaveHistorialCtrl{uc:uc}
}

func (ctrl *SaveHistorialCtrl) Run(c *gin.Context) {
	var historial domain.Historial

	if err := c.ShouldBindJSON(&historial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.uc.Run(historial)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status":true,
			"data": gin.H{
				"type":"Historial de entrega",
				"id Historial": historial.IdHistorial,
				"attributes": gin.H{
					"ID del Pedido":historial.IdPedido,
					"ID del circuito":historial.IdCircuito,
					"ID del robot": historial.IdRobot,
					"Estatus de la entrega": historial.Estatus_entrega,
				},
			},
		})
	}
}