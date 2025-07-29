package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductRequest_Bind(t *testing.T) {
	// Common values for all tests
	productCode := "PROD001"
	description := "Delicious example product"
	width := 10.5
	height := 20.0
	length := 30.0
	netWeight := 500.75
	expirationRate := 0.5
	recommendedFreezingTemperature := -18.0
	freezingRate := 0.8
	productTypeId := 101
	sellerId := 500 // Optional, so we'll test with and without it

	// Test case definitions
	tests := []struct {
		title         string
		request       *ProductRequest
		expectedError string
	}{
		{
			title: "Success - All required fields valid",
			request: &ProductRequest{
				ProductCode:                    &productCode,
				Description:                    &description,
				Width:                          &width,
				Height:                         &height,
				Length:                         &length,
				NetWeight:                      &netWeight,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   &freezingRate,
				ProductTypeId:                  &productTypeId,
				SellerId:                       &sellerId, // Including optional field
			},
			expectedError: "",
		},
		{
			title: "Success - All required fields valid, optional SellerId missing",
			request: &ProductRequest{
				ProductCode:                    &productCode,
				Description:                    &description,
				Width:                          &width,
				Height:                         &height,
				Length:                         &length,
				NetWeight:                      &netWeight,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   &freezingRate,
				ProductTypeId:                  &productTypeId,
				SellerId:                       nil, // Omitting optional field
			},
			expectedError: "",
		},
		{
			title: "Error - Missing ProductCode",
			request: &ProductRequest{
				ProductCode:                    nil,
				Description:                    &description,
				Width:                          &width,
				Height:                         &height,
				Length:                         &length,
				NetWeight:                      &netWeight,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   &freezingRate,
				ProductTypeId:                  &productTypeId,
			},
			expectedError: "productCode must not be null",
		},
		{
			title: "Error - Missing Description",
			request: &ProductRequest{
				ProductCode:                    &productCode,
				Description:                    nil,
				Width:                          &width,
				Height:                         &height,
				Length:                         &length,
				NetWeight:                      &netWeight,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   &freezingRate,
				ProductTypeId:                  &productTypeId,
			},
			expectedError: "description must not be null",
		},
		{
			title: "Error - Missing Width",
			request: &ProductRequest{
				ProductCode:                    &productCode,
				Description:                    &description,
				Width:                          nil,
				Height:                         &height,
				Length:                         &length,
				NetWeight:                      &netWeight,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   &freezingRate,
				ProductTypeId:                  &productTypeId,
			},
			expectedError: "width must not be null",
		},
		{
			title: "Error - Missing Height",
			request: &ProductRequest{
				ProductCode:                    &productCode,
				Description:                    &description,
				Width:                          &width,
				Height:                         nil,
				Length:                         &length,
				NetWeight:                      &netWeight,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   &freezingRate,
				ProductTypeId:                  &productTypeId,
			},
			expectedError: "height must not be null",
		},
		{
			title: "Error - Missing Length",
			request: &ProductRequest{
				ProductCode:                    &productCode,
				Description:                    &description,
				Width:                          &width,
				Height:                         &height,
				Length:                         nil,
				NetWeight:                      &netWeight,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   &freezingRate,
				ProductTypeId:                  &productTypeId,
			},
			expectedError: "length must not be null",
		},
		{
			title: "Error - Missing NetWeight",
			request: &ProductRequest{
				ProductCode:                    &productCode,
				Description:                    &description,
				Width:                          &width,
				Height:                         &height,
				Length:                         &length,
				NetWeight:                      nil,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   &freezingRate,
				ProductTypeId:                  &productTypeId,
			},
			expectedError: "netWeight must not be null",
		},
		{
			title: "Error - Missing ExpirationRate",
			request: &ProductRequest{
				ProductCode:                    &productCode,
				Description:                    &description,
				Width:                          &width,
				Height:                         &height,
				Length:                         &length,
				NetWeight:                      &netWeight,
				ExpirationRate:                 nil,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   &freezingRate,
				ProductTypeId:                  &productTypeId,
			},
			expectedError: "expirationRate must not be null",
		},
		{
			title: "Error - Missing RecommendedFreezingTemperature",
			request: &ProductRequest{
				ProductCode:                    &productCode,
				Description:                    &description,
				Width:                          &width,
				Height:                         &height,
				Length:                         &length,
				NetWeight:                      &netWeight,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: nil,
				FreezingRate:                   &freezingRate,
				ProductTypeId:                  &productTypeId,
			},
			expectedError: "recommendedFreezingTemperature must not be null",
		},
		{
			title: "Error - Missing FreezingRate",
			request: &ProductRequest{
				ProductCode:                    &productCode,
				Description:                    &description,
				Width:                          &width,
				Height:                         &height,
				Length:                         &length,
				NetWeight:                      &netWeight,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   nil,
				ProductTypeId:                  &productTypeId,
			},
			expectedError: "freezingRate must not be null",
		},
		{
			title: "Error - Missing ProductTypeId",
			request: &ProductRequest{
				ProductCode:                    &productCode,
				Description:                    &description,
				Width:                          &width,
				Height:                         &height,
				Length:                         &length,
				NetWeight:                      &netWeight,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   &freezingRate,
				ProductTypeId:                  nil,
			},
			expectedError: "productTypeId must not be null",
		},
		{
			title: "Error - All required fields missing (first error should be ProductCode)",
			request: &ProductRequest{
				ProductCode:                    nil,
				Description:                    nil,
				Width:                          nil,
				Height:                         nil,
				Length:                         nil,
				NetWeight:                      nil,
				ExpirationRate:                 nil,
				RecommendedFreezingTemperature: nil,
				FreezingRate:                   nil,
				ProductTypeId:                  nil,
				SellerId:                       nil,
			},
			expectedError: "productCode must not be null",
		},
	}

	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			// Act
			// For the purpose of the Bind method as provided (which only checks for nil pointers),
			// an empty http.Request is sufficient. If Bind were to parse from request body/query,
			// this would need to be mocked appropriately.
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
