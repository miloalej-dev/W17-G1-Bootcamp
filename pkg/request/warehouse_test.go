package request

import (
	"testing"
	"github.com/stretchr/testify/require"
	"net/http"
)

func Test_Bind(t *testing.T) {
	t.Run("success", func(t *testing.T){

		code := "WAR-001"
		address := "Calle estación"
		telephone := "111-222333"
		minimumCapacity := 10
		minimumTemperature := 1
		locality_id := 1

		warehouse := WarehouseRequest {
			WarehouseCode:      &code,
			Address:            &address,
			Telephone:          &telephone,
			MinimumCapacity:    &minimumCapacity,
			MinimumTemperature: &minimumTemperature,
			LocalityId:         &locality_id,
		}

		err := warehouse.Bind(&http.Request{})
		require.NoError(t, err)
	})

	t.Run("nil fields", func(t *testing.T) {

		code := ""
		address := ""
		telephone := ""
		var minimumCapacity int
		var minimumTemperature int

		warehouse := WarehouseRequest {
			WarehouseCode:      nil,
			Address:            nil,
			Telephone:          nil,
			MinimumCapacity:    nil,
			MinimumTemperature: nil,
			LocalityId:         nil,
		}
		err := warehouse.Bind(&http.Request{})
		require.Error(t, err)

		warehouse = WarehouseRequest {
			WarehouseCode:      &code,
			Address:            nil,
			Telephone:          nil,
			MinimumCapacity:    nil,
			MinimumTemperature: nil,
			LocalityId:         nil,
		}
		err = warehouse.Bind(&http.Request{})
		require.Error(t, err)

		warehouse = WarehouseRequest {
			WarehouseCode:      &code,
			Address:            &address,
			Telephone:          nil,
			MinimumCapacity:    nil,
			MinimumTemperature: nil,
			LocalityId:         nil,
		}
		err = warehouse.Bind(&http.Request{})
		require.Error(t, err)

		warehouse = WarehouseRequest {
			WarehouseCode:      &code,
			Address:            &address,
			Telephone:          &telephone,
			MinimumCapacity:    nil,
			MinimumTemperature: nil,
			LocalityId:         nil,
		}
		err = warehouse.Bind(&http.Request{})
		require.Error(t, err)

		warehouse = WarehouseRequest {
			WarehouseCode:      &code,
			Address:            &address,
			Telephone:          &telephone,
			MinimumCapacity:    &minimumCapacity,
			MinimumTemperature: nil,
			LocalityId:         nil,
		}
		err = warehouse.Bind(&http.Request{})
		require.Error(t, err)

		warehouse = WarehouseRequest {
			WarehouseCode:      &code,
			Address:            &address,
			Telephone:          &telephone,
			MinimumCapacity:    &minimumCapacity,
			MinimumTemperature: &minimumTemperature,
			LocalityId:         nil,
		}
		err = warehouse.Bind(&http.Request{})
		require.Error(t, err)
	})
}
