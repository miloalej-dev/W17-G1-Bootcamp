package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

// SellerRoutes sets up the routes for product-related operations.
func SellerRoutes(router chi.Router, handler *handler.SellerHandler) {
	router.Route("/api/v1/sellers", func(r chi.Router) {
		r.Get("/", handler.GetSellers)
		r.Get("/{id}", handler.GetSeller)
		r.Post("/", handler.PostSeller)
		r.Put("/{id}", handler.PutSeller)
		r.Patch("/{id}", handler.PatchSeller)
		r.Delete("/{id}", handler.DeleteSeller)
	})
}
