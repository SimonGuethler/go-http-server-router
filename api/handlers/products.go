package products

import (
	"net"
)

func ListProducts(conn net.Conn) {
	response := "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n\r\n[{\"name\":\"Product 1\"}, {\"name\":\"Product 2\"}]"
	conn.Write([]byte(response))
}

func CreateProduct(conn net.Conn) {
	response := "HTTP/1.1 201 Created\r\nContent-Type: application/json\r\n\r\n{\"name\":\"New Product\"}"
	conn.Write([]byte(response))
}
