package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

func PurchaseOrderRoutes(router chi.Router, handler *handler.PurchaseOrderHandler) {
	router.Route("/api/v1/purchaseOrders", func(r chi.Router) {
		r.Get("/", handler.GetPurchaseOrdersReport)
		r.Post("/", handler.PostPurchaseOrders)

	})
}
