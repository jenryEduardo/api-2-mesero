package controllers

import (
	"second/robot-status/application"
	"net/http"
	"strconv"
    
	"github.com/gin-gonic/gin"
)

type GetRsIDCtrl struct {
	uc *application.GetbyIdRS
}

func NewGetRsIDCtrl(uc *application.GetbyIdRS) *GetRsIDCtrl {
	return &GetRsIDCtrl{uc: uc}
}

func (ctrl *GetRsIDCtrl) Run(c *gin.Context) {
	id := c.Param("idEstado")
	idStatus, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rs, err := ctrl.uc.Run(idStatus)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": rs,
		})
	}
}