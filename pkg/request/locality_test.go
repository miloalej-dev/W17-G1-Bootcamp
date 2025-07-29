package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocalityRequest_Bind(t *testing.T) {
	// Valores comunes para todas las pruebas
	id := 1
	locality := "Buenos Aires"
	province := "Buenos Aires"
	country := "Argentina"

	// Definición de los casos de prueba
	tests := []struct {
		title         string
		request       *LocalityRequest
		expectedError string
	}{
		{
			title: "Success - All fields valid",
			request: &LocalityRequest{
				Id:       id,
				Locality: &locality,
				Province: &province,
				Country:  &country,
			},
			expectedError: "",
		},
		{
			title: "Error - Invalid Id (0)",
			request: &LocalityRequest{
				Id:       0,
				Locality: &locality,
				Province: &province,
				Country:  &country,
			},
			expectedError: "locality_id must be greater than 0",
		},
		{
			title: "Error - Invalid Id (-1)",
			request: &LocalityRequest{
				Id:       -1,
				Locality: &locality,
				Province: &province,
				Country:  &country,
			},
			expectedError: "locality_id must be greater than 0",
		},
		{
			title: "Error - Missing Locality",
			request: &LocalityRequest{
				Id:       id,
				Locality: nil,
				Province: &province,
				Country:  &country,
			},
			expectedError: "locality_name is required",
		},
		{
			title: "Error - Missing Province",
			request: &LocalityRequest{
				Id:       id,
				Locality: &locality,
				Province: nil,
				Country:  &country,
			},
			expectedError: "province_name is required",
		},
		{
			title: "Error - Missing Country",
			request: &LocalityRequest{
				Id:       id,
				Locality: &locality,
				Province: &province,
				Country:  nil,
			},
			expectedError: "country_name is required",
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
