package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

func ProductRecordRoutes(rt chi.Router, handler *handler.ProductRecordHandler) {
	rt.Route("/api/v1/productRecords", func(rt chi.Router) {
		// - GET /products
		rt.Get("/", handler.GetProductRecords)
		rt.Post("/", handler.PostProductRecord)
		rt.Get("/{id}", handler.GetProductRecord)
		rt.Patch("/{id}", handler.PatchProductRecord)
		rt.Delete("/{id}", handler.DeleteProductRecord)
	})
}
