package router

import (
	"net"
)

type HTTPMethod string

const (
	GET     HTTPMethod = "GET"
	POST    HTTPMethod = "POST"
	PUT     HTTPMethod = "PUT"
	DELETE  HTTPMethod = "DELETE"
	PATCH   HTTPMethod = "PATCH"
	OPTIONS HTTPMethod = "OPTIONS"
	HEAD    HTTPMethod = "HEAD"
)

type Route struct {
	//id      uuid.UUID
	method  HTTPMethod
	path    string
	pattern string
	handler RouteHandler
}

type RouteHandler func(conn net.Conn)

type Router struct {
	routes map[string]map[string]RouteHandler
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]RouteHandler),
	}
}

type RouteGroup struct {
	prefix string
	router *Router
}

//func (r *Router) RouteGroup(prefix string, callback func(*RouteGroup)) {
//	group := &RouteGroup{
//		prefix: SanitizePath(prefix),
//		router: r,
//	}
//	callback(group)
//}
//
//// TODO: Handle infinite nesting
//func (rg *RouteGroup) RouteGroup(prefix string, callback func(*RouteGroup)) {
//	nestedGroup := &RouteGroup{
//		prefix: SanitizePath(rg.prefix + prefix),
//		router: rg.router,
//	}
//	callback(nestedGroup)
//}
//
//func (rg *RouteGroup) Route(method HTTPMethod, path string, handler RouteHandler) {
//	fullPath := SanitizePath(SanitizePath(rg.prefix) + SanitizePath(path))
//
//	// Add to Router's central route map
//	if _, exists := rg.router.routes[string(method)]; !exists {
//		rg.router.routes[string(method)] = make(map[string]RouteHandler)
//	}
//	rg.router.routes[string(method)][fullPath] = handler
//}
//
//func (rg *RouteGroup) GET(path string, handler RouteHandler)     { rg.Route(GET, path, handler) }
//func (rg *RouteGroup) POST(path string, handler RouteHandler)    { rg.Route(POST, path, handler) }
//func (rg *RouteGroup) PUT(path string, handler RouteHandler)     { rg.Route(PUT, path, handler) }
//func (rg *RouteGroup) DELETE(path string, handler RouteHandler)  { rg.Route(DELETE, path, handler) }
//func (rg *RouteGroup) PATCH(path string, handler RouteHandler)   { rg.Route(PATCH, path, handler) }
//func (rg *RouteGroup) OPTIONS(path string, handler RouteHandler) { rg.Route(OPTIONS, path, handler) }
//func (rg *RouteGroup) HEAD(path string, handler RouteHandler)    { rg.Route(HEAD, path, handler) }
//
//func (r *Router) HandleRequest(conn net.Conn) {
//	defer conn.Close()
//
//	buffer := make([]byte, 4096)
//	_, err := conn.Read(buffer)
//	if err != nil {
//		if err.Error() != "EOF" {
//			log.Println("Error reading:", err)
//		}
//		return
//	}
//
//	requestLine := string(buffer)
//	lines := strings.Split(requestLine, "\r\n")
//	if len(lines) < 1 {
//		return
//	}
//	requestLine = lines[0]
//
//	parts := strings.Fields(requestLine)
//	if len(parts) < 2 {
//		return
//	}
//
//	method := HTTPMethod(parts[0])
//	path := SanitizePath(parts[1])
//
//	// TODO: Check if path length is valid
//
//	var contentLength int
//	for _, line := range lines {
//		if strings.HasPrefix(line, "Content-Length:") {
//			fmt.Sscanf(line, "Content-Length: %d", &contentLength)
//		}
//	}
//
//	// TODO: Check if content length is valid
//
//	if method == POST && contentLength > 0 {
//		bodyBuffer := make([]byte, contentLength)
//		_, err := conn.Read(bodyBuffer)
//		if err != nil {
//			log.Println("Error reading POST body:", err)
//			return
//		}
//		log.Println("POST body:", string(bodyBuffer))
//	}
//
//	// Match the request method and path
//	if methodRoutes, exists := r.routes[string(method)]; exists {
//		if handler, exists := methodRoutes[path]; exists {
//			handler(conn)
//			return
//		}
//
//		// Handle dynamic paths like /{id}
//		if handler := r.matchDynamicRoute(method, path); handler != nil {
//			handler(conn)
//			return
//		}
//	}
//
//	r.notFound(conn)
//}
//
//// TODO: Find better handling for instant dynamic route matching and add path variable extraction
//// matchDynamicRoute tries to match dynamic routes like /{id}
//func (r *Router) matchDynamicRoute(method HTTPMethod, path string) RouteHandler {
//	// Search for routes with dynamic segments (e.g., /products/{id})
//	for routePath, handler := range r.routes[string(method)] {
//		// Match dynamic segments using regex (e.g., /products/{id})
//		re := regexp.MustCompile(`{[a-zA-Z0-9_-]+}`)
//		if re.MatchString(routePath) {
//			// Replace dynamic segments with a general wildcard match
//			pattern := re.ReplaceAllString(routePath, `([^/]+)`)
//			matched, _ := regexp.MatchString(pattern, path)
//			if matched {
//				return handler
//			}
//		}
//	}
//	return nil
//}
//
//// notFound sends a 404 response if no route matches
//func (r *Router) notFound(conn net.Conn) {
//	response := "HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\nPage Not Found"
//	conn.Write([]byte(response))
//	conn.Close()
//}
