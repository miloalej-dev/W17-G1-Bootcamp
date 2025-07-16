package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/application/route"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/json"
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

	// - loader

	lfSection := json.NewFile(a.LoaderFilePathSection)
	dbSection, err := lfSection.LoadSections()

	if err != nil {
		return
	}

	// - repositories

	productRepository := database.NewProductDB(db)
	warehouseRepo := memory.NewWarehouseMap()

	sellerRepository := database.NewSellerRepository(db)
	employeeRepository := database.NewEmployeeRepository(db)
  buyerRepository := database.NewBuyerRepository(db)
	sectionRepository := memory.NewSectionMap(dbSection)

	// - services
	buyerService := _default.NewBuyerDefault(buyerRepository)
	productService := _default.NewProductDefault(productRepository)
	warehouseServ := _default.NewWarehouseDefault(warehouseRepo)
	sellerService := _default.NewSellerService(sellerRepository)
	sectionService := _default.NewSectionDefault(sectionRepository)
	employeeService := _default.NewEmployeeService(employeeRepository)

	// - handlers
	productHandler := handler.NewProductDefault(productService)
	buyerHandler := handler.NewBuyerHandler(buyerService)
	warehouseHand := handler.NewWarehouseDefault(warehouseServ)
	sellerHandler := handler.NewSellerHandler(sellerService)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	sectionHandler := handler.NewSectionDefault(sectionService)

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

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
