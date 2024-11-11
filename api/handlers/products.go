package products

import (
	"go-http-server-router/router"
)

// ListProducts handles the GET request to list products
func ListProducts(ctx router.Context) router.HttpResponse {
	body := `[{"name":"Product 1"}, {"name":"Product 2"}]`
	return router.HttpResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       body,
	}
}

// CreateProduct handles the POST request to create a product
func CreateProduct(ctx router.Context) router.HttpResponse {
	body := `{"name":"New Product"}`
	return router.HttpResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       body,
	}
}
