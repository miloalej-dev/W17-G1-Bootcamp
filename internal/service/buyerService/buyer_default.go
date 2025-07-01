package buyerService

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/buyerRepository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// NewBuyerDefault is a function that returns a new instance of BuyerDefault
func NewBuyerDefault(rp buyerRepository.BuyerRepository) *BuyerDefault {
	return &BuyerDefault{rp: rp}
}

// BuyerDefault is a struct that represents the default service for buyers
type BuyerDefault struct {
	// rp is the repository that will be used by the service
	rp buyerRepository.BuyerRepository
}

// FindAll is a method that returns a map of all buyers
func (s *BuyerDefault) FindAll() (v map[int]models.Buyer, err error) {
	v, err = s.rp.FindAll()
	return
}
