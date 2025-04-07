package controllers

import (
	"net/http"
	"second/robot/application"
	"second/robot/domain"

	"github.com/gin-gonic/gin"
)

// type RobotSerialGenerator struct {
// 	Prefix  string
// 	Counter int
// }


// func NewRobotSerialGenerator(prefix string) *RobotSerialGenerator {
// 	return &RobotSerialGenerator{
// 		Prefix:  prefix,
// 		Counter: 1,
// 	}
// }

// func (r *RobotSerialGenerator) GenerateSerial() string {
// 	year := time.Now().Year() % 100 // Últimos 2 dígitos del año
// 	month := int(time.Now().Month()) // Mes en dos dígitos
// 	serialNumber := fmt.Sprintf("%04d", r.Counter) // Número de serie con 4 dígitos
// 	uuidPart := uuid.New().String()[:8] // 8 primeros caracteres del UUID

// 	serialCode := fmt.Sprintf("%s-%02d-%02d-%s-%s", r.Prefix, year, month, serialNumber, uuidPart)

// 	r.Counter++ // Incrementa el contador para el próximo robot
// 	return serialCode
// }

type SaveRobotCtrl struct {
	uc *application.SaveRobot
}

func NewSaveRobotCtrl(uc *application.SaveRobot) *SaveRobotCtrl {
	return &SaveRobotCtrl{uc: uc}
}

func (ctrl *SaveRobotCtrl) Run(c *gin.Context) {
	var robot domain.Robot

	if err := c.ShouldBindJSON(&robot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	 // Asigna el ID generado

	// 🛠️ Imprimir el ID para depuración
	// fmt.Println("ID Generado:", robot.IdRobot)

	err := ctrl.uc.Run(robot)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respuesta con el ID generado
	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type":    "robot",
			"idRobot": robot.IdRobot, // Ahora incluye el id creado
			"attributes": gin.H{
				"alias": robot.Alias,
			},
		},
	})
}
