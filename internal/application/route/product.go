package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

func ProductRoutes(rt chi.Router, handler *handler.ProductDefault) {
	rt.Route("/api/v1/products", func(rt chi.Router) {
		// - GET /products
		rt.Get("/", handler.GetProducts)
		rt.Get("/reportRecords", handler.GetProductReport)
		rt.Post("/", handler.PostProduct)
		rt.Get("/{id}", handler.GetProduct)
		rt.Patch("/{id}", handler.PatchProduct)
		rt.Delete("/{id}", handler.DeleteProduct)

	})
}
