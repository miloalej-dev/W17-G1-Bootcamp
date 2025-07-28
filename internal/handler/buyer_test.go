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

type BuyerServiceMock struct {
	mock.Mock
}

type BuyerHandlerTestSuite struct {
	suite.Suite
	mock    *BuyerServiceMock
	handler *BuyerHandler
	path    string
}

// Mock methods for BuyerService

func (s *BuyerServiceMock) RetrieveAll() ([]models.Buyer, error) {
	args := s.Called()
	return args.Get(0).([]models.Buyer), args.Error(1)
}

func (s *BuyerServiceMock) Retrieve(id int) (models.Buyer, error) {
	args := s.Called(id)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (s *BuyerServiceMock) Register(buyer models.Buyer) (models.Buyer, error) {
	args := s.Called(buyer)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (s *BuyerServiceMock) Modify(buyer models.Buyer) (models.Buyer, error) {
	args := s.Called(buyer)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (s *BuyerServiceMock) PartialModify(id int, fields map[string]any) (models.Buyer, error) {
	args := s.Called(id, fields)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (s *BuyerServiceMock) Remove(id int) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *BuyerServiceMock) RetrieveByPurchaseOrderReport(id int) ([]models.BuyerReport, error) {
	args := s.Called(id)
	return args.Get(0).([]models.BuyerReport), args.Error(1)
}

func (s *BuyerHandlerTestSuite) SetupTest() {
	s.mock = new(BuyerServiceMock)
	s.handler = NewBuyerHandler(s.mock)
	s.path = "/api/v1/buyers"
}

// Test cases for BuyerHandler

func (s *BuyerHandlerTestSuite) TestGetBuyers_Success() {
	// Arrange

	expectedBuyers := []models.Buyer{
		{
			Id:           1,
			CardNumberId: "189-58-5819",
			FirstName:    "Donnamarie",
			LastName:     "Sharpless",
		},
		{
			Id:           2,
			CardNumberId: "174-53-5631",
			FirstName:    "john",
			LastName:     "Smith",
		},
	}

	expectedData := response.Response{Data: expectedBuyers, StatusCode: http.StatusOK}

	s.mock.On("RetrieveAll").Return(expectedBuyers, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetBuyers(recorder, request)

	expectedResponse, _ := json.Marshal(expectedData)

	// Assert

	s.Equal(expectedData.StatusCode, recorder.Code)
	s.JSONEq(string(expectedResponse), recorder.Body.String())
}

func (s *BuyerHandlerTestSuite) TestGetBuyers_InternalError() {
	// Arrange
	var expectedBody []byte
	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("RetrieveAll").Return([]models.Buyer{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetBuyers(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// GetBuyer tests
func (s *BuyerHandlerTestSuite) TestGetBuyer_Success() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedBuyer := models.Buyer{

		Id:           1,
		CardNumberId: "189-58-5819",
		FirstName:    "Donnamarie",
		LastName:     "Sharpless",
	}
	expectedResponse := response.Response{Data: expectedBuyer, StatusCode: http.StatusOK}

	s.mock.On("Retrieve", id).Return(expectedBuyer, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetBuyer(recorder, request)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *BuyerHandlerTestSuite) TestGetBuyer_BadRequest() {
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
	s.handler.GetBuyer(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *BuyerHandlerTestSuite) TestGetBuyer_NotFound() {
	// Arrange
	var expectedBody []byte
	id := 999

	expectedError := errors.New("entity not found")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNotFound}

	s.mock.On("Retrieve", id).Return(models.Buyer{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", s.path, id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "999")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetBuyer(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PostBuyer tests
func (s *BuyerHandlerTestSuite) TestPostBuyer_Success() {
	// Arrange
	var expectedBody []byte

	requestBody := map[string]interface{}{
		"card_number_id": "189-58-5819",
		"first_name":     "Donnamarie",
		"last_name":      "Sharpless",
	}

	expectedBuyer := models.Buyer{
		Id: 1, CardNumberId: "189-58-5819", FirstName: "Donnamarie", LastName: "Sharpless"}

	expectedResponse := response.Response{Data: expectedBuyer, StatusCode: http.StatusCreated}

	inputBuyer := models.Buyer{
		CardNumberId: "189-58-5819", FirstName: "Donnamarie", LastName: "Sharpless"}

	s.mock.On("Register", inputBuyer).Return(expectedBuyer, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostBuyer(recorder, request)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusCreated, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *BuyerHandlerTestSuite) TestPostBuyer_BadRequest() {
	// Arrange
	var expectedBody []byte

	requestBody := map[string]interface{}{
		"first_name": "Donnamarie",
		"last_name":  "Sharpless",
	}

	expectedResponse := response.Response{Message: "card number Id must be not null", StatusCode: http.StatusBadRequest}

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostBuyer(recorder, request)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *BuyerHandlerTestSuite) TestPostBuyer_InternalError() {
	// Arrange

	var expectedBody []byte

	requestBody := map[string]interface{}{
		"card_number_id": "189-58-5819",
		"first_name":     "Donnamarie",
		"last_name":      "Sharpless",
	}

	expectedError := errors.New("something went wrong")

	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	inputBuyer := models.Buyer{
		CardNumberId: "189-58-5819", FirstName: "Donnamarie", LastName: "Sharpless"}

	s.mock.On("Register", inputBuyer).Return(models.Buyer{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostBuyer(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PatchBuyer tests
func (s *BuyerHandlerTestSuite) TestPatchBuyer_Success() {
	// Arrange
	var expectedBody []byte
	id := 1
	fields := map[string]interface{}{
		"first_name": "Donnamarie",
		"last_name":  "Sharpless",
	}

	expectedBuyer := models.Buyer{
		Id:           id,
		CardNumberId: "189-58-5819",
		FirstName:    "Don",
		LastName:     "Sharp",
	}

	expectedResponse := response.Response{Data: expectedBuyer}

	s.mock.On("PartialModify", id, fields).Return(expectedBuyer, nil)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchBuyer(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *BuyerHandlerTestSuite) TestPatchBuyer_BadRequest() {
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
	s.handler.PatchBuyer(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *BuyerHandlerTestSuite) TestPatchBuyer_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := -1
	expectedError := errors.New("invalid request")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusBadRequest}

	fields := map[string]interface{}{
		"first_name": "Donnamarie",
		"last_name":  "Sharpless",
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
	s.handler.PatchBuyer(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *BuyerHandlerTestSuite) TestPatchBuyer_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1

	fields := map[string]interface{}{
		"first_name": "Donnamarie",
		"last_name":  "Sharpless",
	}

	expectedError := errors.New("partial update failed")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("PartialModify", id, fields).Return(models.Buyer{}, expectedError)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchBuyer(recorder, request)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// DeleteBuyer tests
func (s *BuyerHandlerTestSuite) TestDeleteBuyer_Success() {
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
	s.handler.DeleteBuyer(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *BuyerHandlerTestSuite) TestDeleteBuyer_InvalidId() {
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
	s.handler.DeleteBuyer(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *BuyerHandlerTestSuite) TestDeleteBuyer_InternalError() {
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
	s.handler.DeleteBuyer(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(expectedResponse.StatusCode, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Run the test suite
func TestBuyerHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BuyerHandlerTestSuite))
}
