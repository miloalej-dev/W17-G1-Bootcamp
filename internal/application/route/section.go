package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

// SellerRoutes sets up the routes for product-related operations.
func SectionrRoutes(router chi.Router, handler *handler.SectionDefault) {
	router.Route("/api/v1/sections", func(r chi.Router) {
		r.Get("/", handler.GetAll())
		r.Get("/{id}", handler.FindByID())
		r.Post("/", handler.Create())
		r.Patch("/{id}", handler.Update())
		r.Delete("/{id}", handler.Delete())
	})
}
