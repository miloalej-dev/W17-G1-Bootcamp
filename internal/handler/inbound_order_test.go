package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

type InboundOrderServiceMock struct {
	mock.Mock
}

type InboundOrderHandlerTestSuite struct {
	suite.Suite
	mock    *InboundOrderServiceMock
	handler *InboundOrderHandler
	path    string
}

// Mock methods for InboundOrderService

func (m *InboundOrderServiceMock) RetrieveAll() ([]models.InboundOrder, error) {
	args := m.Called()
	return args.Get(0).([]models.InboundOrder), args.Error(1)
}

func (m *InboundOrderServiceMock) Retrieve(id int) (models.InboundOrder, error) {
	args := m.Called(id)
	return args.Get(0).(models.InboundOrder), args.Error(1)
}

func (m *InboundOrderServiceMock) Register(inboundOrder models.InboundOrder) (models.InboundOrder, error) {
	args := m.Called(inboundOrder)
	return args.Get(0).(models.InboundOrder), args.Error(1)
}

func (m *InboundOrderServiceMock) Modify(inboundOrder models.InboundOrder) (models.InboundOrder, error) {
	args := m.Called(inboundOrder)
	return args.Get(0).(models.InboundOrder), args.Error(1)
}

func (m *InboundOrderServiceMock) PartialModify(id int, fields map[string]any) (models.InboundOrder, error) {
	args := m.Called(id, fields)
	return args.Get(0).(models.InboundOrder), args.Error(1)
}

