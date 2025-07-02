package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/application/route"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
	"net/http"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/buyerLoader"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/buyerRepository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/buyerService"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/seller"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/memory"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/warehouse"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	warehouseService "github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/warehouse"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler/product"
	loaderProduct "github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/product"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/product"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/product"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the Buyers
	LoaderFilePathBuyer string
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
	loaderFilePathSeller   string
	// loaderFilePathProducts is the path to the file that contains the buyers
	loaderFilePathBuyer string
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

		if cfg.LoaderFilePathProducts != "" {
			defaultConfig.LoaderFilePathProducts = cfg.LoaderFilePathProducts
		}
		if cfg.LoaderFilePathSeller != "" {
			defaultConfig.LoaderFilePathSeller = cfg.LoaderFilePathSeller

		}
	}

	return &ServerChi{
		serverAddress:          defaultConfig.ServerAddress,
		loaderFilePathBuyer:    defaultConfig.LoaderFilePathBuyer,
		loaderFilePathProducts: defaultConfig.LoaderFilePathProducts,
		loaderFilePathSeller:   defaultConfig.LoaderFilePathSeller,
	}
}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {
	// dependencies

	// - loader
	ldBuyer := buyerLoader.NewBuyerJSONFile(a.loaderFilePathBuyer)
	dbBuyer, err := ldBuyer.Load()

	ldProduct := loaderProduct.NewProductJSONFile(a.loaderFilePathProducts)
	dbProduct, err := ldProduct.Load()

	ldSeller := seller.NewJSONFile(a.loaderFilePathSeller)
	dbSeller, err := ldSeller.Load()

	if err != nil {
		return
	}
	// - repositories
	rpProduct := productRepository.NewProductMap(dbProduct)
	warehouseRepo := repository.NewWarehouseMap()
	rpBuyer := buyerRepository.NewBuyerMap(dbBuyer)
	sellerRepository := memory.NewSellerMap(dbSeller)

	// - services
	svBuyer := buyerService.NewBuyerDefault(rpBuyer)
	svProduct := productService.NewProductDefault(rpProduct)
	warehouseServ := warehouseService.NewWarehouseDefault(warehouseRepo)
	sellerService := service.NewSellerService(sellerRepository)

	// - handlers
	hdBuyer := handler.NewBuyerHandler(svBuyer)
	hdProduct := productHandler.NewProductDefault(svProduct)
	warehouseHand := handler.NewWarehouseDefault(warehouseServ)
	sellerHandler := handler.NewSellerHandler(sellerService)

	//hd := handler.NewFooHandler()
	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	route.DefaultRoutes(rt)
	route.BuyerRoutes(rt, hdBuyer)
	route.SellerRoutes(rt, sellerHandler)

	/*
		rt.Route("/foo", func(rt chi.Router) {
			rt.Get("/", hd.GetAllFoo)
			rt.Post("/", hd.PostFoo)
		})*/

	rt.Route("/api/v1/", func(rt chi.Router) {
		// - GET /products
		rt.Get("/products", hdProduct.GetAll())
		rt.Post("/products", hdProduct.Create())
		rt.Get("/products/{ID}", hdProduct.FindyByID())
		rt.Patch("/products/{ID}", hdProduct.UpdateProduct())
		rt.Delete("/products/{ID}", hdProduct.Delete())
	})

	rt.Route("/api/v1/warehouses", func(rt chi.Router) {
		rt.Get("/", warehouseHand.FindAll())
		rt.Get("/{id}", warehouseHand.FindById())
		rt.Post("/", warehouseHand.Create())
		rt.Patch("/{id}", warehouseHand.Update())
		rt.Delete("/{id}", warehouseHand.Delete())
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
