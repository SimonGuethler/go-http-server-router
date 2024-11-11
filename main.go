package main

import (
	"fmt"
	"go-basic-http-server/api"
	"go-basic-http-server/router"
	"log"
	"net"
)

func main() {
	// Initialize the router
	r := router.NewRouter()

	// Register product routes
	api.RegisterRoutes(r)

	// Start the server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("Server running at http://localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle the incoming request
		go r.HandleRequest(conn)
	}
}
