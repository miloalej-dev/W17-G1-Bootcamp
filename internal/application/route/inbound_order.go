package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

// InboundOrderRoutes sets up the routes for inbound order related operations.
func InboundOrderRoutes(router chi.Router, handler *handler.InboundOrderHandler) {
	router.Route("/api/v1/inbound-orders", func(r chi.Router) {
		r.Get("/", handler.GetInboundOrders)
		r.Get("/{id}", handler.GetInboundOrder)
		r.Post("/", handler.PostInboundOrder)
		r.Put("/{id}", handler.PutInboundOrder)
		r.Patch("/{id}", handler.PatchInboundOrder)
		r.Delete("/{id}", handler.DeleteInboundOrder)
	})
}
