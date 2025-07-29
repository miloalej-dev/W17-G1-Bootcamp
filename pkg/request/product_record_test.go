package request

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_ProductRecordBind(t *testing.T) {

	id := 1
	lastUpdate := "2025-05-12"
	purchasePrice := 5.99
	salePrice := 6.99
	productId := 1

	tests := []struct {
		title         string
		request       *ProductRecordRequest
		expectedError string
	}{
		{
			title: "Success - All fields valid",
			request: &ProductRecordRequest{
				Id:            &id,
				LastUpdate:    &lastUpdate,
				PurchasePrice: &purchasePrice,
				SalePrice:     &salePrice,
				ProductId:     &productId,
			},
			expectedError: "",
		},
		{
			title: "Error - Missing Last Update",
			request: &ProductRecordRequest{
				Id:            &id,
				LastUpdate:    nil,
				PurchasePrice: &purchasePrice,
				SalePrice:     &salePrice,
				ProductId:     &productId,
			},
			expectedError: "last update date must be not null",
		},
		{
			title: "Error - Missing Purchase Price",
			request: &ProductRecordRequest{
				Id:            &id,
				LastUpdate:    &lastUpdate,
				PurchasePrice: nil,
				SalePrice:     &salePrice,
				ProductId:     &productId,
			},
			expectedError: "purchase price  must be not null and greater than 0",
		},
		{
			title: "Error - Missing Sale Price",
			request: &ProductRecordRequest{
				Id:            &id,
				LastUpdate:    &lastUpdate,
				PurchasePrice: &purchasePrice,
				SalePrice:     nil,
				ProductId:     &productId,
			},
			expectedError: "sale price must be not null and greater than 0",
		},
		{
			title: "Error - Missing Product Id",
			request: &ProductRecordRequest{
				Id:            &id,
				LastUpdate:    &lastUpdate,
				PurchasePrice: &purchasePrice,
				SalePrice:     &salePrice,
				ProductId:     nil,
			},
			expectedError: "products Id must be not null and greater than 0",
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
