package application

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/application/route"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/memory"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/default"
	"net/http"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string

	// LoaderFilePath is the path to the file that contains the products
	LoaderFilePathProducts string
	// LoaderFilePath is the path to the file that contains the warehouses
	LoaderFilePathEmployee string
}
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePathProducts is the path to the file that contains the buyers

	loaderFilePathProducts  string
	loaderFilePathEmployee  string
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
		if cfg.LoaderFilePathEmployee != "" {
			defaultConfig.LoaderFilePathEmployee = cfg.LoaderFilePathEmployee
		}
	}

	return &ServerChi{
		serverAddress:           defaultConfig.ServerAddress,
		loaderFilePathProducts:  defaultConfig.LoaderFilePathProducts,
		loaderFilePathEmployee:  defaultConfig.LoaderFilePathEmployee,
	}
}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {
	// dependencies

	// - loader

	ldProduct := json.NewProductFile(a.loaderFilePathProducts)
	dbProduct, err := ldProduct.Load()

	ldEmployee := json.NewEmployeeFile(a.loaderFilePathEmployee)
	dbEmployee, err := ldEmployee.Load()

	if err != nil {
		return
	}
	// - repositories
	rpProduct := memory.NewProductMap(dbProduct)
	warehouseRepo := memory.NewWarehouseMap()
	sellerRepository := memory.NewSellerMap()
	employeeRepository := memory.NewEmployeeMap(dbEmployee)
	buyerRepository := memory.NewBuyerMap()
	sectionRepository := memory.NewSectionMap()

	// - services
	buyerService := _default.NewBuyerDefault(buyerRepository)
	svProduct := _default.NewProductDefault(rpProduct)
	warehouseServ := _default.NewWarehouseDefault(warehouseRepo)
	sellerService := _default.NewSellerService(sellerRepository)
	sectionService := _default.NewSectionDefault(sectionRepository)
	employeeService := _default.NewEmployeeService(employeeRepository)

	// - handlers
	buyerHandler := handler.NewBuyerHandler(buyerService)
	hdProduct := handler.NewProductDefault(svProduct)
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
	route.ProductRoutes(rt, hdProduct)

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
