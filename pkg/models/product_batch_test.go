package models

import (
	"testing"

	"github.com/stretchr/testify/assert" // Using assert for direct value comparisons
)

func TestNewProductBatch(t *testing.T) {
	// Define test input values
	expectedId := 1
	expectedBatchNumber := 12345
	expectedCurrentQuantity := 100
	expectedCurrentTemperature := 5.5
	expectedDueDate := "2026-01-15" // Updated to a future date in current context
	expectedInitialQuantity := 200
	expectedManufacturingDate := "2025-07-28" // Current date in context
	expectedManufacturingHour := 14
	expectedMinimumTemperature := -10.0
	expectedSectionId := 7
	expectedProductId := 300

	// Call the constructor function
	productBatch := NewProductBatch(
		expectedId,
		expectedBatchNumber,
		expectedCurrentQuantity,
		expectedCurrentTemperature,
		expectedDueDate,
		expectedInitialQuantity,
		expectedManufacturingDate,
		expectedManufacturingHour,
		expectedMinimumTemperature,
		expectedSectionId,
		expectedProductId,
	)

	// Assert that each field in the returned struct matches the expected input values
	assert.Equal(t, expectedId, productBatch.Id, "Id should match")
	assert.Equal(t, expectedBatchNumber, productBatch.BatchNumber, "BatchNumber should match")
	assert.Equal(t, expectedCurrentQuantity, productBatch.CurrentQuantity, "CurrentQuantity should match")
	assert.Equal(t, expectedCurrentTemperature, productBatch.CurrentTemperature, "CurrentTemperature should match")
	assert.Equal(t, expectedDueDate, productBatch.DueDate, "DueDate should match")
	assert.Equal(t, expectedInitialQuantity, productBatch.InitialQuantity, "InitialQuantity should match")
	assert.Equal(t, expectedManufacturingDate, productBatch.ManufacturingDate, "ManufacturingDate should match")
	assert.Equal(t, expectedManufacturingHour, productBatch.ManufacturingHour, "ManufacturingHour should match")
	assert.Equal(t, expectedMinimumTemperature, productBatch.MinimumTemperature, "MinimumTemperature should match")
	assert.Equal(t, expectedSectionId, productBatch.SectionId, "SectionId should match")
	assert.Equal(t, expectedProductId, productBatch.ProductId, "ProductId should match")
}
