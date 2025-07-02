package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/application/route"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/buyerLoader"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/buyerRepository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/buyerService"
	"net/http"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/warehouse"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/warehouse"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePathBuyer string
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
		if cfg.LoaderFilePathBuyer != "" {
			defaultConfig.LoaderFilePathBuyer = cfg.LoaderFilePathBuyer
		}
	}

	return &ServerChi{
		serverAddress:       defaultConfig.ServerAddress,
		loaderFilePathBuyer: defaultConfig.LoaderFilePathBuyer,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePathBuyer string
}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {
	// dependencies

	// - loader
	ldBuyer := buyerLoader.NewBuyerJSONFile(a.loaderFilePathBuyer)
	dbBuyer, err := ldBuyer.Load()

	// - repositories
	warehouseRepo := repository.NewWarehouseMap()
	rpBuyer := buyerRepository.NewBuyerMap(dbBuyer)

	// - services
	svBuyer := buyerService.NewBuyerDefault(rpBuyer)
	warehouseServ := service.NewWarehouseDefault(warehouseRepo)

	// - handlers
	hdBuyer := handler.NewBuyerHandler(svBuyer)
	warehouseHand := handler.NewWarehouseDefault(warehouseServ)

	//hd := handler.NewFooHandler()
	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	route.DefaultRoutes(rt)
	route.BuyerRoutes(rt, hdBuyer)

	/*
		rt.Route("/foo", func(rt chi.Router) {
			rt.Get("/", hd.GetAllFoo)
			rt.Post("/", hd.PostFoo)
		})*/
	rt.Route("/api/v1/warehouses", func(rt chi.Router) {
		rt.Get("/", warehouseHand.GetAll())
		rt.Get("/{id}", warehouseHand.GetById())
		rt.Post("/", warehouseHand.Create())
		rt.Patch("/{id}", warehouseHand.Update())
		rt.Delete("/{id}", warehouseHand.Delete())
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
