package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/application/route"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/database"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/memory"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/default"
	"log"
	"net/http"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the warehouses
	LoaderFilePathEmployee string
	// LoaderFilePath is the path to the file that contains the sections
	LoaderFilePathSection string
}
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePathProducts is the path to the file that contains the buyers

	loaderFilePathEmployee string
	LoaderFilePathSection  string
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
		if cfg.LoaderFilePathEmployee != "" {
			defaultConfig.LoaderFilePathEmployee = cfg.LoaderFilePathEmployee
		}

		if cfg.LoaderFilePathSection != "" {
			defaultConfig.LoaderFilePathSection = cfg.LoaderFilePathSection
		}

	}

	return &ServerChi{
		serverAddress:          defaultConfig.ServerAddress,
		loaderFilePathEmployee: defaultConfig.LoaderFilePathEmployee,
		LoaderFilePathSection:  defaultConfig.LoaderFilePathSection,
	}
}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {
	// Database connection
	db, err := database.NewConnection()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if err != nil {
		return
	}

	// - repositories

	productRepository := database.NewProductDB(db)
	productBatchRepository := database.NewProductBatchDB(db)
	warehouseRepo := memory.NewWarehouseMap()

	sellerRepository := database.NewSellerRepository(db)
	employeeRepository := database.NewEmployeeRepository(db)

	buyerRepository := database.NewBuyerRepository(db)
	sectionRepository := database.NewSectionRepository(db)
	inboundOrderRepository := database.NewInboundOrderRepository(db)
	localityRepository := database.NewLocalityRepository(db)

	// - services
	productService := _default.NewProductDefault(productRepository)
	productBatchService := _default.NewProductBatchDefault(productBatchRepository)
	buyerService := _default.NewBuyerDefault(buyerRepository)
	warehouseServ := _default.NewWarehouseDefault(warehouseRepo)
	sellerService := _default.NewSellerService(sellerRepository)
	sectionService := _default.NewSectionService(sectionRepository)
	employeeService := _default.NewEmployeeService(employeeRepository)
	inboundOrderService := _default.NewInboundOrderService(inboundOrderRepository)
	localityService := _default.NewLocalityService(localityRepository)

	// - handlers
	productHandler := handler.NewProductDefault(productService)
	productBatchHandler := handler.NewProductBatchDefault(productBatchService)
	buyerHandler := handler.NewBuyerHandler(buyerService)
	warehouseHand := handler.NewWarehouseDefault(warehouseServ)
	sellerHandler := handler.NewSellerHandler(sellerService)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	sectionHandler := handler.NewSectionDefault(sectionService)
	inboundOrderHandler := handler.NewInboundOrderHandler(inboundOrderService)
	localityHandler := handler.NewLocalityHandler(localityService)

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints

	route.DefaultRoutes(rt)
	route.BuyerRoutes(rt, buyerHandler)
	route.WarehouseRoutes(rt, warehouseHand)
	route.SellerRoutes(rt, sellerHandler)
	route.EmployeeRoutes(rt, employeeHandler)
	route.SectionRoutes(rt, sectionHandler)
	route.ProductRoutes(rt, productHandler)
	route.ProductBatchRoutes(rt, productBatchHandler)
	route.InboundOrderRoutes(rt, inboundOrderHandler)
	route.LocalityRoutes(rt, localityHandler)

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
