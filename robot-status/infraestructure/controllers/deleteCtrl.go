package controllers

import (
	"second/robot-status/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteRSCtrl struct {
	uc *application.DeleteRS
}

func NewDeleteRSCtrl(uc *application.DeleteRS) *DeleteRSCtrl {
	return &DeleteRSCtrl{uc: uc}
}

func (ctrl *DeleteRSCtrl) Run(c *gin.Context) {
	id := c.Param("idEstado")
	idStatus, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser númerico"})
		return
	}

	err = ctrl.uc.Run(idStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": "Estado eliminado",
		})
	}
}