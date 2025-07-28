package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewBuyer(t *testing.T) {
	id := 10
	cardNumber := "1234-5678"
	firstName := "Juan"
	lastName := "Pérez"

	buyer := NewBuyer(id, cardNumber, firstName, lastName)

	require.NotNil(t, buyer)
	require.Equal(t, id, buyer.Id)
	require.Equal(t, cardNumber, buyer.CardNumberId)
	require.Equal(t, firstName, buyer.FirstName)
	require.Equal(t, lastName, buyer.LastName)
}
