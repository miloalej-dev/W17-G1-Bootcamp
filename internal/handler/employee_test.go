package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EmployeeServiceMock struct {
	mock.Mock
}

type EmployeeHandlerTestSuite struct {
	suite.Suite
	mock    *EmployeeServiceMock
	handler *EmployeeHandler
	path    string
}

// Mock methods for EmployeeService

func (m *EmployeeServiceMock) RetrieveAll() ([]models.Employee, error) {
	args := m.Called()
	return args.Get(0).([]models.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) Retrieve(id int) (models.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) Register(employee models.Employee) (models.Employee, error) {
	args := m.Called(employee)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) Modify(employee models.Employee) (models.Employee, error) {
	args := m.Called(employee)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) PartialModify(id int, fields map[string]any) (models.Employee, error) {
	args := m.Called(id, fields)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) Remove(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *EmployeeServiceMock) RetrieveInboundOrdersReport() ([]models.EmployeeInboundOrdersReport, error) {
	args := m.Called()
	return args.Get(0).([]models.EmployeeInboundOrdersReport), args.Error(1)
}

func (m *EmployeeServiceMock) RetrieveInboundOrdersReportById(id int) (models.EmployeeInboundOrdersReport, error) {
	args := m.Called(id)
	return args.Get(0).(models.EmployeeInboundOrdersReport), args.Error(1)
}

func (s *EmployeeHandlerTestSuite) SetupTest() {
	s.mock = new(EmployeeServiceMock)
	s.handler = NewEmployeeHandler(s.mock)
	s.path = "/api/v1/employees"
}

// Test cases for EmployeeHandler

// GetEmployees tests
func (s *EmployeeHandlerTestSuite) TestGetEmployees_Ok() {
	// Arrange
	var expectedBody []byte
	expectedEmployees := []models.Employee{
		{Id: 1, CardNumberId: "123456789", FirstName: "John", LastName: "Doe", WarehouseId: 1},
		{Id: 2, CardNumberId: "987654321", FirstName: "Jane", LastName: "Smith", WarehouseId: 2},
	}

	expectedResponse := response.Response{Data: expectedEmployees}

	s.mock.On("RetrieveAll").Return(expectedEmployees, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetEmployees(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.Len(resp.Data, 2)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestGetEmployees_InternalError() {
	// Arrange
	var expectedBody []byte
	expectedError := errors.New("database connection error")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNoContent}

	s.mock.On("RetrieveAll").Return([]models.Employee{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetEmployees(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusNoContent, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// GetEmployee tests
func (s *EmployeeHandlerTestSuite) TestGetEmployee_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedEmployee := models.Employee{
		Id: id, CardNumberId: "123456789", FirstName: "John", LastName: "Doe", WarehouseId: 1}
	expectedResponse := response.Response{Data: expectedEmployee}

	s.mock.On("Retrieve", id).Return(expectedEmployee, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestGetEmployee_BadRequest_InvalidId() {
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
	s.handler.GetEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestGetEmployee_BadRequest_NegativeId() {
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
	s.handler.GetEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestGetEmployee_NotFound() {
	// Arrange
	var expectedBody []byte
	id := 999
	expectedError := errors.New("entity not found")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNotFound}

	s.mock.On("Retrieve", id).Return(models.Employee{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PostEmployee tests
func (s *EmployeeHandlerTestSuite) TestPostEmployee_Created() {
	// Arrange
	var expectedBody []byte
	cardNumberId := "123456789"
	firstName := "John"
	lastName := "Doe"
	warehouseId := 1

	requestBody := map[string]interface{}{
		"card_number_id": cardNumberId,
		"first_name":     firstName,
		"last_name":      lastName,
		"warehouse_id":   warehouseId,
	}

	expectedEmployee := models.Employee{
		Id: 1, CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: warehouseId}
	expectedResponse := response.Response{Data: expectedEmployee}

	inputEmployee := models.Employee{
		CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: warehouseId}

	s.mock.On("Register", inputEmployee).Return(expectedEmployee, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.CreateEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusCreated, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestPostEmployee_BadRequest_MissingCardNumberId() {
	// Arrange
	var expectedBody []byte
	requestBody := map[string]interface{}{
		"first_name":   "John",
		"last_name":    "Doe",
		"warehouse_id": 1,
	}
	expectedResponse := response.Response{Message: "CardNumberId must not be null", StatusCode: http.StatusUnprocessableEntity}

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.CreateEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusUnprocessableEntity, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestPostEmployee_BadRequest_InvalidJSON() {
	// Arrange
	var expectedBody []byte
	invalidJSON := `{"invalid": json}`
	expectedResponse := response.Response{Message: "invalid character 'j' looking for beginning of value", StatusCode: http.StatusUnprocessableEntity}

	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBufferString(invalidJSON))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.CreateEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusUnprocessableEntity, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestPostEmployee_InternalError() {
	// Arrange
	var expectedBody []byte
	cardNumberId := "123456789"
	firstName := "John"
	lastName := "Doe"
	warehouseId := 1

	requestBody := map[string]interface{}{
		"card_number_id": cardNumberId,
		"first_name":     firstName,
		"last_name":      lastName,
		"warehouse_id":   warehouseId,
	}

	inputEmployee := models.Employee{
		CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: warehouseId}

	expectedError := errors.New("database error")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("Register", inputEmployee).Return(models.Employee{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.CreateEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PutEmployee tests
func (s *EmployeeHandlerTestSuite) TestPutEmployee_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	cardNumberId := "987654321"
	firstName := "Jane"
	lastName := "Smith"
	warehouseId := 2

	requestBody := map[string]interface{}{
		"card_number_id": cardNumberId,
		"first_name":     firstName,
		"last_name":      lastName,
		"warehouse_id":   warehouseId,
	}

	expectedEmployee := models.Employee{
		Id: id, CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: warehouseId}
	expectedResponse := response.Response{Data: expectedEmployee}

	inputEmployee := models.Employee{
		Id: id, CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: warehouseId}

	s.mock.On("Modify", inputEmployee).Return(expectedEmployee, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestPutEmployee_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	requestBody := map[string]interface{}{
		"card_number_id": "987654321",
		"first_name":     "Jane",
		"last_name":      "Smith",
		"warehouse_id":   2,
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
	s.handler.PutEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestPutEmployee_BadRequest_InvalidJSON() {
	// Arrange
	var expectedBody []byte
	id := 1
	invalidJSON := `{"invalid": json}`
	expectedResponse := response.Response{Message: "invalid character 'j' looking for beginning of value", StatusCode: http.StatusBadRequest}

	request := httptest.NewRequest(http.MethodPut, fmt.Sprint(s.path, "/", id), bytes.NewBufferString(invalidJSON))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestPutEmployee_NotFound() {
	// Arrange
	var expectedBody []byte
	id := 999
	cardNumberId := "987654321"
	firstName := "Jane"
	lastName := "Smith"
	warehouseId := 2

	requestBody := map[string]interface{}{
		"card_number_id": cardNumberId,
		"first_name":     firstName,
		"last_name":      lastName,
		"warehouse_id":   warehouseId,
	}

	inputEmployee := models.Employee{
		Id: id, CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: warehouseId}

	expectedError := errors.New("entity not found")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNotFound}

	s.mock.On("Modify", inputEmployee).Return(models.Employee{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PatchEmployee tests
func (s *EmployeeHandlerTestSuite) TestPatchEmployee_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	fields := map[string]interface{}{
		"first_name": "UpdatedName",
	}

	expectedEmployee := models.Employee{
		Id: id, CardNumberId: "123456789", FirstName: "UpdatedName", LastName: "Doe", WarehouseId: 1}
	expectedResponse := response.Response{Data: expectedEmployee}

	s.mock.On("PartialModify", id, fields).Return(expectedEmployee, nil)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestPatchEmployee_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	fields := map[string]interface{}{
		"first_name": "UpdatedName",
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
	s.handler.PatchEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestPatchEmployee_BadRequest_NegativeId() {
	// Arrange
	var expectedBody []byte
	id := -1
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	fields := map[string]interface{}{
		"first_name": "UpdatedName",
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
	s.handler.PatchEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestPatchEmployee_BadRequest_InvalidJSON() {
	// Arrange
	var expectedBody []byte
	id := 1
	invalidJSON := `{invalid json}`
	expectedResponse := response.Response{Message: "invalid character 'i' looking for beginning of object key string", StatusCode: http.StatusBadRequest}

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBufferString(invalidJSON))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestPatchEmployee_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	fields := map[string]interface{}{
		"first_name": "UpdatedName",
	}

	expectedError := errors.New("service error")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("PartialModify", id, fields).Return(models.Employee{}, expectedError)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// DeleteEmployee tests
func (s *EmployeeHandlerTestSuite) TestDeleteEmployee_NoContent() {
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
	s.handler.DeleteEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusNoContent, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestDeleteEmployee_BadRequest_InvalidId() {
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
	s.handler.DeleteEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestDeleteEmployee_BadRequest_NegativeId() {
	// Arrange
	var expectedBody []byte
	id := -1
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *EmployeeHandlerTestSuite) TestDeleteEmployee_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 999
	expectedError := errors.New("employee not found")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	s.mock.On("Remove", id).Return(expectedError)

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteEmployee(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func TestEmployeeHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(EmployeeHandlerTestSuite))
}
