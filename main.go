package main

import (
	"fmt"
	"mouse-mousements-thesis-backend/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hi")
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//r.Use()
	//avantsecure.EndpointProtection("dEdfdNYGegYIneeEkwMLEL6iIXGBjqAiZul7kSBWMLLh80NOG7m6HhrjDDgxKIUl")
	r.POST("/proccessdata", handlers.ProccessData)
	r.GET("/viewdata", handlers.ViewData)
	r.Run(":8000")
}
