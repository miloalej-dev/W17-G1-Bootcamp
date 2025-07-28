package request

import (
	"testing"
	"github.com/stretchr/testify/require"
	"net/http"
)

func Test_Bind(t *testing.T) {
	t.Run("success", func(t *testing.T){

		cId := "CID-001"
		companyName := "Meli"
		address	:= "Dirección"
		telephone := "111-222333"
		locality_id := 1

		carrier := CarrierRequest {
			CId:      			&cId,
			CompanyName:        &companyName,
			Address:            &address,
			Telephone:          &telephone,
			LocalityId:         &locality_id,
		}

		err := carrier.Bind(&http.Request{})
		require.NoError(t, err)
	})

	t.Run("nil fields", func(t *testing.T) {

		cId := "CID-001"
		companyName := "Meli"
		address	:= "Dirección"
		telephone := "111-222333"

		carrier := CarrierRequest {
			CId:      			nil,
			CompanyName:        nil,
			Address:            nil,
			Telephone:          nil,
			LocalityId:         nil,
		}
		err := carrier.Bind(&http.Request{})
		require.Error(t, err)

		carrier = CarrierRequest {
			CId:      			&cId,
			CompanyName:        nil,
			Address:            nil,
			Telephone:          nil,
			LocalityId:         nil,
		}
		err = carrier.Bind(&http.Request{})
		require.Error(t, err)

		carrier = CarrierRequest {
			CId:      			&cId,
			CompanyName:        &companyName,
			Address:            nil,
			Telephone:          nil,
			LocalityId:         nil,
		}
		err = carrier.Bind(&http.Request{})
		require.Error(t, err)

		carrier = CarrierRequest {
			CId:      			&cId,
			CompanyName:        &companyName,
			Address:            &address,
			Telephone:          nil,
			LocalityId:         nil,
		}
		err = carrier.Bind(&http.Request{})
		require.Error(t, err)

		carrier = CarrierRequest {
			CId:      			&cId,
			CompanyName:        &companyName,
			Address:            &address,
			Telephone:          &telephone,
			LocalityId:         nil,
		}
		err = carrier.Bind(&http.Request{})
		require.Error(t, err)

		carrier = CarrierRequest {
			CId:      			&cId,
			CompanyName:        &companyName,
			Address:            &address,
			Telephone:          &telephone,
			LocalityId:         nil,
		}
		err = carrier.Bind(&http.Request{})
		require.Error(t, err)
	})
}
