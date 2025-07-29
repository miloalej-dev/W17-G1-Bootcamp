package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductBatchRequest_Bind(t *testing.T) {

	id := 123
	batchNumber := 1001
	currentQuantity := 50
	currentTemperature := 2.5
	dueDate := "2025-12-31"
	initialQuantity := 100
	manufacturingDate := "2025-07-28"
	manufacturingHour := 10
	minimumTemperature := -5.0
	sectionId := 1
	productId := 201

	tests := []struct {
		title         string
		request       *ProductBatchRequest
		expectedError string
	}{
		{
			title: "Success - All required fields valid",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        &batchNumber,
				CurrentQuantity:    &currentQuantity,
				CurrentTemperature: &currentTemperature,
				DueDate:            &dueDate,
				InitialQuantity:    &initialQuantity,
				ManufacturingDate:  &manufacturingDate,
				ManufacturingHour:  &manufacturingHour,
				MinimumTemperature: &minimumTemperature,
				SectionId:          &sectionId,
				ProductId:          &productId,
			},
			expectedError: "",
		},
		{
			title: "Success - All required fields valid, Id is nil (as often optional for creation)",
			request: &ProductBatchRequest{
				Id:                 nil,
				BatchNumber:        &batchNumber,
				CurrentQuantity:    &currentQuantity,
				CurrentTemperature: &currentTemperature,
				DueDate:            &dueDate,
				InitialQuantity:    &initialQuantity,
				ManufacturingDate:  &manufacturingDate,
				ManufacturingHour:  &manufacturingHour,
				MinimumTemperature: &minimumTemperature,
				SectionId:          &sectionId,
				ProductId:          &productId,
			},
			expectedError: "",
		},
		{
			title: "Error - Missing BatchNumber",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        nil,
				CurrentQuantity:    &currentQuantity,
				CurrentTemperature: &currentTemperature,
				DueDate:            &dueDate,
				InitialQuantity:    &initialQuantity,
				ManufacturingDate:  &manufacturingDate,
				ManufacturingHour:  &manufacturingHour,
				MinimumTemperature: &minimumTemperature,
				SectionId:          &sectionId,
				ProductId:          &productId,
			},
			expectedError: "BatchNumber must not be null",
		},
		{
			title: "Error - Missing CurrentQuantity",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        &batchNumber,
				CurrentQuantity:    nil,
				CurrentTemperature: &currentTemperature,
				DueDate:            &dueDate,
				InitialQuantity:    &initialQuantity,
				ManufacturingDate:  &manufacturingDate,
				ManufacturingHour:  &manufacturingHour,
				MinimumTemperature: &minimumTemperature,
				SectionId:          &sectionId,
				ProductId:          &productId,
			},
			expectedError: "CurrentQuantity must not be null",
		},
		{
			title: "Error - Missing CurrentTemperature",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        &batchNumber,
				CurrentQuantity:    &currentQuantity,
				CurrentTemperature: nil,
				DueDate:            &dueDate,
				InitialQuantity:    &initialQuantity,
				ManufacturingDate:  &manufacturingDate,
				ManufacturingHour:  &manufacturingHour,
				MinimumTemperature: &minimumTemperature,
				SectionId:          &sectionId,
				ProductId:          &productId,
			},
			expectedError: "CurrentTemperature must not be null",
		},
		{
			title: "Error - Missing DueDate",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        &batchNumber,
				CurrentQuantity:    &currentQuantity,
				CurrentTemperature: &currentTemperature,
				DueDate:            nil,
				InitialQuantity:    &initialQuantity,
				ManufacturingDate:  &manufacturingDate,
				ManufacturingHour:  &manufacturingHour,
				MinimumTemperature: &minimumTemperature,
				SectionId:          &sectionId,
				ProductId:          &productId,
			},
			expectedError: "DueDate must not be null",
		},
		{
			title: "Error - Missing InitialQuantity",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        &batchNumber,
				CurrentQuantity:    &currentQuantity,
				CurrentTemperature: &currentTemperature,
				DueDate:            &dueDate,
				InitialQuantity:    nil,
				ManufacturingDate:  &manufacturingDate,
				ManufacturingHour:  &manufacturingHour,
				MinimumTemperature: &minimumTemperature,
				SectionId:          &sectionId,
				ProductId:          &productId,
			},
			expectedError: "InitialQuantity must not be null",
		},
		{
			title: "Error - Missing ManufacturingDate",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        &batchNumber,
				CurrentQuantity:    &currentQuantity,
				CurrentTemperature: &currentTemperature,
				DueDate:            &dueDate,
				InitialQuantity:    &initialQuantity,
				ManufacturingDate:  nil,
				ManufacturingHour:  &manufacturingHour,
				MinimumTemperature: &minimumTemperature,
				SectionId:          &sectionId,
				ProductId:          &productId,
			},
			expectedError: "ManufacturingDate must not be null",
		},
		{
			title: "Error - Missing ManufacturingHour",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        &batchNumber,
				CurrentQuantity:    &currentQuantity,
				CurrentTemperature: &currentTemperature,
				DueDate:            &dueDate,
				InitialQuantity:    &initialQuantity,
				ManufacturingDate:  &manufacturingDate,
				ManufacturingHour:  nil,
				MinimumTemperature: &minimumTemperature,
				SectionId:          &sectionId,
				ProductId:          &productId,
			},
			expectedError: "ManufacturingHour must not be null",
		},
		{
			title: "Error - Missing MinimumTemperature",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        &batchNumber,
				CurrentQuantity:    &currentQuantity,
				CurrentTemperature: &currentTemperature,
				DueDate:            &dueDate,
				InitialQuantity:    &initialQuantity,
				ManufacturingDate:  &manufacturingDate,
				ManufacturingHour:  &manufacturingHour,
				MinimumTemperature: nil,
				SectionId:          &sectionId,
				ProductId:          &productId,
			},
			expectedError: "MinimumTemperature must not be null",
		},
		{
			title: "Error - Missing SectionId",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        &batchNumber,
				CurrentQuantity:    &currentQuantity,
				CurrentTemperature: &currentTemperature,
				DueDate:            &dueDate,
				InitialQuantity:    &initialQuantity,
				ManufacturingDate:  &manufacturingDate,
				ManufacturingHour:  &manufacturingHour,
				MinimumTemperature: &minimumTemperature,
				SectionId:          nil,
				ProductId:          &productId,
			},
			expectedError: "SectionId must not be null",
		},
		{
			title: "Error - Missing ProductId",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        &batchNumber,
				CurrentQuantity:    &currentQuantity,
				CurrentTemperature: &currentTemperature,
				DueDate:            &dueDate,
				InitialQuantity:    &initialQuantity,
				ManufacturingDate:  &manufacturingDate,
				ManufacturingHour:  &manufacturingHour,
				MinimumTemperature: &minimumTemperature,
				SectionId:          &sectionId,
				ProductId:          nil,
			},
			expectedError: "ProductId must not be null",
		},
		{
			title: "Error - All required fields missing (first error should be BatchNumber)",
			request: &ProductBatchRequest{
				Id:                 &id,
				BatchNumber:        nil,
				CurrentQuantity:    nil,
				CurrentTemperature: nil,
				DueDate:            nil,
				InitialQuantity:    nil,
				ManufacturingDate:  nil,
				ManufacturingHour:  nil,
				MinimumTemperature: nil,
				SectionId:          nil,
				ProductId:          nil,
			},
			expectedError: "BatchNumber must not be null",
		},
	}

	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
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
