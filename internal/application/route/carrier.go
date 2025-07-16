package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
)

func CarrierRoutes(router chi.Router, handler *handler.CarrierDefault) {
	router.Route("/api/v1/carriers", func(rt chi.Router) {
		rt.Post("/", handler.PostCarrier)
	})
}
