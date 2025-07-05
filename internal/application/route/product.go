package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

func ProductRoutes(rt chi.Router, handler *handler.ProductDefault) {
	rt.Route("/api/v1/products", func(rt chi.Router) {
		// - GET /products
		rt.Get("/", handler.GetAllProducts())
		rt.Post("/", handler.CreateProduct())
		rt.Get("/{ID}", handler.FindByIDProduct())
		rt.Patch("/{ID}", handler.UpdateProduct())
		rt.Delete("/{ID}", handler.DeleteProduct())
	})
}
