package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSectionRequest_Bind_Success(t *testing.T) {
	sectionNumber := "A1"
	currentTemp := 5.5
	minTemp := 2.0
	currentCap := 30
	minCap := 10
	maxCap := 100
	warehouseId := 1
	productTypeId := 2

	req := &SectionRequest{
		SectionNumber:      &sectionNumber,
		CurrentTemperature: &currentTemp,
		MinimumTemperature: &minTemp,
		CurrentCapacity:    &currentCap,
		MinimumCapacity:    &minCap,
		MaximumCapacity:    &maxCap,
		WarehouseId:        &warehouseId,
		ProductTypeId:      &productTypeId,
	}

	err := req.Bind(&http.Request{})
	require.NoError(t, err)
}

func validSectionRequest() *SectionRequest {
	sectionNumber := "A1"
	currentTemp := 5.5
	minTemp := 2.0
	currentCap := 30
	minCap := 10
	maxCap := 100
	warehouseId := 1
	productTypeId := 2

	return &SectionRequest{
		SectionNumber:      &sectionNumber,
		CurrentTemperature: &currentTemp,
		MinimumTemperature: &minTemp,
		CurrentCapacity:    &currentCap,
		MinimumCapacity:    &minCap,
		MaximumCapacity:    &maxCap,
		WarehouseId:        &warehouseId,
		ProductTypeId:      &productTypeId,
	}
}

func TestSectionRequest_Bind_SectionNumber_Error(t *testing.T) {
	req := validSectionRequest()
	req.SectionNumber = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "section_number is required")
}

func TestSectionRequest_Bind_CurrentTemperature_Error(t *testing.T) {
	req := validSectionRequest()
	req.CurrentTemperature = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "current_temperature is required")
}

func TestSectionRequest_Bind_MinimumTemperature_Error(t *testing.T) {
	req := validSectionRequest()
	req.MinimumTemperature = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "minimum_temperature is required")
}

func TestSectionRequest_Bind_CurrentCapacity_Error(t *testing.T) {
	req := validSectionRequest()
	req.CurrentCapacity = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "current_capacity is required")
}

func TestSectionRequest_Bind_MinimumCapacity_Error(t *testing.T) {
	req := validSectionRequest()
	req.MinimumCapacity = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "minimum_capacity is required")
}

func TestSectionRequest_Bind_MaximumCapacity_Error(t *testing.T) {
	req := validSectionRequest()
	req.MaximumCapacity = nil

	err := req.Bind(&http.Request{})
	require.EqualError(t, err, "maximum_capacity is required")
}
