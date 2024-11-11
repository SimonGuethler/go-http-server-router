package products

import (
	"go-http-server-router/router"
)

func ListProducts(ctx router.Context) router.HttpResponse {
	body := `[{"name":"Product 1"}, {"name":"Product 2"}]`
	return router.HttpResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       body,
	}
}

func CreateProduct(ctx router.Context) router.HttpResponse {
	pathId := ctx.PathParams["id"]
	queryId := ctx.QueryParams["id"]
	body := ctx.Body

	return router.HttpResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       `{"id": "` + pathId + `", "query": "` + queryId + `", "body": ` + body + `}`,
	}
}
