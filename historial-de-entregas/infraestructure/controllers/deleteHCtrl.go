package controllers

import (
	"second/historial-de-entregas/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteHistorialCtrl struct {
	uc *application.DeleteHistorial
}

func NewDeleteHistorialCtrl(uc *application.DeleteHistorial) *DeleteHistorialCtrl {
	return &DeleteHistorialCtrl{uc:uc}
}

func (ctrl *DeleteHistorialCtrl) Run(c *gin.Context) {
	id := c.Param("id_historial")
	idHistorial,err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"El id debe ser númerico"})
		return
	}

	err = ctrl.uc.Run(idHistorial)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":true,
			"data":"historial eliminado",
		})
	}
}