package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/application/route"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/handler"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/database"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/default"
	"log"
	"net/http"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
}
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
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
	}

	return &ServerChi{
		serverAddress: defaultConfig.ServerAddress,
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

	// - repositories

	productRecordRepository := database.NewProductRecordRepository(db)
	productRepository := database.NewProductRepository(db)
	warehouseRepository := database.NewWarehouseDB(db)
	carrierRepository := database.NewCarrierDB(db)
	productBatchRepository := database.NewProductBatchDB(db)
	sellerRepository := database.NewSellerRepository(db)
	employeeRepository := database.NewEmployeeRepository(db)
	buyerRepository := database.NewBuyerRepository(db)
	sectionRepository := database.NewSectionRepository(db)
	inboundOrderRepository := database.NewInboundOrderRepository(db)
	localityRepository := database.NewLocalityRepository(db)
	purchaseOrderRepository := database.NewPurchaseOrderRepository(db)

	// - services

	productRecordService := _default.NewProductRecordDefault(productRecordRepository)
	productService := _default.NewProductDefault(productRepository)
	warehouseService := _default.NewWarehouseDefault(warehouseRepository)
	carrierService := _default.NewCarrierDefault(carrierRepository)
	productBatchService := _default.NewProductBatchDefault(productBatchRepository)
	buyerService := _default.NewBuyerDefault(buyerRepository)
	sellerService := _default.NewSellerService(sellerRepository)
	sectionService := _default.NewSectionService(sectionRepository)
	employeeService := _default.NewEmployeeService(employeeRepository)
	purchaseOrderService := _default.NewPurchaseOrderDefault(purchaseOrderRepository)
	inboundOrderService := _default.NewInboundOrderService(inboundOrderRepository)
	localityService := _default.NewLocalityService(localityRepository)

	// - handlers
	productHandler := handler.NewProductDefault(productService)
	productBatchHandler := handler.NewProductBatchDefault(productBatchService)
	productRecordHandler := handler.NewProductRecordHandler(productRecordService)
	buyerHandler := handler.NewBuyerHandler(buyerService)
	warehouseHandler := handler.NewWarehouseDefault(warehouseService)
	carrierHandler := handler.NewCarrierDefault(carrierService)
	sellerHandler := handler.NewSellerHandler(sellerService)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	sectionHandler := handler.NewSectionDefault(sectionService)
	purchaseOrderHandler := handler.NewPurchaseOrderDefault(purchaseOrderService)
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
	route.WarehouseRoutes(rt, warehouseHandler)
	route.CarrierRoutes(rt, carrierHandler)
	route.SellerRoutes(rt, sellerHandler)
	route.EmployeeRoutes(rt, employeeHandler)
	route.SectionRoutes(rt, sectionHandler)
	route.ProductRoutes(rt, productHandler)
	route.ProductRecordRoutes(rt, productRecordHandler)
	route.ProductBatchRoutes(rt, productBatchHandler)
	route.PurchaseOrderRoutes(rt, purchaseOrderHandler)
	route.InboundOrderRoutes(rt, inboundOrderHandler)
	route.LocalityRoutes(rt, localityHandler)

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
