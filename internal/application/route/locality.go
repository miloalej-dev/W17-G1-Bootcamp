package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

// EmployeeRoutes sets up the routes for employee related operations.
func LocalityRoutes(router chi.Router, handler *handler.LocalityHandler) {

	router.Route("/api/v1/localities", func(r chi.Router) {
		r.Get("/", handler.GetLocalities)
		r.Get("/reportSellers", handler.GetLocality)
		r.Post("/", handler.CreateLocality)
		r.Get("/reportCarriers", handler.GetCarrier)

		//r.Post("/", handler.CreateEmployee)
		//r.Patch("/{id}", handler.PatchEmployee)
		//r.Delete("/{id}", handler.DeleteEmployee)
	})
}
