package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderDetailRequest_Bind_Success(t *testing.T) {
	q := 10
	status := "ok"
	temp := 4.0
	productID := 1
	orderID := 2

	req := &OrderDetailRequest{
		Quantity:         &q,
		CleanLinesStatus: &status,
		Temperature:      &temp,
		ProductRecordID:  &productID,
		PurchaseOrderID:  &orderID,
	}

	err := req.Bind(&http.Request{})
	require.NoError(t, err)
}

func validOrderDetailRequest() *OrderDetailRequest {
	q := 10
	status := "ok"
	temp := 4.0
	productID := 1
	orderID := 2

	return &OrderDetailRequest{
		Quantity:         &q,
		CleanLinesStatus: &status,
		Temperature:      &temp,
		ProductRecordID:  &productID,
		PurchaseOrderID:  &orderID,
	}
}

func TestOrderDetailRequest_Bind_Quantity_Error(t *testing.T) {
	req := validOrderDetailRequest()
	req.Quantity = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "quantity must not be null")
}

func TestOrderDetailRequest_Bind_CleanLinesStatus_Error(t *testing.T) {
	req := validOrderDetailRequest()
	req.CleanLinesStatus = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "clean line status must not be null")
}

func TestOrderDetailRequest_Bind_Temperature_Error(t *testing.T) {
	req := validOrderDetailRequest()
	req.Temperature = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "temperature must not be null")
}

func TestOrderDetailRequest_Bind_ProductRecordID_Error(t *testing.T) {
	req := validOrderDetailRequest()
	req.ProductRecordID = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "product record Id must not be null")
}

func TestOrderDetailRequest_Bind_PurchaseOrderID_Error(t *testing.T) {
	req := validOrderDetailRequest()
	req.PurchaseOrderID = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "purchase order Id must not be null")
}
