package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSectionModel(t *testing.T) {
	section := Section{
		Id:                 1,
		SectionNumber:      "A1",
		CurrentTemperature: 5.5,
		MinimumTemperature: 2.0,
		CurrentCapacity:    50,
		MinimumCapacity:    10,
		MaximumCapacity:    100,
		WarehouseId:        3,
		ProductTypeId:      7,
	}

	require.Equal(t, 1, section.Id)
	require.Equal(t, "A1", section.SectionNumber)
	require.Equal(t, 5.5, section.CurrentTemperature)
	require.Equal(t, 2.0, section.MinimumTemperature)
	require.Equal(t, 50, section.CurrentCapacity)
	require.Equal(t, 10, section.MinimumCapacity)
	require.Equal(t, 100, section.MaximumCapacity)
	require.Equal(t, 3, section.WarehouseId)
	require.Equal(t, 7, section.ProductTypeId)
}

func TestSectionReportModel(t *testing.T) {
	report := SectionReport{
		SectionId:     1,
		SectionNumber: "B2",
		ProductsCount: 25,
	}

	require.Equal(t, 1, report.SectionId)
	require.Equal(t, "B2", report.SectionNumber)
	require.Equal(t, 25, report.ProductsCount)
}
