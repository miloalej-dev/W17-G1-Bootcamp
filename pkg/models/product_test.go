package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	expectedId := 1
	expectedProductCode := "SKU-XYZ-001"
	expectedDescription := "Premium Organic Apples"
	expectedWidth := 15.2
	expectedHeight := 10.1
	expectedLength := 25.4
	expectedNetWeight := 1.25
	expectedExpirationRate := 0.1
	expectedRecommendedFreezingTemperature := -0.5
	expectedFreezingRate := 0.2
	expectedProductTypeId := 1

	tests := []struct {
		title            string
		sellerIdInput    *int
		expectedSellerId *int
	}{
		{
			title:            "NewProduct with SellerId",
			sellerIdInput:    func(i int) *int { return &i }(100),
			expectedSellerId: func(i int) *int { return &i }(100),
		},
		{
			title:            "NewProduct without SellerId (nil)",
			sellerIdInput:    nil,
			expectedSellerId: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			product := NewProduct(
				expectedId,
				expectedProductCode,
				expectedDescription,
				expectedWidth,
				expectedHeight,
				expectedLength,
				expectedNetWeight,
				expectedExpirationRate,
				expectedRecommendedFreezingTemperature,
				expectedFreezingRate,
				expectedProductTypeId,
				tc.sellerIdInput,
			)

			assert.NotNil(t, product, "NewProduct should return a non-nil Product pointer")
			assert.Equal(t, expectedId, product.Id, "Id should match")
			assert.Equal(t, expectedProductCode, product.ProductCode, "ProductCode should match")
			assert.Equal(t, expectedDescription, product.Description, "Description should match")
			assert.Equal(t, expectedWidth, product.Width, "Width should match")
			assert.Equal(t, expectedHeight, product.Height, "Height should match")
			assert.Equal(t, expectedLength, product.Length, "Length should match")
			assert.Equal(t, expectedNetWeight, product.NetWeight, "NetWeight should match")
			assert.Equal(t, expectedExpirationRate, product.ExpirationRate, "ExpirationRate should match")
			assert.Equal(t, expectedRecommendedFreezingTemperature, product.RecommendedFreezingTemperature, "RecommendedFreezingTemperature should match")
			assert.Equal(t, expectedFreezingRate, product.FreezingRate, "FreezingRate should match")
			assert.Equal(t, expectedProductTypeId, product.ProductTypeId, "ProductTypeId should match")

			if tc.expectedSellerId == nil {
				assert.Nil(t, product.SellerId, "SellerId should be nil")
			} else {
				assert.NotNil(t, product.SellerId, "SellerId should not be nil")
				assert.Equal(t, *tc.expectedSellerId, *product.SellerId, "SellerId value should match")
			}
		})
	}
}