func (m *InboundOrderServiceMock) Remove(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (s *InboundOrderHandlerTestSuite) SetupTest() {
	s.mock = new(InboundOrderServiceMock)
	s.handler = NewInboundOrderHandler(s.mock)
	s.path = "/api/v1/inbound-orders"
}

// Test cases for InboundOrderHandler

// GetInboundOrders tests
func (s *InboundOrderHandlerTestSuite) TestGetInboundOrders_Ok() {
	// Arrange
	var expectedBody []byte
	orderDate := time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC)
	expectedInboundOrders := []models.InboundOrder{
		{Id: 1, OrderNumber: "ORD001", OrderDate: orderDate, EmployeeId: 1, ProductBatchId: 1, WarehouseId: 1},
		{Id: 2, OrderNumber: "ORD002", OrderDate: orderDate, EmployeeId: 2, ProductBatchId: 2, WarehouseId: 2},
	}

	expectedResponse := response.Response{Data: expectedInboundOrders}

	s.mock.On("RetrieveAll").Return(expectedInboundOrders, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetInboundOrders(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.Len(resp.Data, 2)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestGetInboundOrders_InternalError() {
	// Arrange
	var expectedBody []byte
	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("RetrieveAll").Return([]models.InboundOrder{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetInboundOrders(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// GetInboundOrder tests
func (s *InboundOrderHandlerTestSuite) TestGetInboundOrder_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	orderDate := time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC)
	expectedInboundOrder := models.InboundOrder{
		Id: id, OrderNumber: "ORD001", OrderDate: orderDate, EmployeeId: 1, ProductBatchId: 1, WarehouseId: 1}
	expectedResponse := response.Response{Data: expectedInboundOrder}

	s.mock.On("Retrieve", id).Return(expectedInboundOrder, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestGetInboundOrder_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", id)
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestGetInboundOrder_BadRequest_NegativeId() {
	// Arrange
	var expectedBody []byte
	id := -1
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestGetInboundOrder_NotFound() {
	// Arrange
	var expectedBody []byte
	id := 999
	expectedError := errors.New("entity not found")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNotFound}

	s.mock.On("Retrieve", id).Return(models.InboundOrder{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "999")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PostInboundOrder tests
func (s *InboundOrderHandlerTestSuite) TestPostInboundOrder_Created() {
	// Arrange
	var expectedBody []byte
	orderNumber := "ORD003"
	employeeId := 3
	productBatchId := 3
	warehouseId := 3

	requestBody := map[string]interface{}{
		"order_number":     orderNumber,
		"employee_id":      employeeId,
		"product_batch_id": productBatchId,
		"warehouse_id":     warehouseId,
	}

	orderDate := time.Now()
	expectedInboundOrder := models.InboundOrder{
		Id: 1, OrderNumber: orderNumber, OrderDate: orderDate, EmployeeId: employeeId, ProductBatchId: productBatchId, WarehouseId: warehouseId}
	expectedResponse := response.Response{Data: expectedInboundOrder}

	inputInboundOrder := models.InboundOrder{
		OrderNumber: orderNumber, EmployeeId: employeeId, ProductBatchId: productBatchId, WarehouseId: warehouseId}

	s.mock.On("Register", inputInboundOrder).Return(expectedInboundOrder, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusCreated, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestPostInboundOrder_BadRequest_MissingOrderNumber() {
	// Arrange
	var expectedBody []byte
	requestBody := map[string]interface{}{
		"employee_id":      3,
		"product_batch_id": 3,
		"warehouse_id":     3,
	}
	expectedResponse := response.Response{Message: "order_number must not be null", StatusCode: http.StatusBadRequest}

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestPostInboundOrder_InternalError() {
	// Arrange
	var expectedBody []byte
	orderNumber := "ORD003"
	employeeId := 3
	productBatchId := 3
	warehouseId := 3

	requestBody := map[string]interface{}{
		"order_number":     orderNumber,
		"employee_id":      employeeId,
		"product_batch_id": productBatchId,
		"warehouse_id":     warehouseId,
	}

	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	inputInboundOrder := models.InboundOrder{
		OrderNumber: orderNumber, EmployeeId: employeeId, ProductBatchId: productBatchId, WarehouseId: warehouseId}

	s.mock.On("Register", inputInboundOrder).Return(models.InboundOrder{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PutInboundOrder tests
func (s *InboundOrderHandlerTestSuite) TestPutInboundOrder_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	orderNumber := "ORD001-UPDATED"
	employeeId := 4
	productBatchId := 4
	warehouseId := 4

	requestBody := map[string]interface{}{
		"order_number":     orderNumber,
		"employee_id":      employeeId,
		"product_batch_id": productBatchId,
		"warehouse_id":     warehouseId,
	}

	orderDate := time.Now()
	expectedInboundOrder := models.InboundOrder{
		Id: id, OrderNumber: orderNumber, OrderDate: orderDate, EmployeeId: employeeId, ProductBatchId: productBatchId, WarehouseId: warehouseId}
	expectedResponse := response.Response{Data: expectedInboundOrder}

	inputInboundOrder := models.InboundOrder{
		Id: id, OrderNumber: orderNumber, EmployeeId: employeeId, ProductBatchId: productBatchId, WarehouseId: warehouseId}

	s.mock.On("Modify", inputInboundOrder).Return(expectedInboundOrder, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestPutInboundOrder_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	requestBody := map[string]interface{}{
		"order_number":     "ORD001-UPDATED",
		"employee_id":      4,
		"product_batch_id": 4,
		"warehouse_id":     4,
	}

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", id)
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestPutInboundOrder_BadRequest_NegativeId() {
	// Arrange
	var expectedBody []byte
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	id := -1
	requestBody := map[string]interface{}{
		"order_number":     "ORD001-UPDATED",
		"employee_id":      4,
		"product_batch_id": 4,
		"warehouse_id":     4,
	}

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestPutInboundOrder_BadRequest() {
	// Arrange
	var expectedBody []byte
	id := 1
	requestBody := map[string]interface{}{
		"employee_id":      4,
		"product_batch_id": 4,
		"warehouse_id":     4,
	}
	expectedResponse := response.Response{Message: "order_number must not be null", StatusCode: http.StatusBadRequest}

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestPutInboundOrder_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	orderNumber := "ORD001-UPDATED"
	employeeId := 4
	productBatchId := 4
	warehouseId := 4

	requestBody := map[string]interface{}{
		"order_number":     orderNumber,
		"employee_id":      employeeId,
		"product_batch_id": productBatchId,
		"warehouse_id":     warehouseId,
	}

	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	inputInboundOrder := models.InboundOrder{
		Id: id, OrderNumber: orderNumber, EmployeeId: employeeId, ProductBatchId: productBatchId, WarehouseId: warehouseId}

	s.mock.On("Modify", inputInboundOrder).Return(models.InboundOrder{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PatchInboundOrder tests
func (s *InboundOrderHandlerTestSuite) TestPatchInboundOrder_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	fields := map[string]interface{}{
		"order_number": "ORD001-PARTIAL-UPDATE",
	}

	orderDate := time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC)
	expectedInboundOrder := models.InboundOrder{
		Id: id, OrderNumber: "ORD001-PARTIAL-UPDATE", OrderDate: orderDate, EmployeeId: 1, ProductBatchId: 1, WarehouseId: 1}
	expectedResponse := response.Response{Data: expectedInboundOrder}

	s.mock.On("PartialModify", id, fields).Return(expectedInboundOrder, nil)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestPatchInboundOrder_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	fields := map[string]interface{}{
		"order_number": "ORD001-PARTIAL-UPDATE",
	}

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", id)
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestPatchInboundOrder_BadRequest_NegativeId() {
	// Arrange
	var expectedBody []byte
	id := -1
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	fields := map[string]interface{}{
		"order_number": "ORD001-PARTIAL-UPDATE",
	}

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestPatchInboundOrder_BadRequest() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedResponse := response.Response{Message: ErrUnexpectedJSON.Error(), StatusCode: http.StatusBadRequest}

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer([]byte("invalid json")))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestPatchInboundOrder_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	fields := map[string]interface{}{
		"order_number": "ORD001-PARTIAL-UPDATE",
	}

	expectedError := errors.New("partial update failed")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("PartialModify", id, fields).Return(models.InboundOrder{}, expectedError)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// DeleteInboundOrder tests
func (s *InboundOrderHandlerTestSuite) TestDeleteInboundOrder_NoContent() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedResponse := response.Response{Data: nil}

	s.mock.On("Remove", id).Return(nil)

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusNoContent, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestDeleteInboundOrder_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", id)
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *InboundOrderHandlerTestSuite) TestDeleteInboundOrder_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedError := errors.New("delete failed")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("Remove", id).Return(expectedError)

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteInboundOrder(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Run the test suite
func TestInboundOrderHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(InboundOrderHandlerTestSuite))
}
