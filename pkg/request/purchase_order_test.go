package request

import (
	"net/http"
	"testing"
	"time"

	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestPurchaseOrderRequest_Bind(t *testing.T) {
	orderNumber := "ORD123"
	now := time.Now()
	tracingCode := "TRC456"
	buyerID := 1
	warehouseID := 2
	carrierID := 3
	statusID := 4
	details := []models.OrderDetail{}

	tests := []struct {
		name          string
		request       *PurchaseOrderRequest
		expectedError string
	}{
		{
			name: "Success - All fields present",
			request: &PurchaseOrderRequest{
				OrderNumber:   &orderNumber,
				OrderDate:     &now,
				TracingCode:   &tracingCode,
				BuyersID:      &buyerID,
				WarehousesID:  &warehouseID,
				CarriersID:    &carrierID,
				OrderStatusID: &statusID,
				OrderDetails:  &details,
			},
			expectedError: "",
		},
		{
			name: "Error - OrderNumber is nil",
			request: &PurchaseOrderRequest{
				OrderNumber:   nil,
				OrderDate:     &now,
				TracingCode:   &tracingCode,
				BuyersID:      &buyerID,
				WarehousesID:  &warehouseID,
				CarriersID:    &carrierID,
				OrderStatusID: &statusID,
				OrderDetails:  &details,
			},
			expectedError: "OrderNumber must not be null",
		},
		{
			name: "Error - OrderDate is nil",
			request: &PurchaseOrderRequest{
				OrderNumber:   &orderNumber,
				OrderDate:     nil,
				TracingCode:   &tracingCode,
				BuyersID:      &buyerID,
				WarehousesID:  &warehouseID,
				CarriersID:    &carrierID,
				OrderStatusID: &statusID,
				OrderDetails:  &details,
			},
			expectedError: "OrderDate must not be null",
		},
		{
			name: "Error - TracingCode is nil",
			request: &PurchaseOrderRequest{
				OrderNumber:   &orderNumber,
				OrderDate:     &now,
				TracingCode:   nil,
				BuyersID:      &buyerID,
				WarehousesID:  &warehouseID,
				CarriersID:    &carrierID,
				OrderStatusID: &statusID,
				OrderDetails:  &details,
			},
			expectedError: "TracingCode must not be null",
		},
		{
			name: "Error - BuyersID is nil",
			request: &PurchaseOrderRequest{
				OrderNumber:   &orderNumber,
				OrderDate:     &now,
				TracingCode:   &tracingCode,
				BuyersID:      nil,
				WarehousesID:  &warehouseID,
				CarriersID:    &carrierID,
				OrderStatusID: &statusID,
				OrderDetails:  &details,
			},
			expectedError: "BuyersID must not be null",
		},
		{
			name: "Error - WarehousesID is nil",
			request: &PurchaseOrderRequest{
				OrderNumber:   &orderNumber,
				OrderDate:     &now,
				TracingCode:   &tracingCode,
				BuyersID:      &buyerID,
				WarehousesID:  nil,
				CarriersID:    &carrierID,
				OrderStatusID: &statusID,
				OrderDetails:  &details,
			},
			expectedError: "WarehousesID must not be null",
		},
		{
			name: "Error - CarriersID is nil",
			request: &PurchaseOrderRequest{
				OrderNumber:   &orderNumber,
				OrderDate:     &now,
				TracingCode:   &tracingCode,
				BuyersID:      &buyerID,
				WarehousesID:  &warehouseID,
				CarriersID:    nil,
				OrderStatusID: &statusID,
				OrderDetails:  &details,
			},
			expectedError: "CarriersID must not be null",
		},
		{
			name: "Error - OrderStatusID is nil",
			request: &PurchaseOrderRequest{
				OrderNumber:   &orderNumber,
				OrderDate:     &now,
				TracingCode:   &tracingCode,
				BuyersID:      &buyerID,
				WarehousesID:  &warehouseID,
				CarriersID:    &carrierID,
				OrderStatusID: nil,
				OrderDetails:  &details,
			},
			expectedError: "OrderStatusID must not be null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Bind(&http.Request{})
			if tt.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.EqualError(t, err, tt.expectedError)
			}
		})
	}
}
