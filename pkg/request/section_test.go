package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSectionRequest_Bind(t *testing.T) {
	sectionNumber := "A1"
	currentTemp := 5.5
	minTemp := 2.0
	currentCap := 30
	minCap := 10
	maxCap := 100
	warehouseId := 1
	productTypeId := 2

	tests := []struct {
		title         string
		request       *SectionRequest
		expectedError string
	}{
		{
			title: "Success - all fields valid",
			request: &SectionRequest{
				SectionNumber:      &sectionNumber,
				CurrentTemperature: &currentTemp,
				MinimumTemperature: &minTemp,
				CurrentCapacity:    &currentCap,
				MinimumCapacity:    &minCap,
				MaximumCapacity:    &maxCap,
				WarehouseId:        &warehouseId,
				ProductTypeId:      &productTypeId,
			},
			expectedError: "",
		},
		{
			title: "Error - missing SectionNumber",
			request: &SectionRequest{
				SectionNumber:      nil,
				CurrentTemperature: &currentTemp,
				MinimumTemperature: &minTemp,
				CurrentCapacity:    &currentCap,
				MinimumCapacity:    &minCap,
				MaximumCapacity:    &maxCap,
				WarehouseId:        &warehouseId,
				ProductTypeId:      &productTypeId,
			},
			expectedError: "section_number is required",
		},
		{
			title: "Error - missing CurrentTemperature",
			request: &SectionRequest{
				SectionNumber:      &sectionNumber,
				CurrentTemperature: nil,
				MinimumTemperature: &minTemp,
				CurrentCapacity:    &currentCap,
				MinimumCapacity:    &minCap,
				MaximumCapacity:    &maxCap,
				WarehouseId:        &warehouseId,
				ProductTypeId:      &productTypeId,
			},
			expectedError: "current_temperature is required",
		},
		{
			title: "Error - missing MinimumTemperature",
			request: &SectionRequest{
				SectionNumber:      &sectionNumber,
				CurrentTemperature: &currentTemp,
				MinimumTemperature: nil,
				CurrentCapacity:    &currentCap,
				MinimumCapacity:    &minCap,
				MaximumCapacity:    &maxCap,
				WarehouseId:        &warehouseId,
				ProductTypeId:      &productTypeId,
			},
			expectedError: "minimum_temperature is required",
		},
		{
			title: "Error - missing CurrentCapacity",
			request: &SectionRequest{
				SectionNumber:      &sectionNumber,
				CurrentTemperature: &currentTemp,
				MinimumTemperature: &minTemp,
				CurrentCapacity:    nil,
				MinimumCapacity:    &minCap,
				MaximumCapacity:    &maxCap,
				WarehouseId:        &warehouseId,
				ProductTypeId:      &productTypeId,
			},
			expectedError: "current_capacity is required",
		},
		{
			title: "Error - missing MinimumCapacity",
			request: &SectionRequest{
				SectionNumber:      &sectionNumber,
				CurrentTemperature: &currentTemp,
				MinimumTemperature: &minTemp,
				CurrentCapacity:    &currentCap,
				MinimumCapacity:    nil,
				MaximumCapacity:    &maxCap,
				WarehouseId:        &warehouseId,
				ProductTypeId:      &productTypeId,
			},
			expectedError: "minimum_capacity is required",
		},
		{
			title: "Error - missing MaximumCapacity",
			request: &SectionRequest{
				SectionNumber:      &sectionNumber,
				CurrentTemperature: &currentTemp,
				MinimumTemperature: &minTemp,
				CurrentCapacity:    &currentCap,
				MinimumCapacity:    &minCap,
				MaximumCapacity:    nil,
				WarehouseId:        &warehouseId,
				ProductTypeId:      &productTypeId,
			},
			expectedError: "maximum_capacity is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			err := tc.request.Bind(&http.Request{})

			if tc.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.expectedError)
			}
		})
	}
}
