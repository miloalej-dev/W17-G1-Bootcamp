package request

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)


func Test_WarehouseBind(t *testing.T) {


	code := "WAR-001"
	address := "Calle estación"
	telephone := "111-222333"
	minimumCapacity := 10
	minimumTemperature := 1
	locality_id := 1

	tests := []struct {
		title         string
		request       *WarehouseRequest
		expectedError string
	}{
		{
			title: "Success - All fields valid",
			request: &WarehouseRequest{
				WarehouseCode:      &code,
				Address:            &address,
				Telephone:          &telephone,
				MinimumCapacity:    &minimumCapacity,
				MinimumTemperature: &minimumTemperature,
				LocalityId:         &locality_id,
			},
			expectedError: "",
		},
		{
			title: "Error - Missing code",
			request: &WarehouseRequest{
				WarehouseCode:      nil,
				Address:            &address,
				Telephone:          &telephone,
				MinimumCapacity:    &minimumCapacity,
				MinimumTemperature: &minimumTemperature,
				LocalityId:         &locality_id,
			},
			expectedError: "warehouse code must not be null",
		},
		{
			title: "Error - Missing address",
			request: &WarehouseRequest{
				WarehouseCode:      &code,
				Address:            nil,
				Telephone:          &telephone,
				MinimumCapacity:    &minimumCapacity,
				MinimumTemperature: &minimumTemperature,
				LocalityId:         &locality_id,
			},
			expectedError: "address must not be null",
		},
		{
			title: "Error - Missing telephone",
			request: &WarehouseRequest{
				WarehouseCode:      &code,
				Address:            &address,
				Telephone:          nil,
				MinimumCapacity:    &minimumCapacity,
				MinimumTemperature: &minimumTemperature,
				LocalityId:         &locality_id,
			},
			expectedError: "telephone must not be null",
		},
		{
			title: "Error - Missing minimum capacity",
			request: &WarehouseRequest{
				WarehouseCode:      &code,
				Address:            &address,
				Telephone:          &telephone,
				MinimumCapacity:    nil,
				MinimumTemperature: &minimumTemperature,
				LocalityId:         &locality_id,
			},
			expectedError: "minimum Capacity must not be null",
		},
		{
			title: "Error - Missing minimum temperature",
			request: &WarehouseRequest{
				WarehouseCode:      &code,
				Address:            &address,
				Telephone:          &telephone,
				MinimumCapacity:    &minimumCapacity,
				MinimumTemperature: nil,
				LocalityId:         &locality_id,
			},
			expectedError: "minimum Temperature must not be null",
		},
		{
			title: "Error - Missing locality id",
			request: &WarehouseRequest{
				WarehouseCode:      &code,
				Address:            &address,
				Telephone:          &telephone,
				MinimumCapacity:    &minimumCapacity,
				MinimumTemperature: &minimumTemperature,
				LocalityId:         nil,
			},
			expectedError: "locality id must not be null",
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
