package repository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// BuyerRepository is an interface that represents a Buyer repository
type BuyerRepository interface {
	Repository[int, models.Seller]
}
