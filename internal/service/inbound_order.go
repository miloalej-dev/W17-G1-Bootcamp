package service

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type InboundOrderService interface {
	RetrieveAll() ([]models.InboundOrder, error)
	Retrieve(id int) (models.InboundOrder, error)
	Register(inboundOrder models.InboundOrder) (models.InboundOrder, error)
	Modify(inboundOrder models.InboundOrder) (models.InboundOrder, error)
	PartialModify(id int, fields map[string]any) (models.InboundOrder, error)
	Remove(id int) error
}
