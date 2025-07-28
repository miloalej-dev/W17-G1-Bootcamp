package handler

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PurchaseOrderServiceMock struct {
	mock.Mock
}

type PurchaseOrderTestSuite struct {
	suite.Suite
	mock    *PurchaseOrderServiceMock
	handler *PurchaseOrderHandler
	path    string
}

// Mock methods for PurchaseOrder
func (s *PurchaseOrderServiceMock) RetrieveAll() ([]models.PurchaseOrder, error) {
	args := s.Called()
	return args.Get(0).([]models.PurchaseOrder), args.Error(1)
}

func (s *PurchaseOrderServiceMock) Retrieve(id int) (models.PurchaseOrder, error) {
	args := s.Called(id)
	return args.Get(0).(models.PurchaseOrder), args.Error(1)
}

func (s *PurchaseOrderServiceMock) Register(po models.PurchaseOrder) (models.PurchaseOrder, error) {
	args := s.Called(po)
	return args.Get(0).(models.PurchaseOrder), args.Error(1)
}

func (s *PurchaseOrderServiceMock) Modify(po models.PurchaseOrder) (models.PurchaseOrder, error) {
	args := s.Called(po)
	return args.Get(0).(models.PurchaseOrder), args.Error(1)
}

func (s *PurchaseOrderServiceMock) PartialModify(id int, fields map[string]any) (models.PurchaseOrder, error) {
	args := s.Called(id, fields)
	return args.Get(0).(models.PurchaseOrder), args.Error(1)
}

func (s *PurchaseOrderServiceMock) Remove(id int) error {
	args := s.Called(id)
	return args.Error(0)
}
func (s *PurchaseOrderServiceMock) RetrieveByBuyer(id int) ([]models.PurchaseOrder, error) {
	args := s.Called(id)
	return args.Get(0).([]models.PurchaseOrder), args.Error(1)
}

func (s *PurchaseOrderTestSuite) SetupTest() {
	s.mock = new(PurchaseOrderServiceMock)
	s.handler = NewPurchaseOrderDefault(s.mock)
	s.path = "/api/v1/purchase_orders"
}

// Run the test suite
func TestSellerHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(PurchaseOrderTestSuite))
}
