package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

func BuyerRoutes(rt chi.Router, handler *handler.BuyerHandler) {
	rt.Route("/api/v1/buyers", func(rt chi.Router) {

		// - GET /
		rt.Get("/", handler.GetAll())
		rt.Get("/{id}", handler.GetById())

		// - POST /
		rt.Post("/", handler.Post())

		// - PATCH /
		rt.Patch("/{id}", handler.Patch())

		// - DELETE/
		rt.Delete("/{id}", handler.Delete())

	})
}
