package application

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/buyerLoader"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/buyerRepository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/buyerService"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	rpBuyer := buyerRepository.NewBuyerMap(dbBuyer)
	// - services
	svBuyer := buyerService.NewBuyerDefault(rpBuyer)
	// - handlers
	hdBuyer := handler.NewBuyerDefault(svBuyer)
	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	rt.Route("/buyers", func(rt chi.Router) {

		// - GET /
		rt.Get("/", hdBuyer.GetAll())
		rt.Get("/{id}", hdBuyer.GetById())

		// - POST /
		rt.Post("/", hdBuyer.Post())

		// - PATCH /
		rt.Patch("/{id}", hdBuyer.Patch())
		// - DELETE/
		rt.Delete("/{id}", hdBuyer.Delete())

	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
