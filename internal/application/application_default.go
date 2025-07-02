package application

import (
	"net/http"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/warehouse"


	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler/product"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/product"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/product"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/product"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the products
	LoaderFilePathProducts string
	//
	LoaderFilePathSeller string
}
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePathProducts is the path to the file that contains the products
	loaderFilePathProducts string
	LoaderFilePathSeller   string
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

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {
	// dependencies

	// - loader
	ldProduct := productLoader.NewProductJSONFile(a.loaderFilePathProducts)
	dbProduct, err := ldProduct.Load()

	if err != nil {
		return
	}
	// - repositories
	rpProduct := productRepository.NewProductMap(dbProduct)
	warehouseRepo := repository.NewWarehouseMap()

	// - services
	svProduct := productService.NewProductDefault(rpProduct)
	warehouseServ := service.NewWarehouseDefault(warehouseRepo)

	// - handlers
	hdProduct := productHandler.NewProductDefault(svProduct)
	warehouseHand := handler.NewWarehouseDefault(warehouseServ)

	hd := handler.NewFooHandler()
	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	rt.Route("/api/v1/", func(rt chi.Router) {
		// - GET /products
		rt.Get("/products", hdProduct.GetAll())
		rt.Post("/products", hdProduct.Create())
		rt.Get("/products/{ID}", hdProduct.FindyByID())
		rt.Patch("/products/{ID}", hdProduct.UpdateProduct())
		rt.Delete("/products/{ID}", hdProduct.Delete())
	})
	rt.Route("/foo", func(rt chi.Router) {
		rt.Get("/", hd.GetAllFoo)
		rt.Post("/", hd.PostFoo)
	})
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
