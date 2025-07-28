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

type SellerServiceMock struct {
	mock.Mock
}

type SellerHandlerTestSuite struct {
	suite.Suite
	mock    *SellerServiceMock
	handler *SellerHandler
	path    string
}

// Mock methods for SellerService

func (s *SellerServiceMock) RetrieveAll() ([]models.Seller, error) {
	args := s.Called()
	return args.Get(0).([]models.Seller), args.Error(1)
}

func (s *SellerServiceMock) Retrieve(id int) (models.Seller, error) {
	args := s.Called(id)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (s *SellerServiceMock) Register(seller models.Seller) (models.Seller, error) {
	args := s.Called(seller)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (s *SellerServiceMock) Modify(seller models.Seller) (models.Seller, error) {
	args := s.Called(seller)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (s *SellerServiceMock) PartialModify(id int, fields map[string]any) (models.Seller, error) {
	args := s.Called(id, fields)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (s *SellerServiceMock) Remove(id int) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *SellerHandlerTestSuite) SetupTest() {
	s.mock = new(SellerServiceMock)
	s.handler = NewSellerHandler(s.mock)
	s.path = "/api/v1/sellers"
}

// Test cases for SellerHandler

func (s *SellerHandlerTestSuite) TestGetSellers_Ok() {
	// Arrange
	var expectedBody []byte
	expectedSellers := []models.Seller{
		{Id: 1, Name: "Company A", Address: "123 Main St", Telephone: "555-0001", LocalityId: 1},
		{Id: 2, Name: "Company B", Address: "456 Oak Ave", Telephone: "555-0002", LocalityId: 2}}

	expectedResponse := response.Response{Data: expectedSellers}

	s.mock.On("RetrieveAll").Return(expectedSellers, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetSellers(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.Len(resp.Data, 2)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestGetSellers_InternalError() {
	// Arrange
	var expectedBody []byte
	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("RetrieveAll").Return([]models.Seller{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetSellers(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// GetSeller tests
func (s *SellerHandlerTestSuite) TestGetSeller_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedSeller := models.Seller{
		Id: id, Name: "Company A", Address: "123 Main St", Telephone: "555-0001", LocalityId: 1}
	expectedResponse := response.Response{Data: expectedSeller}

	s.mock.On("Retrieve", id).Return(expectedSeller, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestGetSeller_BadRequest_InvalidId() {
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
	s.handler.GetSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestGetSeller_BadRequest_NegativeId() {
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
	s.handler.GetSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestGetSeller_NotFound() {
	// Arrange
	var expectedBody []byte
	id := 999
	expectedError := errors.New("entity not found")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNotFound}

	s.mock.On("Retrieve", id).Return(models.Seller{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", s.path, id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "999")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PostSeller tests
func (s *SellerHandlerTestSuite) TestPostSeller_Ok() {
	// Arrange
	var expectedBody []byte
	name := "New Company"
	address := "789 New St"
	telephone := "555-0003"
	localityId := 3

	requestBody := map[string]interface{}{
		"name":        name,
		"address":     address,
		"telephone":   telephone,
		"locality_id": localityId,
	}

	expectedSeller := models.Seller{
		Id: 1, Name: name, Address: address, Telephone: telephone, LocalityId: localityId}
	expectedResponse := response.Response{Data: expectedSeller}

	inputSeller := models.Seller{
		Name: name, Address: address, Telephone: telephone, LocalityId: localityId}

	s.mock.On("Register", inputSeller).Return(expectedSeller, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusCreated, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestPostSeller_BadRequest_MissingName() {
	// Arrange
	var expectedBody []byte
	requestBody := map[string]interface{}{
		"address":     "789 New St",
		"telephone":   "555-0003",
		"locality_id": 3,
	}
	expectedResponse := response.Response{Message: "name must not be null", StatusCode: http.StatusBadRequest}

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestPostSeller_InternalError() {
	// Arrange
	var expectedBody []byte
	name := "New Company"
	address := "789 New St"
	telephone := "555-0003"
	localityId := 3

	requestBody := map[string]interface{}{
		"name":        name,
		"address":     address,
		"telephone":   telephone,
		"locality_id": localityId,
	}

	expectedError := errors.New("something went wrong")

	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	inputSeller := models.Seller{
		Name: name, Address: address, Telephone: telephone, LocalityId: localityId}

	s.mock.On("Register", inputSeller).Return(models.Seller{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PutSeller tests
func (s *SellerHandlerTestSuite) TestPutSeller_Ok() {
	// Arrange
	var expectedBody []byte
	sellerId := 1
	name := "Updated Company"
	address := "123 Updated St"
	telephone := "555-0004"
	localityId := 4

	requestBody := map[string]interface{}{
		"name":        name,
		"address":     address,
		"telephone":   telephone,
		"locality_id": localityId,
	}

	expectedSeller := models.Seller{
		Id: sellerId, Name: name, Address: address, Telephone: telephone, LocalityId: localityId}
	expectedResponse := response.Response{Data: expectedSeller}

	inputSeller := models.Seller{
		Id: sellerId, Name: name, Address: address, Telephone: telephone, LocalityId: localityId}

	s.mock.On("Modify", inputSeller).Return(expectedSeller, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", s.path, sellerId), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestPutSeller_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	requestBody := map[string]interface{}{
		"name":        "Updated Company",
		"address":     "123 Updated St",
		"telephone":   "555-0004",
		"locality_id": 4,
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
	s.handler.PutSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestPutSeller_BadRequest_NegativeId() {
	// Arrange
	var expectedBody []byte
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	id := -1
	requestBody := map[string]interface{}{
		"name":        "Updated Company",
		"address":     "123 Updated St",
		"telephone":   "555-0004",
		"locality_id": 4,
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
	s.handler.PutSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestPutSeller_BadRequest() {
	// Arrange
	var expectedBody []byte
	sellerId := 1
	requestBody := map[string]interface{}{
		"address":     "123 Updated St",
		"telephone":   "555-0004",
		"locality_id": 4,
	}
	expectedResponse := response.Response{Message: "name must not be null", StatusCode: http.StatusBadRequest}

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", s.path, sellerId), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestPutSeller_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	name := "Updated Company"
	address := "123 Updated St"
	telephone := "555-0004"
	localityId := 4

	requestBody := map[string]interface{}{
		"name":        name,
		"address":     address,
		"telephone":   telephone,
		"locality_id": localityId,
	}

	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	inputSeller := models.Seller{
		Id: id, Name: name, Address: address, Telephone: telephone, LocalityId: localityId}

	s.mock.On("Modify", inputSeller).Return(models.Seller{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PatchSeller tests
func (s *SellerHandlerTestSuite) TestPatchSeller_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	fields := map[string]interface{}{
		"name": "Partially Updated Company",
	}

	expectedSeller := models.Seller{
		Id: id, Name: "Partially Updated Company", Address: "123 Main St", Telephone: "555-0001", LocalityId: 1}
	expectedResponse := response.Response{Data: expectedSeller}

	s.mock.On("PartialModify", id, fields).Return(expectedSeller, nil)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestPatchSeller_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	fields := map[string]interface{}{
		"name": "Partially Updated Company",
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
	s.handler.PatchSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestPatchSeller_BadRequest_NegativeId() {
	// Arrange
	var expectedBody []byte
	id := -1
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	fields := map[string]interface{}{
		"name": "Partially Updated Company",
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
	s.handler.PatchSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestPatchSeller_BadRequest() {
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
	s.handler.PatchSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestPatchSeller_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	fields := map[string]interface{}{
		"name": "Partially Updated Company",
	}

	expectedError := errors.New("partial update failed")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("PartialModify", id, fields).Return(models.Seller{}, expectedError)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// DeleteSeller tests
func (s *SellerHandlerTestSuite) TestDeleteSeller_NoContent() {
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
	s.handler.DeleteSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusNoContent, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestDeleteSeller_InvalidId() {
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
	s.handler.DeleteSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SellerHandlerTestSuite) TestDeleteSeller_InternalError() {
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
	s.handler.DeleteSeller(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Run the test suite
func TestSellerHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(SellerHandlerTestSuite))
}
