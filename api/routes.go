package api

import (
	products "go-basic-http-server/api/handlers"
	"go-basic-http-server/router"
	"net"
)

func RegisterRoutes(r *router.Router) {
	r.RouteGroup("/", func(rootGroup *router.RouteGroup) {
		rootGroup.Get("", func(conn net.Conn) {
			response := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n<h1>Welcome to the API</h1>"
			conn.Write([]byte(response))
		})
		rootGroup.Get("/healthcheck", func(conn net.Conn) {
			response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nOK"
			conn.Write([]byte(response))
		})
	})
	r.RouteGroup("/products", func(productsGroup *router.RouteGroup) {
		productsGroup.Get("/list", products.ListProducts)
		productsGroup.Post("/{id}", products.CreateProduct)
		productsGroup.RouteGroup("/cars", func(group *router.RouteGroup) {
			group.Get("", func(conn net.Conn) {
				response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\n Products Cars"
				conn.Write([]byte(response))
			})
		})
	})
}
