package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

// EmployeeRoutes sets up the routes for employee related operations.
func EmployeeRoutes(router chi.Router, handler *handler.EmployeeHandler) {

	router.Route("/api/v1/employees", func(r chi.Router) {
		r.Get("/", handler.GetEmployees)
		r.Get("/{id}", handler.GetEmployee)
		r.Get("/reportInboundOrders", handler.GetInboundOrdersReport)
		r.Post("/", handler.CreateEmployee)
		r.Patch("/{id}", handler.PatchEmployee)
		r.Delete("/{id}", handler.DeleteEmployee)
	})
}
