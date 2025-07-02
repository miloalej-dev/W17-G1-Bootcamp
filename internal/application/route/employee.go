package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

// EmployeeRoutes sets up the routes for employee related operations.
func EmployeeRoutes(router chi.Router, handler *handler.EmployeeHandler) {

	router.Route("/api/v1/employees", func(r chi.Router) {
		r.Get("/", handler.GetAllEmployees)
		r.Get("/{id}", handler.GetEmployee)

	})
}
