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
)

type WarehouseServiceMock struct {
	mock.Mock
}

type WarehouseHandlerTestSuite struct {
	suite.Suite
	mock    *WarehouseServiceMock
	handler *WarehouseDefault
	path    string
}

// Mock methods for WarehouseService

func (s *WarehouseServiceMock) RetrieveAll() ([]models.Warehouse, error) {
	args := s.Called()
	return args.Get(0).([]models.Warehouse), args.Error(1)
}

func (s *WarehouseServiceMock) Retrieve(id int) (models.Warehouse, error) {
	args := s.Called(id)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (s *WarehouseServiceMock) Register(warehouse models.Warehouse) (models.Warehouse, error) {
	args := s.Called(warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (s *WarehouseServiceMock) Modify(warehouse models.Warehouse) (models.Warehouse, error) {
	args := s.Called(warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (s *WarehouseServiceMock) PartialModify(id int, fields map[string]any) (models.Warehouse, error) {
	args := s.Called(id, fields)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (s *WarehouseServiceMock) Remove(id int) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *WarehouseHandlerTestSuite) SetupTest() {
	s.mock = new(WarehouseServiceMock)
	s.handler = NewWarehouseDefault(s.mock)
	s.path = "/api/v1/warehouses"
}

// Test cases for WarehouseHandler

func (s *WarehouseHandlerTestSuite) TestGetWarehouses_Ok() {
	// Arrange
	var expectedBody []byte
	expectedWarehouses := []models.Warehouse{
		{Id:1, WarehouseCode:"AAA-111", Address:"Boulevard", Telephone:"123-456789", MinimumCapacity:10, MinimumTemperature:10, LocalityId:1},
		{Id:2, WarehouseCode:"AAA-222", Address:"Plaza", Telephone:"223-456789", MinimumCapacity:20, MinimumTemperature:20, LocalityId:1},
	}


	expectedResponse := response.Response{Data: expectedWarehouses}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("RetrieveAll").Return(expectedWarehouses, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetWarehouses(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.Len(resp.Data, 2)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestGetWarehouses_InternalError() {
	// Arrange
	var expectedBody []byte
	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("RetrieveAll").Return([]models.Warehouse{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetWarehouses(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

//// GetWarehouse tests
func (s *WarehouseHandlerTestSuite) TestGetWarehouse_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedWarehouse := models.Warehouse{
		Id:1, WarehouseCode:"AAA-111", Address:"Boulevard", Telephone:"123-456789", MinimumCapacity:10, MinimumTemperature:10, LocalityId:1,
	}
	expectedResponse := response.Response{Data: expectedWarehouse}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("Retrieve", id).Return(expectedWarehouse, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestGetWarehouse_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", id)
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestGetWarehouse_BadRequest_NegativeId() {
	// Arrange
	var expectedBody []byte
	id := -1
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestGetWarehouse_NotFound() {
	// Arrange
	var expectedBody []byte
	id := 999
	expectedError := errors.New("entity not found")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNotFound}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("Retrieve", id).Return(models.Warehouse{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", s.path, id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "999")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PostWarehouse tests
func (s *WarehouseHandlerTestSuite) TestPostWarehouse_Ok() {
	// Arrange
	var expectedBody []byte
	warehouseCode := "AAA-111"
	address := "Boulevard"
	telephone := "123-456789"
	minimumCapacity := 10
	minimumTemperature := 10
	localityId := 1

	requestBody := map[string]interface{}{
		"warehouse_code":		warehouseCode,
		"address":				address,
		"telephone":			telephone,
		"minimum_capacity":		minimumCapacity,
		"minimum_temperature":	minimumTemperature,
		"locality_id":			localityId,
	}

	expectedWarehouse := models.Warehouse{
		Id: 1, WarehouseCode: warehouseCode, Address: address, Telephone: telephone, MinimumCapacity: minimumCapacity, MinimumTemperature: minimumTemperature, LocalityId: localityId,
	}
	expectedResponse := response.Response{Data: expectedWarehouse}
	expectedBody, _ = json.Marshal(expectedResponse)

	inputWarehouse := models.Warehouse{
		WarehouseCode: warehouseCode, Address: address, Telephone: telephone, MinimumCapacity: minimumCapacity, MinimumTemperature: minimumTemperature, LocalityId: localityId,
	}

	s.mock.On("Register", inputWarehouse).Return(expectedWarehouse, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusCreated, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestPostWarehouse_BadRequest_MissingCode() {
	// Arrange
	var expectedBody []byte
	address := "Boulevard"
	telephone := "123-456789"
	minimumCapacity := 10
	minimumTemperature := 10
	localityId := 1
	requestBody := map[string]interface{}{
		"address":				address,
		"telephone":			telephone,
		"minimum_capacity":		minimumCapacity,
		"minimum_temperature":	minimumTemperature,
		"locality_id":			localityId,
	}
	expectedResponse := response.Response{Message: "warehouse code must not be null", StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestPostWarehouse_InternalError() {
	// Arrange
	var expectedBody []byte
	warehouseCode := "AAA-111"
	address := "Boulevard"
	telephone := "123-456789"
	minimumCapacity := 10
	minimumTemperature := 10
	localityId := 1

	requestBody := map[string]interface{}{
		"warehouse_code":		warehouseCode,
		"address":				address,
		"telephone":			telephone,
		"minimum_capacity":		minimumCapacity,
		"minimum_temperature":	minimumTemperature,
		"locality_id":			localityId,
	}


	expectedError := errors.New("internal error")

	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	inputWarehouse := models.Warehouse{
		WarehouseCode: warehouseCode, Address: address, Telephone: telephone, MinimumCapacity: minimumCapacity, MinimumTemperature: minimumTemperature, LocalityId: localityId,
	}

	s.mock.On("Register", inputWarehouse).Return(models.Warehouse{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PutWarehouse tests
func (s *WarehouseHandlerTestSuite) TestPutWarehouse_Ok() {
	// Arrange
	var expectedBody []byte
	warehouseId := 1
	warehouseCode := "AAA-111"
	address := "Boulevard"
	telephone := "123-456789"
	minimumCapacity := 10
	minimumTemperature := 10
	localityId := 1

	requestBody := map[string]interface{}{
		"warehouse_code":		warehouseCode,
		"address":				address,
		"telephone":			telephone,
		"minimum_capacity":		minimumCapacity,
		"minimum_temperature":	minimumTemperature,
		"locality_id":			localityId,
	}


	expectedWarehouse := models.Warehouse{
		Id: warehouseId, WarehouseCode: warehouseCode, Address: address, Telephone: telephone, MinimumCapacity: minimumCapacity, MinimumTemperature: minimumTemperature, LocalityId: localityId,
	}
	expectedResponse := response.Response{Data: expectedWarehouse}
	expectedBody, _ = json.Marshal(expectedResponse)

	inputWarehouse := models.Warehouse{
		Id: warehouseId, WarehouseCode: warehouseCode, Address: address, Telephone: telephone, MinimumCapacity: minimumCapacity, MinimumTemperature: minimumTemperature, LocalityId: localityId,
	}

	s.mock.On("Modify", inputWarehouse).Return(expectedWarehouse, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", s.path, warehouseId), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestPutWarehouse_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	warehouseCode := "AAA-111"
	address := "Boulevard"
	telephone := "123-456789"
	minimumCapacity := 10
	minimumTemperature := 10
	localityId := 1

	requestBody := map[string]interface{}{
		"warehouse_code":		warehouseCode,
		"address":				address,
		"telephone":			telephone,
		"minimum_capacity":		minimumCapacity,
		"minimum_temperature":	minimumTemperature,
		"locality_id":			localityId,
	}
	requestBodyBytes, _ := json.Marshal(requestBody)

	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	request := httptest.NewRequest(http.MethodPut, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", id)
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestPutWarehouse_BadRequest_NegativeId() {
	// Arrange
	var expectedBody []byte
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	id := -1
	warehouseCode := "AAA-111"
	address := "Boulevard"
	telephone := "123-456789"
	minimumCapacity := 10
	minimumTemperature := 10
	localityId := 1

	requestBody := map[string]interface{}{
		"warehouse_code":		warehouseCode,
		"address":				address,
		"telephone":			telephone,
		"minimum_capacity":		minimumCapacity,
		"minimum_temperature":	minimumTemperature,
		"locality_id":			localityId,
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
	s.handler.PutWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestPutWarehouse_BadRequest() {
	// Arrange
	var expectedBody []byte
	warehouseId := 1
	address := "Boulevard"
	telephone := "123-456789"
	minimumCapacity := 10
	minimumTemperature := 10
	localityId := 1

	requestBody := map[string]interface{}{
		"address":				address,
		"telephone":			telephone,
		"minimum_capacity":		minimumCapacity,
		"minimum_temperature":	minimumTemperature,
		"locality_id":			localityId,
	}
	expectedResponse := response.Response{Message: "warehouse code must not be null", StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", s.path, warehouseId), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestPutWarehouse_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	warehouseCode := "111-AAA"
	address := "Boulevard"
	telephone := "123-456789"
	minimumCapacity := 10
	minimumTemperature := 10
	localityId := 1

	requestBody := map[string]interface{}{
		"warehouse_code":		warehouseCode,
		"address":				address,
		"telephone":			telephone,
		"minimum_capacity":		minimumCapacity,
		"minimum_temperature":	minimumTemperature,
		"locality_id":			localityId,
	}

	expectedError := errors.New("internal error")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	inputWarehouse := models.Warehouse{
		Id: id, WarehouseCode: warehouseCode, Address: address, Telephone: telephone, MinimumCapacity: minimumCapacity, MinimumTemperature: minimumTemperature, LocalityId: localityId,
	}

	s.mock.On("Modify", inputWarehouse).Return(models.Warehouse{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PatchWarehouse tests
func (s *WarehouseHandlerTestSuite) TestPatchWarehouse_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	warehouseCode := "111-AAA"
	address := "Boulevard"
	telephone := "123-456789"
	minimumCapacity := 10
	minimumTemperature := 11
	localityId := 1
	fields := map[string]interface{}{
		"minimum_temperature":	11.0,
	}

	expectedWarehouse := models.Warehouse{
		Id: id, WarehouseCode: warehouseCode, Address: address, Telephone: telephone, MinimumCapacity: minimumCapacity, MinimumTemperature: minimumTemperature, LocalityId: localityId,
	}
	expectedResponse := response.Response{Data: expectedWarehouse}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("PartialModify", id, fields).Return(expectedWarehouse, nil)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestPatchWarehouse_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	fields := map[string]interface{}{
		"warehouse_code": "Partially Updated Company",
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
	s.handler.PatchWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestPatchWarehouse_BadRequest_NegativeId() {
	// Arrange
	var expectedBody []byte
	id := -1
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	fields := map[string]interface{}{
		"address": "Partially Updated Company",
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
	s.handler.PatchWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestPatchWarehouse_BadRequest() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedResponse := response.Response{Message: ErrUnexpectedJSON.Error(), StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer([]byte("invalid json")))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestPatchWarehouse_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	fields := map[string]interface{}{
		"address": "Partially Updated Company",
	}

	expectedError := errors.New("internal error")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("PartialModify", id, fields).Return(models.Warehouse{}, expectedError)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// DeleteWarehouse tests
func (s *WarehouseHandlerTestSuite) TestDeleteWarehouse_NoContent() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedResponse := response.Response{Data: nil}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("Remove", id).Return(nil)

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusNoContent, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestDeleteWarehouse_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", id)
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *WarehouseHandlerTestSuite) TestDeleteWarehouse_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedError := errors.New("internal error")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("Remove", id).Return(expectedError)

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteWarehouse(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Run the test suite
func TestWarehouseHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(WarehouseHandlerTestSuite))
}
