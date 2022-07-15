package handlers

import (
	"fmt"
	"log"
	"mouse-mousements-thesis-backend/calculations"
	"net/http"

	"github.com/gin-gonic/gin"
)

type dataPayload struct {
	MouseMovement []struct {
		X int     `json:"x"`
		Y int     `json:"y"`
		Z float64 `json:"t"`
	} `json:"mouseMovement"`
}

func ProccessData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Content-Type", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var userData dataPayload

	if err := c.BindJSON(&userData); err != nil {
		log.Fatal(err)
		c.String(http.StatusInternalServerError, "invalid data submitted")
		return
	}

	movementData := userData.MouseMovement

	for i := 0; i < len(movementData)-1; i++ {
		fmt.Println(calculations.CalculateDistance(movementData[i].X, movementData[i+1].X, movementData[i].Y, movementData[i+1].Y))
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
