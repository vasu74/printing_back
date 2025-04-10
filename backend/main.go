package main

import (
	"fmt"
	"myproject/db" // Ensure the correct import path
	"myproject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB connection
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
	
	// defer db.CloseDB()

	fmt.Println("Database setup complete!")
}

