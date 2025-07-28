package handler

import (
	"bytes"
	"encoding/json"
	"errors"
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

func (s *PurchaseOrderHandlerTestSuite) TestGetPurchaseOrdersReport_InvalidID() {
	expectedResponse := response.Response{
		Message:    "invalid", // Actualización clave aquí
		StatusCode: http.StatusBadRequest,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/report?id=invalid", s.path), nil)
	rec := httptest.NewRecorder()

	s.handler.GetPurchaseOrdersReport(rec, req)

	s.Equal(http.StatusBadRequest, rec.Code)
	s.JSONEq(string(expectedBody), rec.Body.String())
}

func (s *PurchaseOrderHandlerTestSuite) TestGetPurchaseOrdersReport_NotFound() {
	// Arrange
	id := 999
	expectedErr := errors.New("purchase orders not found")
	expectedResponse := response.Response{
		Message:    expectedErr.Error(),
		StatusCode: http.StatusNotFound,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("RetrieveByBuyer", id).Return([]models.PurchaseOrder{}, expectedErr) // ✅ CORREGIDO

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/report?id=%d", s.path, id), nil)
	rec := httptest.NewRecorder()

	// Act
	s.handler.GetPurchaseOrdersReport(rec, req)

	// Assert
	s.Equal(http.StatusNotFound, rec.Code)
	s.JSONEq(string(expectedBody), rec.Body.String())
}

func (s *PurchaseOrderHandlerTestSuite) TestGetPurchaseOrdersReport_WithoutID() {
	// Arrange
	id := 0
	expectedErr := errors.New("purchase orders not found")
	expectedResponse := response.Response{
		Message:    expectedErr.Error(),
		StatusCode: http.StatusNotFound,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("RetrieveByBuyer", id).Return([]models.PurchaseOrder{}, expectedErr)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/report", s.path), nil)
	rec := httptest.NewRecorder()

	// Act
	s.handler.GetPurchaseOrdersReport(rec, req)

	// Assert
	s.Equal(http.StatusNotFound, rec.Code)
	s.JSONEq(string(expectedBody), rec.Body.String())
}

func (s *PurchaseOrderHandlerTestSuite) TestPostPurchaseOrders_Success() {
	// Arrange
	payload := `{
		"order_number": "PO123",
		"order_date": "2025-07-28T00:00:00Z",
		"tracing_code": "ABC123",
		"buyer_id": 1,
		"warehouse_id": 1,
		"carrier_id": 2,
		"order_status_id": 1,
		"order_details": [{
			"quantity": 10,
			"clean_lines_status": "clean",
			"temperature": 4.5,
			"product_record_id": 1001
		}]
	}`

	expectedOrder := models.PurchaseOrder{
		OrderNumber:   "PO123",
		OrderDate:     time.Date(2025, 7, 28, 0, 0, 0, 0, time.UTC),
		TracingCode:   "ABC123",
		BuyerID:       1,
		WarehouseID:   1,
		CarrierID:     2,
		OrderStatusID: 1,
		OrderDetails: &[]models.OrderDetail{
			{
				Quantity:         10,
				CleanLinesStatus: "clean",
				Temperature:      4.5,
				ProductRecordID:  1001,
			},
		},
	}

	// Esperado como respuesta
	expectedResponse := response.Response{
		Data:       expectedOrder,
		Message:    "",
		StatusCode: http.StatusCreated,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("Register", mock.AnythingOfType("models.PurchaseOrder")).Return(expectedOrder, nil)

	req := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Act
	s.handler.PostPurchaseOrders(rec, req)

	// Assert
	s.Equal(http.StatusCreated, rec.Code)
	s.JSONEq(string(expectedBody), rec.Body.String())
}

func (s *PurchaseOrderHandlerTestSuite) TestPostPurchaseOrders_BindError() {
	// Arrange: JSON inválido (comilla faltante)
	payload := `{
		"order_number": "PO123,
		"order_date": "2025-07-28T00:00:00Z"
	}`

	expectedResponse := response.Response{
		Data:       nil,
		Message:    "invalid character '\\n' in string literal",
		StatusCode: http.StatusUnprocessableEntity,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	req := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Act
	s.handler.PostPurchaseOrders(rec, req)

	// Assert
	s.Equal(http.StatusUnprocessableEntity, rec.Code)
	s.JSONEq(string(expectedBody), rec.Body.String())
}

func (s *PurchaseOrderHandlerTestSuite) TestPostPurchaseOrders_RegisterConflict() {
	// Arrange
	payload := `{
		"order_number": "PO123",
		"order_date": "2025-07-28T00:00:00Z",
		"tracing_code": "ABC123",
		"buyer_id": 1,
		"warehouse_id": 1,
		"carrier_id": 2,
		"order_status_id": 1,
		"order_details": [{
			"quantity": 10,
			"clean_lines_status": "clean",
			"temperature": 4.5,
			"product_record_id": 1001
		}]
	}`

	expectedErr := errors.New("order number already exists")

	s.mock.On("Register", mock.AnythingOfType("models.PurchaseOrder")).Return(models.PurchaseOrder{}, expectedErr)

	expectedResponse := response.Response{
		Data:       nil,
		Message:    expectedErr.Error(),
		StatusCode: http.StatusConflict,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	req := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Act
	s.handler.PostPurchaseOrders(rec, req)

	// Assert
	s.Equal(http.StatusConflict, rec.Code)
	s.JSONEq(string(expectedBody), rec.Body.String())
}
