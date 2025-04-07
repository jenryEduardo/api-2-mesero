package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"second/historial-de-entregas/application"
)

type FindIdCircuitoCtrl struct {
	uc *application.FindIdCircuito
}

func NewFindIdCircuitoCtrl(uc *application.FindIdCircuito) *FindIdCircuitoCtrl {
	return &FindIdCircuitoCtrl{uc: uc}
}

// Configura el cliente MQTT
func newMQTTClient() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883") // Dirección del broker MQTT (ajústalo si es necesario)
	opts.SetClientID("go-publisher")
	opts.SetUsername("user") // Si el broker requiere autenticación
	opts.SetPassword("password") // Si el broker requiere autenticación

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error al conectar al broker MQTT: %v", token.Error())
	}
	return client
}

// Función para publicar el idCircuito en un tópico MQTT
func publishToMQTT(client mqtt.Client, topic string, message string) {
	token := client.Publish(topic, 0, false, message)
	token.Wait()
	log.Printf("Mensaje enviado al tópico %s: %s", topic, message)
}

// El controlador que maneja la solicitud GET y publica en MQTT
func (ctrl *FindIdCircuitoCtrl) Run(c *gin.Context) {
	// Obtener el parámetro idPedido de la URL (por ejemplo: /circuito?idPedido=123)
	idPedidoStr := c.DefaultQuery("idPedido", "")
	if idPedidoStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idPedido es obligatorio"})
		return
	}

	// Convertir el idPedido a entero
	idPedido, err := strconv.Atoi(idPedidoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idPedido debe ser un número válido"})
		return
	}

	// Llamar al caso de uso para obtener el idCircuito
	idCircuito, err := ctrl.uc.Run(idPedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener idCircuito"})
		return
	}

	// Crear un cliente MQTT
	client := newMQTTClient()
	defer client.Disconnect(250)

	// Publicar el idCircuito en un tópico MQTT
	topic := "esp32/circuito" // El tópico que el ESP32 está escuchando
	message := fmt.Sprintf("idCircuito: %d", idCircuito)

	publishToMQTT(client, topic, message)

	// Devolver el idCircuito como respuesta
	c.JSON(http.StatusOK, gin.H{
		"idPedido":   idPedido,
		"idCircuito": idCircuito,
	})
}
