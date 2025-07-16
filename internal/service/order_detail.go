package service

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type OrderDetailService interface {
	RetrieveAll() (v []models.OrderDetail, err error)
	Retrieve(id int) (models.OrderDetail, error)
	Register(s models.OrderDetail) (models.OrderDetail, error)
	Modify(s models.OrderDetail) (models.OrderDetail, error)
	PartialModify(id int, fields map[string]any) (models.OrderDetail, error)
	Remove(id int) error
}
