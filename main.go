package main

import (
	"assignment/config"
	"assignment/routes"
	"log"
)

func main() {
	// Initialize configurations
	config.InitDB()
	config.InitRedis()

	// Start the server
	router := routes.SetupRouter()
	log.Fatal(router.Run(":8080"))
}
