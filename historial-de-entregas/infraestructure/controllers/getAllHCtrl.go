package controllers

import (
	"second/historial-de-entregas/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllHistorialCtrl struct {
	uc *application.GetAllHistorial
}

func NewGetAllHistorialCtrl(uc *application.GetAllHistorial) *GetAllHistorialCtrl {
	return &GetAllHistorialCtrl{uc:uc}
}

func (ctrl *GetAllHistorialCtrl) Run(c *gin.Context) {
	historial, err := ctrl.uc.Run()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":true,
			"data": historial,
		})
	}
}