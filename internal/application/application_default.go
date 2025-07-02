package application

import (
	"net/http"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
	sectionRepository "github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/section"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/warehouse"
	sectionService "github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/section"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/warehouse"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress  string
	LoaderFilePath string
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
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {
	// dependencies

	// - loader

	// - repositories
	warehouseRepo := repository.NewWarehouseMap()
	sectionRepo := sectionRepository.NewSectionMap()

	// - services
	warehouseServ := service.NewWarehouseDefault(warehouseRepo)
	sectionServ := sectionService.NewSectionDefault(sectionRepo)

	// - handlers
	warehouseHand := handler.NewWarehouseDefault(warehouseServ)
	sectionHand := handler.NewSectionDefault(sectionServ)

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	rt.Route("/api/v1/warehouses", func(rt chi.Router) {
		rt.Get("/", warehouseHand.GetAll())
		rt.Get("/{id}", warehouseHand.GetById())
		rt.Post("/", warehouseHand.Create())
		rt.Patch("/{id}", warehouseHand.Update())
		rt.Delete("/{id}", warehouseHand.Delete())
	})

	rt.Route("/api/v1/section", func(rt chi.Router) {
		rt.Get("/", sectionHand.GetAll())
		rt.Get("/{id}", sectionHand.FindByID())
		rt.Post("/", sectionHand.Create())
		rt.Patch("/{id}", sectionHand.Update())
		rt.Delete("/{id}", sectionHand.Delete())

	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
