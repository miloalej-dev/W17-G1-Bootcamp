package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type ProductRecordServiceMock struct {
	mock.Mock
}

type ProductRecordHandlerTestSuite struct {
	suite.Suite
	mock    *ProductRecordServiceMock
	handler *ProductRecordHandler
	path    string
}

// Mock methods for BuyerService

func (s *ProductRecordServiceMock) RetrieveAll() ([]models.ProductRecord, error) {
	args := s.Called()
	return args.Get(0).([]models.ProductRecord), args.Error(1)
}

func (s *ProductRecordServiceMock) Retrieve(id int) (models.ProductRecord, error) {
	args := s.Called(id)
	return args.Get(0).(models.ProductRecord), args.Error(1)
}

func (s *ProductRecordServiceMock) Register(productRecord models.ProductRecord) (models.ProductRecord, error) {
	args := s.Called(productRecord)
	return args.Get(0).(models.ProductRecord), args.Error(1)
}

func (s *ProductRecordServiceMock) Modify(productRecord models.ProductRecord) (models.ProductRecord, error) {
	args := s.Called(productRecord)
	return args.Get(0).(models.ProductRecord), args.Error(1)
}

func (s *ProductRecordServiceMock) PartialModify(id int, fields map[string]any) (models.ProductRecord, error) {
	args := s.Called(id, fields)
	return args.Get(0).(models.ProductRecord), args.Error(1)
}

func (s *ProductRecordServiceMock) Remove(id int) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *ProductRecordHandlerTestSuite) SetupTest() {
	s.mock = new(ProductRecordServiceMock)
	s.handler = NewProductRecordHandler(s.mock)
	s.path = "/api/v1/productRecords"
}

// Test cases for BuyerHandler

