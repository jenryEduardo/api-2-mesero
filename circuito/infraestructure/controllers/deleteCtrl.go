package controllers

import (
	"second/circuito/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCircuitoCtrl struct {
	uc *application.DeleteCircuito
}

func NewDeleteCircuitoCtrl (uc *application.DeleteCircuito) *DeleteCircuitoCtrl {
	return &DeleteCircuitoCtrl{uc:uc}
}

func (ctrl *DeleteCircuitoCtrl) Run(c *gin.Context) {
	id := c.Param("idCircuito")
	idCircuito, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser númerico"})
		return
	}

	err = ctrl.uc.Run(idCircuito)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": "Circuito eliminado",
		})
	}
}