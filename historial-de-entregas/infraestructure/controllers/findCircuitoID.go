package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"second/historial-de-entregas/application"
)

type ColorPayload struct {
	Value string `json:"color"`
}

type FindIdCircuitoCtrl struct {
	uc *application.FindIdCircuito
}

func NewFindIdCircuitoCtrl(uc *application.FindIdCircuito) *FindIdCircuitoCtrl {
	return &FindIdCircuitoCtrl{uc: uc}
}

func (ctrl *FindIdCircuitoCtrl) Run(c *gin.Context) {
	// Obtener el parámetro idPedido de la URL (por ejemplo: /circuito?idPedido=123)
	idPedidoStr := c.Param("idPedido")

	// Convertir el idPedido a entero
	idPedido, err := strconv.Atoi(idPedidoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idPedido debe ser un número válido"})
		return
	}
// Suponiendo que ctrl.uc.Run(idPedido) retorna un string con el color
colorValue, err := ctrl.uc.Run(idPedido)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener color"})
	return
}

// Crear el objeto que cumple con el JSON esperado por el receptor
	payload := ColorPayload{Value: colorValue}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al codificar JSON"})
		return
	}

	// Hacer la solicitud POST a la URL externa
	resp, err := http.Post("http://13.219.94.246:8088/enviar-color", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar datos a servidor remoto"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "El servidor remoto respondió con error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"idPedido": idPedido,
		"color":    colorValue,
		"status":   "Enviado correctamente a /enviar-color",
	})

}