func (s *ProductRecordHandlerTestSuite) TestGetProductRecords_Success() {
	// Arrange

	expectedProductRecord := []models.ProductRecord{
		{
			Id:            1,
			LastUpdate:    "2022-22-10",
			PurchasePrice: 4.99,
			SalePrice:     5.99,
			ProductId:     1,
		},
		{
			Id:            2,
			LastUpdate:    "2022-22-11",
			PurchasePrice: 4.99,
			SalePrice:     6.99,
			ProductId:     1,
		},
	}

	expectedData := response.Response{Data: expectedProductRecord, StatusCode: http.StatusOK}

	s.mock.On("RetrieveAll").Return(expectedProductRecord, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetProductRecords(recorder, request)

	expectedResponse, _ := json.Marshal(expectedData)

	// Assert

	s.Equal(expectedData.StatusCode, recorder.Code)
	s.JSONEq(string(expectedResponse), recorder.Body.String())
}

func (s *ProductRecordHandlerTestSuite) TestGetProductRecord_InternalError() {
	// Arrange
	var expectedBody []byte
	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("RetrieveAll").Return([]models.ProductRecord{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetProductRecords(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// GetBuyer tests
func (s *ProductRecordHandlerTestSuite) TestGetProductRecord_Success() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedProductRecord := models.ProductRecord{
		Id:            1,
		LastUpdate:    "2022-22-10",
		PurchasePrice: 4.99,
		SalePrice:     5.99,
		ProductId:     1,
	}
	expectedResponse := response.Response{Data: expectedProductRecord, StatusCode: http.StatusOK}

	s.mock.On("Retrieve", id).Return(expectedProductRecord, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetProductRecord(recorder, request)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *ProductRecordHandlerTestSuite) TestGetProductRecord_BadRequest() {
	// Arrange
	var expectedBody []byte
	id := -1
	expectedError := errors.New("invalid request")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusBadRequest}

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetProductRecord(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *ProductRecordHandlerTestSuite) TestGetProductRecord_NotFound() {
	// Arrange
	var expectedBody []byte
	id := 999

	expectedError := errors.New("entity not found")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNotFound}

	s.mock.On("Retrieve", id).Return(models.ProductRecord{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", s.path, id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "999")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetProductRecord(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PostBuyer tests
func (s *ProductRecordHandlerTestSuite) TestPostProductRecord_Success() {
	// Arrange
	var expectedBody []byte

	requestBody := map[string]interface{}{

		"last_update":    "2022-22-10",
		"purchase_price": 4.99,
		"sale_price":     5.99,
		"product_id":     1,
	}
	expectedProductRecord := models.ProductRecord{
		Id:            1,
		LastUpdate:    "2022-22-10",
		PurchasePrice: 4.99,
		SalePrice:     5.99,
		ProductId:     1,
	}

	expectedResponse := response.Response{Data: expectedProductRecord, StatusCode: http.StatusCreated}

	inputProductRecord := models.ProductRecord{
		LastUpdate: "2022-22-10", PurchasePrice: 4.99, SalePrice: 5.99, ProductId: 1}

	s.mock.On("Register", inputProductRecord).Return(expectedProductRecord, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostProductRecord(recorder, request)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *ProductRecordHandlerTestSuite) TestPostProductRecord_BadRequest() {
	// Arrange
	var expectedBody []byte

	requestBody := map[string]interface{}{
		"last_update":    "2022-22-10",
		"purchase_price": 4.99,
	}

	expectedResponse := response.Response{Message: "sale price must be not null and greater than 0", StatusCode: http.StatusUnprocessableEntity}

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostProductRecord(recorder, request)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *ProductRecordHandlerTestSuite) TestPostProductRecord_InternalError() {
	// Arrange

	var expectedBody []byte

	requestBody := map[string]interface{}{
		"last_update":    "2022-22-10",
		"purchase_price": 4.99,
		"sale_price":     5.99,
		"product_id":     1,
	}

	expectedError := errors.New("something went wrong")

	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	inputProductRecord := models.ProductRecord{
		LastUpdate: "2022-22-10", PurchasePrice: 4.99, SalePrice: 5.99, ProductId: 1}

	s.mock.On("Register", inputProductRecord).Return(models.ProductRecord{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostProductRecord(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *ProductRecordHandlerTestSuite) TestPostProductRecord_Conflict() {
	// Arrange

	var expectedBody []byte

	requestBody := map[string]interface{}{
		"last_update":    "2022-22-10",
		"purchase_price": 4.99,
		"sale_price":     5.99,
		"product_id":     999,
	}

	expectedError := service.ErrProductIdConflict

	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusConflict}

	inputProductRecord := models.ProductRecord{
		LastUpdate: "2022-22-10", PurchasePrice: 4.99, SalePrice: 5.99, ProductId: 999}

	s.mock.On("Register", inputProductRecord).Return(models.ProductRecord{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostProductRecord(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PatchBuyer tests
func (s *ProductRecordHandlerTestSuite) TestPatchProductRecord_Success() {
	// Arrange
	var expectedBody []byte
	id := 1

	fields := map[string]interface{}{
		"sale_price": 5.99,
	}
	expectedProductRecord := models.ProductRecord{
		Id:            1,
		LastUpdate:    "2022-22-10",
		PurchasePrice: 4.99,
		SalePrice:     5.99,
		ProductId:     1,
	}

	expectedResponse := response.Response{Data: expectedProductRecord, StatusCode: http.StatusOK}

	s.mock.On("PartialModify", id, fields).Return(expectedProductRecord, nil)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchProductRecord(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *ProductRecordHandlerTestSuite) TestPatchProductRecord_BadRequest() {
	// Arrange
	var expectedBody []byte
	id := 1

	expectedError := errors.New("unexpected JSON format, check the request body")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusBadRequest}

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer([]byte("invalid json")))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchProductRecord(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *ProductRecordHandlerTestSuite) TestPatchProductRecord_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := -1
	expectedError := errors.New("invalid request")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusBadRequest}

	fields := map[string]interface{}{

		"last_update":    "2022-22-10",
		"purchase_price": 4.99,
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
	s.handler.PatchProductRecord(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *ProductRecordHandlerTestSuite) TestPatchProductRecord_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1

	fields := map[string]interface{}{

		"last_update":    "2022-22-10",
		"purchase_price": 4.99,
	}

	expectedError := errors.New("partial update failed")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("PartialModify", id, fields).Return(models.ProductRecord{}, expectedError)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchProductRecord(recorder, request)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// DeleteBuyer tests
func (s *ProductRecordHandlerTestSuite) TestDeleteProductRecord_Success() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedResponse := response.Response{Data: nil, StatusCode: http.StatusNoContent}

	s.mock.On("Remove", id).Return(nil)

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteProductRecord(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *ProductRecordHandlerTestSuite) TestDeleteProductRecord_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedError := errors.New("invalid request")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusBadRequest}

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", id)
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteProductRecord(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *ProductRecordHandlerTestSuite) TestDeleteProductRecord_InternalError() {
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
	s.handler.DeleteProductRecord(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Run the test suite
func TestProductRecordHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRecordHandlerTestSuite))
}
