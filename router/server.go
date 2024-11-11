package router

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	router    *Router
	port      int
	isRunning bool
}

func NewServer() *Server {
	return &Server{
		router:    NewRouter(),
		port:      8080,
		isRunning: false,
	}
}

type RegisterRoutes func(r *Router)

func (s *Server) Register(registerRoutes RegisterRoutes) {
	registerRoutes(s.router)
}

func (s *Server) SetPort(port int) {
	s.port = port
}

func (s *Server) Port() int {
	return s.port
}

func (s *Server) Start() {
	if s.isRunning {
		log.Println("Server already running")
		return
	}

	s.isRunning = true

	listener, err := net.Listen("tcp", ":"+fmt.Sprint(s.port))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	defer func(listener net.Listener) {
		s.isRunning = false
		err := listener.Close()
		if err != nil {
			log.Println("Error closing listener:", err)
		}
	}(listener)

	fmt.Printf("Server running at http://localhost:%d\n\n", s.port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle the incoming request
		go s.router.HandleRequest(conn)
	}
}
