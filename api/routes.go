package api

import (
	products "go-http-server-router/api/handlers"
	"go-http-server-router/router"
)

func RegisterRoutes(r *router.Router) {
	// Root route group
	r.RouteGroup("", func(rootGroup *router.RouteGroup) {
		rootGroup.Get("", func(ctx router.Context) router.HttpResponse {
			return router.HttpResponseOK("<h1>Welcome to the API</h1>")
		})
		rootGroup.Get("healthcheck", func(ctx router.Context) router.HttpResponse {
			return router.HttpResponse{
				StatusCode: 200,
				Headers:    map[string]string{"Content-Type": "text/plain"},
				Body:       "OK",
			}
		})
	})

	// Products route group
	r.RouteGroup("/products", func(productsGroup *router.RouteGroup) {
		productsGroup.Get("/list", products.ListProducts)
		productsGroup.Post("/:id", products.CreateProduct)

		// Nested /cars route group under /products
		productsGroup.RouteGroup("/cars", func(carsGroup *router.RouteGroup) {
			carsGroup.Get("", func(ctx router.Context) router.HttpResponse {
				return router.HttpResponseOK("Products Cars")
			})
		})
	})

	// Shops route group
	r.RouteGroup("/shops", func(shopsGroup *router.RouteGroup) {
		shopsGroup.Get("", func(ctx router.Context) router.HttpResponse {
			return router.HttpResponseOK("Shops")
		})
		shopsGroup.Post("asdf", func(ctx router.Context) router.HttpResponse {
			return router.HttpResponseOK("OK")
		})
		shopsGroup.Post("qwer", func(ctx router.Context) router.HttpResponse {
			return router.HttpResponseOK("OK")
		})
		shopsGroup.Post("yxcv", func(ctx router.Context) router.HttpResponse {
			return router.HttpResponseOK("OK")
		})
		shopsGroup.RouteGroup("shops", func(shopsGroup *router.RouteGroup) {
			shopsGroup.RouteGroup("shops", func(shopsGroup *router.RouteGroup) {
				shopsGroup.Get("", func(ctx router.Context) router.HttpResponse {
					return router.HttpResponseOK("Shops Shops Shops")
				})
			})
		})
	})
}
