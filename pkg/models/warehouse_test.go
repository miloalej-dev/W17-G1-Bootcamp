package models

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_NewWarehouse(t *testing.T) {
	t.Run("success", func(t *testing.T){
		id := 1
		code := "WAR-001"
		address := "Calle estación"
		telephone := "111-222333"
		minimumCapacity := 10
		minimumTemperature := 1
		locality_id := 1

		warehouse := NewWarehouse(id, code, address, telephone, minimumCapacity, minimumTemperature, locality_id)

		require.NotNil(t, warehouse)
		require.Equal(t, address, warehouse.Address)
	})
}
