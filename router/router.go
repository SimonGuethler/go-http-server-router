package router

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
)

// RouteHandler is the handler type for routes
type RouteHandler func(conn net.Conn)

// Router stores the routes
type Router struct {
	routes map[string]map[string]RouteHandler
}

// NewRouter creates a new router
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]RouteHandler),
	}
}

// RouteGroup defines a group of routes with a common prefix
type RouteGroup struct {
	prefix string
	router *Router
}

// SanitizePath trims leading/trailing slashes and ensures empty paths are treated as "/"
func SanitizePath(path string) string {
	if path == "" {
		return "/"
	}
	return "/" + strings.Trim(path, "/")
}

// RouteGroup registers a new route group with a given prefix and allows nested sub-groups
func (r *Router) RouteGroup(prefix string, callback func(*RouteGroup)) {
	group := &RouteGroup{
		prefix: SanitizePath(prefix),
		router: r,
	}
	callback(group)
}

// RouteGroup creates a nested route group with a stacked prefix
func (rg *RouteGroup) RouteGroup(prefix string, callback func(*RouteGroup)) {
	nestedGroup := &RouteGroup{
		prefix: SanitizePath(rg.prefix + prefix),
		router: rg.router,
	}
	callback(nestedGroup)
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

// Route handles adding a route to the router
func (rg *RouteGroup) Route(method HTTPMethod, path string, handler RouteHandler) {
	fullPath := SanitizePath(rg.prefix + path)

	// Add to Router's central route map
	if _, exists := rg.router.routes[string(method)]; !exists {
		rg.router.routes[string(method)] = make(map[string]RouteHandler)
	}
	rg.router.routes[string(method)][fullPath] = handler
}

// Route shortcuts for HTTP methods
func (rg *RouteGroup) Get(path string, handler RouteHandler)     { rg.Route(Get, path, handler) }
func (rg *RouteGroup) Post(path string, handler RouteHandler)    { rg.Route(Post, path, handler) }
func (rg *RouteGroup) Put(path string, handler RouteHandler)     { rg.Route(Put, path, handler) }
func (rg *RouteGroup) Delete(path string, handler RouteHandler)  { rg.Route(Delete, path, handler) }
func (rg *RouteGroup) Patch(path string, handler RouteHandler)   { rg.Route(Patch, path, handler) }
func (rg *RouteGroup) Options(path string, handler RouteHandler) { rg.Route(Options, path, handler) }
func (rg *RouteGroup) Head(path string, handler RouteHandler)    { rg.Route(Head, path, handler) }

func (r *Router) HandleRequest(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 4096)
	_, err := conn.Read(buffer)
	if err != nil {
		if err.Error() != "EOF" {
			log.Println("Error reading:", err)
		}
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
	path := SanitizePath(parts[1])

	var contentLength int
	for _, line := range lines {
		if strings.HasPrefix(line, "Content-Length:") {
			fmt.Sscanf(line, "Content-Length: %d", &contentLength)
		}
	}

	if method == Post && contentLength > 0 {
		bodyBuffer := make([]byte, contentLength)
		_, err := conn.Read(bodyBuffer)
		if err != nil {
			log.Println("Error reading POST body:", err)
			return
		}
		log.Println("POST body:", string(bodyBuffer))
	}

	// Match the request method and path
	if methodRoutes, exists := r.routes[string(method)]; exists {
		if handler, exists := methodRoutes[path]; exists {
			handler(conn)
			return
		}

		// Handle dynamic paths like /{id}
		if handler := r.matchDynamicRoute(method, path); handler != nil {
			handler(conn)
			return
		}
	}

	r.notFound(conn)
}

// matchDynamicRoute tries to match dynamic routes like /{id}
func (r *Router) matchDynamicRoute(method HTTPMethod, path string) RouteHandler {
	// Search for routes with dynamic segments (e.g., /products/{id})
	for routePath, handler := range r.routes[string(method)] {
		// Match dynamic segments using regex (e.g., /products/{id})
		re := regexp.MustCompile(`{[a-zA-Z0-9_-]+}`)
		if re.MatchString(routePath) {
			// Replace dynamic segments with a general wildcard match
			pattern := re.ReplaceAllString(routePath, `([^/]+)`)
			matched, _ := regexp.MatchString(pattern, path)
			if matched {
				return handler
			}
		}
	}
	return nil
}

// notFound sends a 404 response if no route matches
func (r *Router) notFound(conn net.Conn) {
	response := "HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\nPage Not Found"
	conn.Write([]byte(response))
	conn.Close()
}
