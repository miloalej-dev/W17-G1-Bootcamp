package request

import (
	"net/http"
	"testing"
	"time"

	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestPurchaseOrderRequest_Bind_Success(t *testing.T) {
	orderNumber := "ORD123"
	now := time.Now()
	tracingCode := "TRC456"
	buyerID := 1
	warehouseID := 2
	carrierID := 3
	statusID := 4
	details := []models.OrderDetail{}

	req := &PurchaseOrderRequest{
		OrderNumber:   &orderNumber,
		OrderDate:     &now,
		TracingCode:   &tracingCode,
		BuyersID:      &buyerID,
		WarehousesID:  &warehouseID,
		CarriersID:    &carrierID,
		OrderStatusID: &statusID,
		OrderDetails:  &details,
	}

	err := req.Bind(&http.Request{})
	require.NoError(t, err)
}

func validPurchaseOrderRequest() *PurchaseOrderRequest {
	orderNumber := "ORD123"
	now := time.Now()
	tracingCode := "TRC456"
	buyerID := 1
	warehouseID := 2
	carrierID := 3
	statusID := 4
	details := []models.OrderDetail{}

	return &PurchaseOrderRequest{
		OrderNumber:   &orderNumber,
		OrderDate:     &now,
		TracingCode:   &tracingCode,
		BuyersID:      &buyerID,
		WarehousesID:  &warehouseID,
		CarriersID:    &carrierID,
		OrderStatusID: &statusID,
		OrderDetails:  &details,
	}
}

func TestPurchaseOrderRequest_Bind_OrderNumber_Error(t *testing.T) {
	req := validPurchaseOrderRequest()
	req.OrderNumber = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "OrderNumber must not be null")
}

func TestPurchaseOrderRequest_Bind_OrderDate_Error(t *testing.T) {
	req := validPurchaseOrderRequest()
	req.OrderDate = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "OrderDate must not be null")
}

func TestPurchaseOrderRequest_Bind_TracingCode_Error(t *testing.T) {
	req := validPurchaseOrderRequest()
	req.TracingCode = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "TracingCode must not be null")
}

func TestPurchaseOrderRequest_Bind_BuyersID_Error(t *testing.T) {
	req := validPurchaseOrderRequest()
	req.BuyersID = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "BuyersID must not be null")
}

func TestPurchaseOrderRequest_Bind_WarehousesID_Error(t *testing.T) {
	req := validPurchaseOrderRequest()
	req.WarehousesID = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "WarehousesID must not be null")
}

func TestPurchaseOrderRequest_Bind_CarriersID_Error(t *testing.T) {
	req := validPurchaseOrderRequest()
	req.CarriersID = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "CarriersID must not be null")
}

func TestPurchaseOrderRequest_Bind_OrderStatusID_Error(t *testing.T) {
	req := validPurchaseOrderRequest()
	req.OrderStatusID = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "OrderStatusID must not be null")
}
