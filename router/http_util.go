package router

import (
	"fmt"
	"net"
)

type HttpResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

func (r *HttpResponse) Write(conn net.Conn) {
	response := fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.StatusCode, httpStatusCodes[r.StatusCode].Description)
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	for key, value := range r.Headers {
		response += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	response += "\r\n" + r.Body
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response:", err)
		closeConn(conn)
		return
	}
}

func closeConn(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Println("Error closing connection:", err)
	}
}

func (h *HTTPMethod) canHaveBody() bool {
	return *h == POST || *h == PUT || *h == PATCH || *h == DELETE || *h == OPTIONS || *h == CONNECT || *h == TRACE
}

func notFound(conn net.Conn) {
	response := &HttpResponse{
		StatusCode: 404,
		Headers:    map[string]string{"Content-Type": PlainText},
		Body:       "Not Found",
	}
	response.Write(conn)
}
