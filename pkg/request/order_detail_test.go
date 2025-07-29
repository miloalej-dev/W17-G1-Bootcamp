package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderDetailRequest_Bind(t *testing.T) {
	q := 10
	status := "ok"
	temp := 4.0
	productID := 1
	orderID := 2

	// Casos de prueba
	tests := []struct {
		name          string
		request       *OrderDetailRequest
		expectedError string
	}{
		{
			name: "Success - All fields present",
			request: &OrderDetailRequest{
				Quantity:         &q,
				CleanLinesStatus: &status,
				Temperature:      &temp,
				ProductRecordID:  &productID,
				PurchaseOrderID:  &orderID,
			},
			expectedError: "",
		},
		{
			name: "Error - Quantity is nil",
			request: &OrderDetailRequest{
				Quantity:         nil,
				CleanLinesStatus: &status,
				Temperature:      &temp,
				ProductRecordID:  &productID,
				PurchaseOrderID:  &orderID,
			},
			expectedError: "quantity must not be null",
		},
		{
			name: "Error - CleanLinesStatus is nil",
			request: &OrderDetailRequest{
				Quantity:         &q,
				CleanLinesStatus: nil,
				Temperature:      &temp,
				ProductRecordID:  &productID,
				PurchaseOrderID:  &orderID,
			},
			expectedError: "clean line status must not be null",
		},
		{
			name: "Error - Temperature is nil",
			request: &OrderDetailRequest{
				Quantity:         &q,
				CleanLinesStatus: &status,
				Temperature:      nil,
				ProductRecordID:  &productID,
				PurchaseOrderID:  &orderID,
			},
			expectedError: "temperature must not be null",
		},
		{
			name: "Error - ProductRecordID is nil",
			request: &OrderDetailRequest{
				Quantity:         &q,
				CleanLinesStatus: &status,
				Temperature:      &temp,
				ProductRecordID:  nil,
				PurchaseOrderID:  &orderID,
			},
			expectedError: "product record Id must not be null",
		},
		{
			name: "Error - PurchaseOrderID is nil",
			request: &OrderDetailRequest{
				Quantity:         &q,
				CleanLinesStatus: &status,
				Temperature:      &temp,
				ProductRecordID:  &productID,
				PurchaseOrderID:  nil,
			},
			expectedError: "purchase order Id must not be null",
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
