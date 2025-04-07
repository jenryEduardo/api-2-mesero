package main

import (
	"log"
	dependenciesCircuito "second/circuito/infraestructure/dependencies"
	routesCircuito "second/circuito/infraestructure/routes"
	dependenciesHistorial "second/historial-de-entregas/infraestructure/dependencies"
	routesHistorial "second/historial-de-entregas/infraestructure/routes"
	dependenciesStatus "second/robot-status/infraestructure/dependencies"
	routesStatus "second/robot-status/infraestructure/routes"
	dependenciesRobot "second/robot/infraestructure/dependencies"
	routesRobot "second/robot/infraestructure/routes"
	pedidosRoute "second/pedido/infraestructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	//CIRCUITO
	dependenciesCircuito.Init()
	routesCircuito.CircuitRoutes(router)

	//HISTORIAL
	dependenciesHistorial.Init()
	routesHistorial.HistorialRoutes(router)

	//ROBOT
	dependenciesRobot.Init()
	routesRobot.RobotRoutes(router)

	//ROBOT-STATUS
	dependenciesStatus.Init()
	routesStatus.RSroutes(router)


	pedidosRoute.SetUpRoutes(router)

	port := ":8081"
	log.Println("Servidor en el puerto:", port)
	log.Fatal(router.Run(port))
}