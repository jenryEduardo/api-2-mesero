package controllers

import (
	"second/robot-status/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllRSCtrl struct {
	uc *application.GetAllRS
}

func NewGetAllRSCtrl(uc *application.GetAllRS) *GetAllRSCtrl {
	return &GetAllRSCtrl{uc: uc}
}

func (ctrl *GetAllRSCtrl) Run(c *gin.Context) {
	rs, err := ctrl.uc.Run()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": rs,
		})
	}
}