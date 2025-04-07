package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// func GetPedidoStatus(g *gin.Context){
// 	var idPedido domain.Pedido

// 	if err := g.ShouldBindJSON(&idPedido.Status); err != nil {
// 		g.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo obtener el id"})
// 		return
// 	}

// 	jsonData, err := json.Marshal(idPedido)
// 	if err != nil {
// 		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al serializar el JSON"})
// 		return
// 	}

// 	req, err := http.NewRequest("POST", "ruta del publisher por implementar", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la solicitud"})
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar la solicitud"})
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		g.JSON(resp.StatusCode, gin.H{"error": "Error en la respuesta del servidor externo"})
// 		return
// 	}

// 	g.JSON(http.StatusOK, gin.H{"message": "statusPedido enviado correctamente"})
// }