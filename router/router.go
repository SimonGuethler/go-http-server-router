package router

import (
	"log"
	"net"
	"strings"
)

type RouteHandler func(conn net.Conn)

type Router struct {
	groups map[string]*RouteGroup
	routes map[string]map[string]RouteHandler
}

func NewRouter() *Router {
	return &Router{
		groups: make(map[string]*RouteGroup),
		routes: make(map[string]map[string]RouteHandler),
	}
}

type RouteGroup struct {
	prefix string
	routes map[string]map[string]RouteHandler
	router *Router
}

func (r *Router) RouteGroup(prefix string, callback func(*RouteGroup)) {
	group := &RouteGroup{
		prefix: prefix,
		routes: make(map[string]map[string]RouteHandler),
		router: r,
	}
	r.groups[prefix] = group
	callback(group)
}

type HTTPMethod string

const (
	Get     HTTPMethod = "GET"
	Post    HTTPMethod = "POST"
	Put     HTTPMethod = "PUT"
	Delete  HTTPMethod = "DELETE"
	Patch   HTTPMethod = "PATCH"
	Options HTTPMethod = "OPTIONS"
	Head    HTTPMethod = "HEAD"
)

func (rg *RouteGroup) Route(method HTTPMethod, path string, handler RouteHandler) {
	rg.addRoute(string(method), path, handler)
}

func (rg *RouteGroup) Get(path string, handler RouteHandler)     { rg.Route(Get, path, handler) }
func (rg *RouteGroup) Post(path string, handler RouteHandler)    { rg.Route(Post, path, handler) }
func (rg *RouteGroup) Put(path string, handler RouteHandler)     { rg.Route(Put, path, handler) }
func (rg *RouteGroup) Delete(path string, handler RouteHandler)  { rg.Route(Delete, path, handler) }
func (rg *RouteGroup) Patch(path string, handler RouteHandler)   { rg.Route(Patch, path, handler) }
func (rg *RouteGroup) Options(path string, handler RouteHandler) { rg.Route(Options, path, handler) }
func (rg *RouteGroup) Head(path string, handler RouteHandler)    { rg.Route(Head, path, handler) }

func (rg *RouteGroup) addRoute(method, path string, handler RouteHandler) {
	fullPath := rg.prefix + path

	if _, exists := rg.routes[method]; !exists {
		rg.routes[method] = make(map[string]RouteHandler)
	}
	rg.routes[method][path] = handler

	if _, exists := rg.router.routes[method]; !exists {
		rg.router.routes[method] = make(map[string]RouteHandler)
	}
	rg.router.routes[method][fullPath] = handler
}

func (r *Router) HandleRequest(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 4096)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error reading:", err)
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

	method := parts[0]
	path := parts[1]

	if handler, exists := r.routes[method][path]; exists {
		handler(conn)
	} else {
		r.notFound(conn)
	}
}

func (rg *RouteGroup) matchRoute(method, path string) (RouteHandler, bool) {
	if routes, ok := rg.routes[method]; ok {
		// Check for exact match
		if handler, exists := routes[path]; exists {
			return handler, true
		}
		// Check for dynamic segments
		for routePath, handler := range routes {
			if matchPath(routePath, path) {
				return handler, true
			}
		}
	}
	return nil, false
}

func matchPath(routePath, requestPath string) bool {
	routeParts := strings.Split(routePath, "/")
	requestParts := strings.Split(requestPath, "/")
	if len(routeParts) != len(requestParts) {
		return false
	}
	for i, part := range routeParts {
		if part != requestParts[i] && !strings.HasPrefix(part, "{") {
			return false
		}
	}
	return true
}

func (r *Router) notFound(conn net.Conn) {
	response := "HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\nPage Not Found"
	conn.Write([]byte(response))
	conn.Close()
}
