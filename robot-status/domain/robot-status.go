package domain

type RobotStatus struct {
	IdEstado int `json:"idEstado"`
	IdRobot int `json:"idRobot"`
	Status string `json:"status"`
}

