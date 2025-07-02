package application

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/application/route"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/buyerLoader"
	loaderProduct "github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/product"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/seller"
	loaderWarehouse "github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/warehouse"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/buyerRepository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/memory"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/buyerService"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/section"
	warehouseService "github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/warehouse"
	"net/http"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler/product"
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
	// LoaderFilePath is the path to the file that contains the warehouses
	LoaderFilePathWarehouse string
}
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePathProducts is the path to the file that contains the buyers
	loaderFilePathBuyer     string
	loaderFilePathProducts  string
	loaderFilePathSeller    string
	loaderFilePathWarehouse string
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
		if cfg.LoaderFilePathWarehouse != "" {
			defaultConfig.LoaderFilePathWarehouse = cfg.LoaderFilePathWarehouse
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
		serverAddress:           defaultConfig.ServerAddress,
		loaderFilePathBuyer:     defaultConfig.LoaderFilePathBuyer,
		loaderFilePathProducts:  defaultConfig.LoaderFilePathProducts,
		loaderFilePathSeller:    defaultConfig.LoaderFilePathSeller,
		loaderFilePathWarehouse: defaultConfig.LoaderFilePathWarehouse,
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

	ldWarehouse := loaderWarehouse.NewJSONFile(a.loaderFilePathWarehouse)
	dbWarehouse, err := ldWarehouse.Load()

	if err != nil {
		return
	}
	// - repositories
	rpProduct := productRepository.NewProductMap(dbProduct)
	warehouseRepo := memory.NewWarehouseMap(dbWarehouse)
	rpBuyer := buyerRepository.NewBuyerMap(dbBuyer)
	sellerRepository := memory.NewSellerMap(dbSeller)
	sectionRepository := memory.NewSectionMap()

	// - services
	svBuyer := buyerService.NewBuyerDefault(rpBuyer)
	svProduct := productService.NewProductDefault(rpProduct)
	warehouseServ := warehouseService.NewWarehouseDefault(warehouseRepo)
	sellerService := service.NewSellerService(sellerRepository)
	sectionService := section.NewSectionDefault(sectionRepository)

	// - handlers
	hdBuyer := handler.NewBuyerHandler(svBuyer)
	hdProduct := productHandler.NewProductDefault(svProduct)
	warehouseHand := handler.NewWarehouseDefault(warehouseServ)
	sellerHandler := handler.NewSellerHandler(sellerService)
	sectionHandler := handler.NewSectionDefault(sectionService)

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
	route.WarehouseRoutes(rt, warehouseHand)
	route.SectionRoutes(rt, sectionHandler)

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

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
