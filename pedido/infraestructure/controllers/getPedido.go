package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdPedido(g *gin.Context) {
	// Obtener el idPedido de la URL
	idString := g.Param("idPedido")
	id, err := strconv.Atoi(idString)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Crear el JSON con el idPedido
	payload := map[string]interface{}{
		"idPedido": id,
	}

	// Serializar el JSON
	body, err := json.Marshal(payload)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al serializar el JSON"})
		return
	}

	// Construir la URL
	url := fmt.Sprintf("http://54.81.41.160:3010/enviarPedido")

	// Crear la solicitud HTTP POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la solicitud"})
		return
	}

	// Establecer los encabezados
	req.Header.Set("Content-Type", "application/json")

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar la solicitud"})
		return
	}
	defer resp.Body.Close()

	// Verificar la respuesta del servidor
	if resp.StatusCode != http.StatusOK {
		g.JSON(resp.StatusCode, gin.H{"error": "Error en la respuesta del servidor externo"})
		return
	}

	// Responder al cliente
	g.JSON(http.StatusOK, gin.H{"message": "Pedido enviado correctamente"})
}
