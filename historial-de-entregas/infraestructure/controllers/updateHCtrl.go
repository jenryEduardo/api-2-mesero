package controllers

import (
	"second/historial-de-entregas/application"
	"second/historial-de-entregas/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateHistorialCtrl struct {
	uc *application.UpdateHistorial
}

func NewUpdateHistorialCtrl(uc *application.UpdateHistorial) *UpdateHistorialCtrl {
	return &UpdateHistorialCtrl{uc:uc}
}

func (ctrl *UpdateHistorialCtrl) Run(c *gin.Context) {
	var historial domain.Historial

	if err := c.ShouldBindJSON(&historial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Todos los campos son requeridos"})
		//elimine el return
	}
	id := c.Param("id_historial")
	idHistorial, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El id del historial debe ser un número"})
			return 
	}

	if err := ctrl.uc.Run(idHistorial, historial); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
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
	return
}