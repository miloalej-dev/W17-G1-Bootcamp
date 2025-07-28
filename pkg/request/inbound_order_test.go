package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInboundOrder_Bind(t *testing.T) {
	// Common values for all tests
	id := 1
	orderNumber := "ORD-2024-001"
	employeeId := 5
	productBatchId := 10
	warehouseId := 3

	// Test case definitions
	tests := []struct {
		title         string
		request       *InboundOrder
		expectedError string
	}{
		{
			title: "Success - All fields valid",
			request: &InboundOrder{
				Id:             &id,
				OrderNumber:    &orderNumber,
				EmployeeId:     &employeeId,
				ProductBatchId: &productBatchId,
				WarehouseId:    &warehouseId,
			},
			expectedError: "",
		},
		{
			title: "Error - Missing OrderNumber",
			request: &InboundOrder{
				Id:             &id,
				OrderNumber:    nil,
				EmployeeId:     &employeeId,
				ProductBatchId: &productBatchId,
				WarehouseId:    &warehouseId,
			},
			expectedError: "order_number must not be null",
		},
		{
			title: "Error - Missing EmployeeId",
			request: &InboundOrder{
				Id:             &id,
				OrderNumber:    &orderNumber,
				EmployeeId:     nil,
				ProductBatchId: &productBatchId,
				WarehouseId:    &warehouseId,
			},
			expectedError: "employee_id must not be null",
		},
		{
			title: "Error - Missing ProductBatchId",
			request: &InboundOrder{
				Id:             &id,
				OrderNumber:    &orderNumber,
				EmployeeId:     &employeeId,
				ProductBatchId: nil,
				WarehouseId:    &warehouseId,
			},
			expectedError: "product_batch_id must not be null",
		},
		{
			title: "Error - Missing WarehouseId",
			request: &InboundOrder{
				Id:             &id,
				OrderNumber:    &orderNumber,
				EmployeeId:     &employeeId,
				ProductBatchId: &productBatchId,
				WarehouseId:    nil,
			},
			expectedError: "warehouse_id must not be null",
		},
		{
			title: "Error - All required fields missing",
			request: &InboundOrder{
				Id:             &id,
				OrderNumber:    nil,
				EmployeeId:     nil,
				ProductBatchId: nil,
				WarehouseId:    nil,
			},
			expectedError: "order_number must not be null",
		},
		{
			title: "Success - Id can be nil (optional field)",
			request: &InboundOrder{
				Id:             nil,
				OrderNumber:    &orderNumber,
				EmployeeId:     &employeeId,
				ProductBatchId: &productBatchId,
				WarehouseId:    &warehouseId,
			},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			// Act
			err := tc.request.Bind(&http.Request{})

			// Assert
			if tc.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Equal(t, tc.expectedError, err.Error())
			}
		})
	}
}
