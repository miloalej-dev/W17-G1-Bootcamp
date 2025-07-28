package handler

import (
	"encoding/json"
	"fmt"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type PurchaseOrderServiceMock struct {
	mock.Mock
}

type PurchaseOrderHandlerTestSuite struct {
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

func (s *PurchaseOrderHandlerTestSuite) SetupTest() {
	s.mock = new(PurchaseOrderServiceMock)
	s.handler = NewPurchaseOrderDefault(s.mock)
	s.path = "/api/v1/purchase_orders"
}

// Run the test suite
func TestSellerHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(PurchaseOrderHandlerTestSuite))
}

func (s *PurchaseOrderHandlerTestSuite) TestGetPurchaseOrdersReport_Ok() {
	// Arrange
	id := 1
	expectedData := []models.PurchaseOrder{
		{
			Id:            1,
			OrderNumber:   "PO123",
			OrderDate:     time.Now(),
			TracingCode:   "ABC123",
			BuyerID:       id,
			WarehouseID:   1,
			CarrierID:     2,
			OrderStatusID: 1,
			OrderDetails: &[]models.OrderDetail{
				{
					Id:               1,
					Quantity:         10,
					CleanLinesStatus: "clean",
					Temperature:      4.5,
					ProductRecordID:  1001,
					PurchaseOrderID:  1,
				},
			},
		},
	}

	expectedResponse := response.Response{Data: expectedData}
	expectedBody, _ := json.Marshal(expectedResponse)

	// Mock del servicio
	s.mock.On("RetrieveByBuyer", id).Return(expectedData, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/report?id=%d", s.path, id), nil)
	rec := httptest.NewRecorder()

	// Act
	s.handler.GetPurchaseOrdersReport(rec, req)

	// Assert
	s.Equal(http.StatusOK, rec.Code)
	s.JSONEq(string(expectedBody), rec.Body.String())
}
