package service

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type SellerService interface {
	RetrieveAll() ([]models.Seller, error)
	Retrieve(id int) (models.Seller, error)
	Register(seller models.Seller) (models.Seller, error)
	Modify(seller models.Seller) (models.Seller, error)
	PartialModify(id int, fields map[string]any) (models.Seller, error)
	Remove(id int) error
}
