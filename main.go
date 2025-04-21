package main

import (
	"fmt"
	"go-assessment/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default() // Create a Gin router with default middleware

	// Define the GET route
	router.GET("/weth-total-supply", handlers.GetWethTotalSupply)

	port := "8080"
	fmt.Printf("Server starting on port %s using Gin\n", port)
	if err := router.Run(":" + port); err != nil { // Start Gin server
		log.Fatal("Failed to run server: ", err)
	}
}
