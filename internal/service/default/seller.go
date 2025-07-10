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

func (s *SellerService) RetrieveAll() ([]models.Seller, error) {
	return s.repository.FindAll()
}

func (s *SellerService) Retrieve(id int) (models.Seller, error) {
	return s.repository.FindById(id)
}

func (s *SellerService) Register(seller models.Seller) (models.Seller, error) {
	return s.repository.Create(seller)
}

func (s *SellerService) Modify(seller models.Seller) (models.Seller, error) {
	return s.repository.Update(seller)
}

func (s *SellerService) PartialModify(id int, fields map[string]any) (models.Seller, error) {
	return s.repository.PartialUpdate(id, fields)
}

func (s *SellerService) Remove(id int) error {
	return s.repository.Delete(id)
}
