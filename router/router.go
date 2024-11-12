package router

import (
	"log"
	"net"
)

const RequestMaxSize = 4096

type Route struct {
	method  HTTPMethod
	path    string
	handler RouteHandler
}

type Context struct {
	Path        string
	PathParams  map[string]string
	QueryParams map[string]string
	Header      map[string]string
	Body        string
}

type RouteHandler func(context Context) HttpResponse

type Router struct {
	routes *RouteTree
}

func NewRouter() *Router {
	return &Router{routes: NewRouteTree()}
}

type RouteGroup struct {
	route  string
	router *Router
}

func (r *Router) RouteGroup(routePart string, callback func(*RouteGroup)) {
	sanitizedPath, err := SanitizePath(routePart, true)
	if err != nil {
		log.Panicf("Error sanitizing Path: %s\nAllowed characters: %s", err, dynamicPathAllowedChars)
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
		log.Panicf("Error sanitizing Path: %s\nAllowed characters: %s", err, dynamicPathAllowedChars)
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
		log.Panicf("Error sanitizing Path: %s\nAllowed characters: %s\nError: %s", path, dynamicPathAllowedChars, err)
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

	buffer, err := ReadRequest(conn, RequestMaxSize)
	if err != nil {
		log.Println("Error reading request:", err)
		notFound(conn)
		return
	}

	if buffer == nil || len(buffer) == 0 {
		return
	}

	httpParser := NewHttpParser(buffer)
	err = httpParser.Parse()
	if err != nil {
		log.Println("Error parsing request:", err)
		notFound(conn)
		return
	}

	route, err := r.routes.Search(httpParser.Path, httpParser.Method)
	if err != nil {
		log.Println("Error searching route:", err)
		notFound(conn)
		return
	}

	var context Context
	context.Path = httpParser.Path
	context.PathParams = ExtractPathVariables(route.path, httpParser.Path)
	context.QueryParams = httpParser.QueryParams
	context.Header = httpParser.Headers
	context.Body = httpParser.Body

	if handler := route.handler; handler != nil {
		response := handler(context)
		response.Write(conn)
		return
	}
}
