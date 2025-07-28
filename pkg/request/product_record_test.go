package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductRecordRequest_Bind_Success(t *testing.T) {

	id := 1
	lastUpdate := "2024-07-28"
	purchasePrice := 10.5
	salePrice := 15.0
	productId := 99

	req := &ProductRecordRequest{
		Id:            &id,
		LastUpdate:    &lastUpdate,
		PurchasePrice: &purchasePrice,
		SalePrice:     &salePrice,
		ProductId:     &productId,
	}

	err := req.Bind(&http.Request{})
	require.NoError(t, err)
}

func TestProductRecordRequest_Bind_LastUpdate_Err(t *testing.T) {
	val := 1
	price := 2.5

	req1 := &ProductRecordRequest{
		Id:            &val,
		LastUpdate:    nil,
		PurchasePrice: &price,
		SalePrice:     &price,
		ProductId:     &val,
	}
	err := req1.Bind(&http.Request{})
	require.EqualError(t, err, "last update date must be not null")

	req2 := &ProductRecordRequest{
		Id:            &val,
		LastUpdate:    ptrString(""),
		PurchasePrice: &price,
		SalePrice:     &price,
		ProductId:     &val,
	}
	err = req2.Bind(&http.Request{})
	require.EqualError(t, err, "last update date must be not null")
}

func TestProductRecordRequest_Bind_PurchasePrice_Err(t *testing.T) {
	val := 1
	text := "2024-07-28"

	req1 := &ProductRecordRequest{
		Id:            &val,
		LastUpdate:    &text,
		PurchasePrice: nil,
		SalePrice:     ptrFloat(10),
		ProductId:     &val,
	}
	err := req1.Bind(&http.Request{})
	require.EqualError(t, err, "purchase price  must be not null and greater than 0")

	badPrice := -1.0
	req2 := &ProductRecordRequest{
		Id:            &val,
		LastUpdate:    &text,
		PurchasePrice: &badPrice,
		SalePrice:     ptrFloat(10),
		ProductId:     &val,
	}
	err = req2.Bind(&http.Request{})
	require.EqualError(t, err, "purchase price  must be not null and greater than 0")
}

func TestProductRecordRequest_Bind_SalePrice_Err(t *testing.T) {
	val := 1
	text := "2024-07-28"
	purchase := 4.7

	req1 := &ProductRecordRequest{
		Id:            &val,
		LastUpdate:    &text,
		PurchasePrice: &purchase,
		SalePrice:     nil,
		ProductId:     &val,
	}
	err := req1.Bind(&http.Request{})
	require.EqualError(t, err, "sale price must be not null and greater than 0")

	badSale := -5.0
	req2 := &ProductRecordRequest{
		Id:            &val,
		LastUpdate:    &text,
		PurchasePrice: &purchase,
		SalePrice:     &badSale,
		ProductId:     &val,
	}
	err = req2.Bind(&http.Request{})
	require.EqualError(t, err, "sale price must be not null and greater than 0")
}

func TestProductRecordRequest_Bind_ProductId_Err(t *testing.T) {
	val := 1
	text := "2024-07-28"
	purchase := 4.7
	sale := 5.2

	req1 := &ProductRecordRequest{
		Id:            &val,
		LastUpdate:    &text,
		PurchasePrice: &purchase,
		SalePrice:     &sale,
		ProductId:     nil,
	}
	err := req1.Bind(&http.Request{})
	require.EqualError(t, err, "products Id must be not null and greater than 0")

	badId := -7
	req2 := &ProductRecordRequest{
		Id:            &val,
		LastUpdate:    &text,
		PurchasePrice: &purchase,
		SalePrice:     &sale,
		ProductId:     &badId,
	}
	err = req2.Bind(&http.Request{})
	require.EqualError(t, err, "products Id must be not null and greater than 0")
}

// Helpers
func ptrString(s string) *string  { return &s }
func ptrFloat(f float64) *float64 { return &f }
