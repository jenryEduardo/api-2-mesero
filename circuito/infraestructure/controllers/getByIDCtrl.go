package controllers

import (
	"second/circuito/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetCircuitoByIdCtrl struct {
	uc *application.GetCircuitoById
}

func NewGetByIDCircuitoCtrl(uc *application.GetCircuitoById) *GetCircuitoByIdCtrl {
	return &GetCircuitoByIdCtrl{uc:uc}
}

func (ctrl *GetCircuitoByIdCtrl) Run(c *gin.Context) {
	id := c.Param("idCircuito")
	idCircuito, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	circuito, err := ctrl.uc.Run(idCircuito)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": circuito,
		})
	}
}