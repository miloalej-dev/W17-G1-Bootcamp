package repository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// SellerRepository defines the interface for seller data operations
type SellerRepository interface {
	Repository[int, models.Seller]
	// More methods specific to seller data operations can be added here...
}
