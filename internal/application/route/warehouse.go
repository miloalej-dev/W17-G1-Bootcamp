package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

func WarehouseRoutes(router chi.Router, handler *handler.WarehouseDefault) {
	router.Route("/api/v1/warehouses", func(rt chi.Router) {
		rt.Get("/", handler.GetWarehouses)
		rt.Get("/{id}", handler.GetWarehouse)
		rt.Post("/", handler.PostWarehouse)
		rt.Patch("/{id}", handler.PatchWarehouse)
		rt.Delete("/{id}", handler.DeleteWarehouse)
	})
}
