package domain


type Robot struct {
	IdRobot int `json:"idRobot" gorm:"primaryKey"`
	Alias   string `json:"alias"`
	//Quedaría solo idrobot, alias
}