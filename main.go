package main

import (
	"log"
	"net/http"

	"github.com/Chinzzii/Kubernetes-Load-Testing-and-Autoscaling/routes"
)

func main() {
	// Initialize the router
	router := routes.InitRouter()

	// Start the server
	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
