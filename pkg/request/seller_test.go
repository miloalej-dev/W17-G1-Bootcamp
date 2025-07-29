package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSellerRequest_Bind(t *testing.T) {
	// Common values for all tests
	name := "ABC Company"
	address := "123 Main Street"
	telephone := "+1-555-123-4567"
	localityId := 1

	// Test case definitions
	tests := []struct {
		title         string
		request       *SellerRequest
		expectedError string
	}{
		{
			title: "Success - All fields valid",
			request: &SellerRequest{
				Name:       &name,
				Address:    &address,
				Telephone:  &telephone,
				LocalityId: &localityId,
			},
			expectedError: "",
		},
		{
			title: "Error - Missing Name",
			request: &SellerRequest{
				Name:       nil,
				Address:    &address,
				Telephone:  &telephone,
				LocalityId: &localityId,
			},
			expectedError: "name must not be null",
		},
		{
			title: "Error - Missing Address",
			request: &SellerRequest{
				Name:       &name,
				Address:    nil,
				Telephone:  &telephone,
				LocalityId: &localityId,
			},
			expectedError: "address must not be null",
		},
		{
			title: "Error - Missing Telephone",
			request: &SellerRequest{
				Name:       &name,
				Address:    &address,
				Telephone:  nil,
				LocalityId: &localityId,
			},
			expectedError: "telephone must not be null",
		},
		{
			title: "Error - Missing LocalityId",
			request: &SellerRequest{
				Name:       &name,
				Address:    &address,
				Telephone:  &telephone,
				LocalityId: nil,
			},
			expectedError: "locality_id must not be null",
		},
		{
			title: "Error - All fields missing",
			request: &SellerRequest{
				Name:       nil,
				Address:    nil,
				Telephone:  nil,
				LocalityId: nil,
			},
			expectedError: "name must not be null",
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
