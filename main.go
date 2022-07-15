package main

import (
	"mouse-mousements-thesis-backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/proccessdata", handlers.ProccessData)

	r.Run(":8000")
}
