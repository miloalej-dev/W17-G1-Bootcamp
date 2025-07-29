package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmployeeRequest_Bind(t *testing.T) {
	// Valores comunes para todas las pruebas
	id := 1
	cardNumberId := "EMP768"
	firstName := "Doe"
	lastName := "John"
	warehouseId := 1

	// Definición de los casos de prueba
	tests := []struct {
		title         string
		request       *EmployeeRequest
		expectedError string
	}{
		{
			title: "Success - All fields valid",
			request: &EmployeeRequest{
				Id:           &id,
				CardNumberId: &cardNumberId,
				FirstName:    &firstName,
				LastName:     &lastName,
				WarehouseId:  &warehouseId,
			},
			expectedError: "",
		},
		{
			title: "Error - Missing CardNumberId",
			request: &EmployeeRequest{
				Id:           &id,
				CardNumberId: nil,
				FirstName:    &firstName,
				LastName:     &lastName,
				WarehouseId:  &warehouseId,
			},
			expectedError: "CardNumberId must not be null",
		},
		{
			title: "Error - Missing FirstName",
			request: &EmployeeRequest{
				Id:           &id,
				CardNumberId: &cardNumberId,
				FirstName:    nil,
				LastName:     &lastName,
				WarehouseId:  &warehouseId,
			},
			expectedError: "FirstName must not be null",
		},
		{
			title: "Error - Missing LastName",
			request: &EmployeeRequest{
				Id:           &id,
				CardNumberId: &cardNumberId,
				FirstName:    &firstName,
				LastName:     nil,
				WarehouseId:  &warehouseId,
			},
			expectedError: "LastName must not be null",
		},
		{
			title: "Error - Missing WarehouseId",
			request: &EmployeeRequest{
				Id:           &id,
				CardNumberId: &cardNumberId,
				FirstName:    &firstName,
				LastName:     &lastName,
				WarehouseId:  nil,
			},
			expectedError: "WarehouseId must not be null",
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
