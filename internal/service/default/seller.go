package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type SellerService struct {
	repository repository.Repository[int, models.Seller]
}

func NewSellerService(repository repository.Repository[int, models.Seller]) *SellerService {
	return &SellerService{
		repository: repository,
	}
}

// GetSellers returns all sellers
func (s *SellerService) GetSellers() ([]models.Seller, error) {
	return s.repository.FindAll()
}

// GetSellerById returns a seller by id
func (s *SellerService) GetSellerById(id int) (models.Seller, error) {
	return s.repository.FindById(id)
}

// RegisterSeller creates a new seller
func (s *SellerService) RegisterSeller(seller models.Seller) (models.Seller, error) {
	return s.repository.Create(seller)
}

// ModifySeller updates an existing seller
func (s *SellerService) ModifySeller(seller models.Seller) (models.Seller, error) {
	return s.repository.Update(seller)
}

// UpdateSellerFields updates specific fields of an existing seller
func (s *SellerService) UpdateSellerFields(id int, fields map[string]interface{}) (models.Seller, error) {
	return s.repository.PartialUpdate(id, fields)
}

// RemoveSeller removes a seller by id
func (s *SellerService) RemoveSeller(id int) error {
	return s.repository.Delete(id)
}
