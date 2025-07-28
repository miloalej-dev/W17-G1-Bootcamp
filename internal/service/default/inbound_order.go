package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"time"
)

type InboundOrderService struct {
	repository.InboundOrderRepository
}

func NewInboundOrderService(inboundOrderRepository repository.InboundOrderRepository) *InboundOrderService {
	return &InboundOrderService{InboundOrderRepository: inboundOrderRepository}
}

func (i *InboundOrderService) RetrieveAll() ([]models.InboundOrder, error) {
	return i.FindAll()
}

func (i *InboundOrderService) Retrieve(id int) (models.InboundOrder, error) {
	return i.FindById(id)
}

func (i *InboundOrderService) Register(inboundOrder models.InboundOrder) (models.InboundOrder, error) {

	if inboundOrder.OrderDate.IsZero() {
		inboundOrder.OrderDate = time.Now()
	}

	return i.Create(inboundOrder)
}

func (i *InboundOrderService) Modify(inboundOrder models.InboundOrder) (models.InboundOrder, error) {
	return i.Update(inboundOrder)
}

func (i *InboundOrderService) PartialModify(id int, fields map[string]any) (models.InboundOrder, error) {
	return i.PartialUpdate(id, fields)
}

func (i *InboundOrderService) Remove(id int) error {
	return i.Delete(id)
}
