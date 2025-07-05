package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

func WarehouseRoutes(router chi.Router, handler *handler.WarehouseDefault) {
	router.Route("/api/v1/warehouses", func(rt chi.Router) {
		rt.Get("/", handler.FindAll)
		rt.Get("/{id}", handler.FindById)
		rt.Post("/", handler.Create)
		rt.Patch("/{id}", handler.Update)
		rt.Delete("/{id}", handler.Delete)
	})
}
