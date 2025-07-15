package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

func BuyerRoutes(rt chi.Router, handler *handler.BuyerHandler) {
	rt.Route("/api/v1/buyers", func(rt chi.Router) {

		// - GET /
		rt.Get("/", handler.GetBuyers)
		rt.Get("/{id}", handler.GetBuyer)

		// - POST /
		rt.Post("/", handler.PostBuyer)

		// - PATCH /
		rt.Patch("/{id}", handler.PatchBuyer)

		// - DELETE/
		rt.Delete("/{id}", handler.DeleteBuyer)

		rt.Get("/reportPurchaseOrders", handler.GetBuyer)
	})
}
