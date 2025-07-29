package request

import (
	"testing"
	"github.com/stretchr/testify/require"
	"net/http"
)

func Test_Bind(t *testing.T) {

	cId := "CID-001"
	companyName := "Meli"
	address	:= "Dirección"
	telephone := "111-222333"
	locality_id := 1

	tests := []struct {
		title         string
		request       *CarrierRequest
		expectedError string
	}{
		{
			title: "Success - All fields valid",
			request: &CarrierRequest {
				CId:      			&cId,
				CompanyName:        &companyName,
				Address:            &address,
				Telephone:          &telephone,
				LocalityId:         &locality_id,
			},
			expectedError: "",
		},
		{
			title: "Error - missing CId",
			request: &CarrierRequest {
				CId:      			nil,
				CompanyName:        &companyName,
				Address:            &address,
				Telephone:          &telephone,
				LocalityId:         &locality_id,
			},
			expectedError: "cid code must not be null",
		},
		{
			title: "Error - missing company name",
			request: &CarrierRequest {
				CId:      			&cId,
				CompanyName:        nil,
				Address:            &address,
				Telephone:          &telephone,
				LocalityId:         &locality_id,
			},
			expectedError: "company name must not be null",
		},
		{
			title: "Error - missing address",
			request: &CarrierRequest {
				CId:      			&cId,
				CompanyName:        &companyName,
				Address:            nil,
				Telephone:          &telephone,
				LocalityId:         &locality_id,
			},
			expectedError: "address must not be null",
		},
		{
			title: "Error - missing telephone",
			request: &CarrierRequest {
				CId:      			&cId,
				CompanyName:        &companyName,
				Address:            &address,
				Telephone:          nil,
				LocalityId:         &locality_id,
			},
			expectedError: "telephone must not be null",
		},
		{
			title: "Error - missing locality Id",
			request: &CarrierRequest {
				CId:      			&cId,
				CompanyName:        &companyName,
				Address:            &address,
				Telephone:          &telephone,
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
