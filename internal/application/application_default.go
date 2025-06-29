package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler/product"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/product"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/product"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/product"
	"net/http"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the products
	LoaderFilePathProducts string
}
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePathProducts is the path to the file that contains the products
	loaderFilePathProducts string
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
		if cfg.LoaderFilePathProducts != "" {
			defaultConfig.LoaderFilePathProducts = cfg.LoaderFilePathProducts
		}
	}

	return &ServerChi{
		serverAddress:          defaultConfig.ServerAddress,
		loaderFilePathProducts: defaultConfig.LoaderFilePathProducts,
	}
}

// ServerChi is a struct that implements the Application interface

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {
	// dependencies
	// - loader
	ld := productLoader.NewProductJSONFile(a.loaderFilePathProducts)
	dbProduct, err := ld.Load()
	if err != nil {
		return
	}
	// - repositories
	rpProduct := productRepository.NewProductMap(dbProduct)
	// - services
	svProduct := productService.NewProductDefault(rpProduct)
	// - handlers
	hdProduct := productHandler.NewProductDefault(svProduct)

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	rt.Route("/api/v1/", func(rt chi.Router) {
		// - GET /products
		rt.Get("/products", hdProduct.GetAll())

	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
