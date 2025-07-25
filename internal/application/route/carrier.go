package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

func CarrierRoutes(router chi.Router, handler *handler.CarrierDefault) {
	router.Route("/api/v1/carriers", func(rt chi.Router) {
		rt.Get("/", handler.GetCarriers)
		rt.Get("/{id}", handler.GetCarrier)
		rt.Post("/", handler.PostCarrier)
		rt.Put("/{id}", handler.PutCarrier)
		rt.Patch("/{id}", handler.PatchCarrier)
		rt.Delete("/{id}", handler.DeleteCarrier)
	})
}
