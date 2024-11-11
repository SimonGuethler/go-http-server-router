package router

import (
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	RequestMaxSize = 4096
	PathMaxSize    = 256
)

type Route struct {
	method  HTTPMethod
	path    string
	handler RouteHandler
}

type Context struct {
	path       string
	pathParams map[string]string
	query      string
	header     map[string]string
	body       string
}

type RouteHandler func(context Context) HttpResponse

type Router struct {
	routes *Trie
}

func NewRouter() *Router {
	return &Router{routes: NewTrie()}
}

type RouteGroup struct {
	route  string
	router *Router
}

func (r *Router) RouteGroup(routePart string, callback func(*RouteGroup)) {
	sanitizedPath, err := SanitizePath(routePart, true)
	if err != nil {
		log.Panicf("Error sanitizing path: %s\nAllowed characters: %s", err, dynamicPathAllowedChars)
	}

	group := &RouteGroup{
		route:  sanitizedPath,
		router: r,
	}
	callback(group)
}

func (rg *RouteGroup) RouteGroup(routePart string, callback func(*RouteGroup)) {
	sanitizedPath, err := SanitizePath(routePart, true)
	if err != nil {
		log.Panicf("Error sanitizing path: %s\nAllowed characters: %s", err, dynamicPathAllowedChars)
	}

	nestedGroup := &RouteGroup{
		route:  rg.route + sanitizedPath,
		router: rg.router,
	}
	callback(nestedGroup)
}

func (rg *RouteGroup) Route(method HTTPMethod, path string, handler RouteHandler) {
	sanitizedPath, err := SanitizePath(path, true)
	if err != nil {
		log.Panicf("Error sanitizing path: %s\nAllowed characters: %s\nError: %s", path, dynamicPathAllowedChars, err)
	}

	route := &Route{
		method:  method,
		path:    rg.route + sanitizedPath,
		handler: handler,
	}

	err = rg.router.routes.Insert(route)
	if err != nil {
		log.Panicf("Error inserting route: %s\nError: %s", sanitizedPath, err)
	}
}

func (rg *RouteGroup) Get(path string, handler RouteHandler)     { rg.Route(GET, path, handler) }
func (rg *RouteGroup) Head(path string, handler RouteHandler)    { rg.Route(HEAD, path, handler) }
func (rg *RouteGroup) Post(path string, handler RouteHandler)    { rg.Route(POST, path, handler) }
func (rg *RouteGroup) Put(path string, handler RouteHandler)     { rg.Route(PUT, path, handler) }
func (rg *RouteGroup) Delete(path string, handler RouteHandler)  { rg.Route(DELETE, path, handler) }
func (rg *RouteGroup) Connect(path string, handler RouteHandler) { rg.Route(CONNECT, path, handler) }
func (rg *RouteGroup) Options(path string, handler RouteHandler) { rg.Route(OPTIONS, path, handler) }
func (rg *RouteGroup) Trace(path string, handler RouteHandler)   { rg.Route(TRACE, path, handler) }
func (rg *RouteGroup) Patch(path string, handler RouteHandler)   { rg.Route(PATCH, path, handler) }

func (r *Router) HandleRequest(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing connection:", err)
		}
	}(conn)

	buffer := make([]byte, RequestMaxSize)
	_, err := conn.Read(buffer)
	if err != nil {
		if err.Error() != "EOF" {
			log.Println("Error reading:", err)
		}
		notFound(conn)
		return
	}

	requestLine := string(buffer)
	lines := strings.Split(requestLine, "\r\n")
	if len(lines) < 1 {
		return
	}
	requestLine = lines[0]
	parts := strings.Fields(requestLine)
	if len(parts) < 2 {
		return
	}

	method := HTTPMethod(parts[0])

	path := parts[1]
	if len(path) > PathMaxSize {
		log.Println("Path too long")
		notFound(conn)
		return
	}
	path, err = SanitizePath(path, false)
	if err != nil {
		log.Println("Error sanitizing path:", err)
		notFound(conn)
		return
	}

	route, err := r.routes.Search(path, method)
	if err != nil {
		log.Println("Error searching route:", err)
		notFound(conn)
		return
	}

	var contentLength int
	for _, line := range lines {
		if strings.HasPrefix(line, "Content-Length:") {
			_, err := fmt.Sscanf(line, "Content-Length: %d", &contentLength)
			if err != nil {
				log.Println("Error reading content length:", err)
				return
			}
		}
	}
	// TODO: Parse headers
	// TODO: Parse query string

	var context Context
	context.path = path

	context.pathParams = ExtractPathVariables(route.path, path)

	if contentLength > 0 && method.canHaveBody() {
		bodyBuffer := make([]byte, contentLength)
		_, err := conn.Read(bodyBuffer)
		if err != nil {
			log.Println("Error reading POST body:", err)
			return
		}
		context.body = string(bodyBuffer)
	}

	if handler := route.handler; handler != nil {
		response := handler(context)
		response.Write(conn)
		return
	}
}
