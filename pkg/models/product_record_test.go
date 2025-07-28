package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProductRecord(t *testing.T) {

	id := 5
	lastUpdate := "2023-10-10"
	purchasePrice := 4.75
	salePrice := 5.50
	productId := 45

	pr := NewProductRecord(id, lastUpdate, purchasePrice, salePrice, productId)

	require.NotNil(t, pr)
	require.Equal(t, id, pr.Id)
	require.Equal(t, lastUpdate, pr.LastUpdate)
	require.Equal(t, purchasePrice, pr.PurchasePrice)
	require.Equal(t, salePrice, pr.SalePrice)
	require.Equal(t, productId, pr.ProductId)
}
