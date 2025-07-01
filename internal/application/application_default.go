package application

import (
	"net/http"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/warehouse"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/warehouse"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {
	// dependencies

	// - repositories
	warehouseRepo := repository.NewWarehouseMap()	

	// - services
	warehouseServ := service.NewWarehouseDefault(warehouseRepo)

	// - handlers
	warehouseHand := handler.NewWarehouseDefault(warehouseServ)

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	rt.Route("/api/v1/warehouses", func(rt chi.Router) {
		rt.Get("/", warehouseHand.GetAll())
		rt.Get("/{id}", warehouseHand.GetById())
	})


	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
