package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

func ProductBatchRoutes(rt chi.Router, handler *handler.ProductBatchDefault) {
	rt.Route("/api/v1/productBatches", func(rt chi.Router) {
		// - GET /products
		rt.Post("/", handler.PostProductBatch)
	})
}
