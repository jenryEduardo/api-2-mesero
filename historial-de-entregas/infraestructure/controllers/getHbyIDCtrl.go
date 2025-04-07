package controllers

import (
	"second/historial-de-entregas/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetHistorialByIDCtrl struct {
	uc *application.GetHistorialByID
}

func NewGetHistorialByIDCtrl (uc *application.GetHistorialByID) *GetHistorialByIDCtrl {
	return &GetHistorialByIDCtrl{uc:uc}
}

func (ctrl *GetHistorialByIDCtrl) Run(c *gin.Context) {
	id := c.Param("id_historial")
	idHistorial, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	historial, err := ctrl.uc.Run(idHistorial)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":true,
			"data": historial,
		})
	}
}