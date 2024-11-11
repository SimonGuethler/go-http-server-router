package main

import (
	"go-http-server-router/api"
	"go-http-server-router/router"
)

func main() {
	server := router.NewServer()
	server.SetPort(8080)
	server.Register(api.RegisterRoutes)
	server.Start()
}
