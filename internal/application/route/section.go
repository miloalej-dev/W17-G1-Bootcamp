package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

// SectionRoutes sets up the routes for product-related operations.
func SectionRoutes(router chi.Router, handler *handler.SectionHandler) {
	router.Route("/api/v1/sections", func(r chi.Router) {
		r.Get("/", handler.GetSections)
		r.Get("/reportProducts", handler.GetSectionReportProducts)
		r.Get("/{id}", handler.GetSection)
		r.Post("/", handler.PostSection)
		r.Patch("/{id}", handler.PatchSection)
		r.Delete("/{id}", handler.DeleteSection)
	})
}
