package request

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_BuyerBind(t *testing.T) {

	cardNumberId := "1"
	firstName := "john"
	lastName := "smith"

	tests := []struct {
		title         string
		request       *BuyerRequest
		expectedError string
	}{
		{
			title: "Success - All fields valid",
			request: &BuyerRequest{
				CardNumberId: &cardNumberId,
				FirstName:    &firstName,
				LastName:     &lastName,
			},
			expectedError: "",
		},
		{
			title: "Error - Missing Card Number Id",
			request: &BuyerRequest{
				CardNumberId: nil,
				FirstName:    &firstName,
				LastName:     &lastName,
			},
			expectedError: "card number Id must be not null",
		},
		{
			title: "Error - Missing First Name",
			request: &BuyerRequest{
				CardNumberId: &cardNumberId,
				FirstName:    nil,
				LastName:     &lastName,
			},
			expectedError: "first name must not be null",
		},
		{
			title: "Error - Missing Last Name",
			request: &BuyerRequest{
				CardNumberId: &cardNumberId,
				FirstName:    &firstName,
				LastName:     nil,
			},
			expectedError: "last name must not be null",
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
