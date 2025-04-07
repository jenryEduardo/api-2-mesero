package controllers

import (
	"second/circuito/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllCircuitoCtrl struct {
	uc *application.GetAllCircuitos
}

func NewGetAllCircuitosCtrl(uc *application.GetAllCircuitos) *GetAllCircuitoCtrl {
	return &GetAllCircuitoCtrl{uc:uc}
}

func(ctrl *GetAllCircuitoCtrl) Run(c *gin.Context) {
	circuito, err := ctrl.uc.Run()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": circuito,
		})
	}
}