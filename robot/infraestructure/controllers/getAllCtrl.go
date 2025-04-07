package controllers

import (
	"second/robot/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllRobotsCtrl struct {
	uc *application.GetAllRobot
}

func NewGetAllRobotsCtrl(uc *application.GetAllRobot) *GetAllRobotsCtrl {
	return &GetAllRobotsCtrl{uc: uc}
}

func (ctrl *GetAllRobotsCtrl) Run(c *gin.Context) {
	robots, err := ctrl.uc.Run()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, robots)
	}
}