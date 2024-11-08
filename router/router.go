package router

import (
	"log"
	"net"
	"strings"
)

// RouteHandler type for handling requests
type RouteHandler func(conn net.Conn)

// Router struct to hold route groups
type Router struct {
	groups map[string]*RouteGroup
}

// NewRouter creates a new router instance
func NewRouter() *Router {
	return &Router{
		groups: make(map[string]*RouteGroup),
	}
}

// RouteGroup holds routes under a common prefix
type RouteGroup struct {
	prefix string
	routes map[string]map[string]RouteHandler
}

// RouteGroup registers a new route group with a callback
func (r *Router) RouteGroup(prefix string, callback func(*RouteGroup)) {
	group := &RouteGroup{
		prefix: prefix,
		routes: make(map[string]map[string]RouteHandler),
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

// Route registers a new route with the given method and path
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

// addRoute adds a new route to the group with the given method and path
func (rg *RouteGroup) addRoute(method, path string, handler RouteHandler) {
	if _, exists := rg.routes[method]; !exists {
		rg.routes[method] = make(map[string]RouteHandler)
	}
	rg.routes[method][path] = handler
}

// HandleRequest handles incoming requests and calls the appropriate handler
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

	for _, group := range r.groups {
		if strings.HasPrefix(path, group.prefix) {
			pathWithoutPrefix := strings.TrimPrefix(path, group.prefix)
			if handler, exists := group.matchRoute(method, pathWithoutPrefix); exists {
				handler(conn)
				return
			}
		}
	}
	r.notFound(conn)
}

// matchRoute matches static and dynamic routes
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

// matchPath matches paths with dynamic segments like /{id}
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
