package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type CarrierServiceMock struct {
	mock.Mock
}

type CarrierHandlerTestSuite struct {
	suite.Suite
	mock    *CarrierServiceMock
	handler *CarrierDefault
	path    string
}

// Mock methods for CarrierService

func (s *CarrierServiceMock) RetrieveAll() ([]models.Carrier, error) {
	args := s.Called()
	return args.Get(0).([]models.Carrier), args.Error(1)
}

func (s *CarrierServiceMock) Retrieve(id int) (models.Carrier, error) {
	args := s.Called(id)
	return args.Get(0).(models.Carrier), args.Error(1)
}

func (s *CarrierServiceMock) Register(carrier models.Carrier) (models.Carrier, error) {
	args := s.Called(carrier)
	return args.Get(0).(models.Carrier), args.Error(1)
}

func (s *CarrierServiceMock) Modify(carrier models.Carrier) (models.Carrier, error) {
	args := s.Called(carrier)
	return args.Get(0).(models.Carrier), args.Error(1)
}

func (s *CarrierServiceMock) PartialModify(id int, fields map[string]any) (models.Carrier, error) {
	args := s.Called(id, fields)
	return args.Get(0).(models.Carrier), args.Error(1)
}

func (s *CarrierServiceMock) Remove(id int) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *CarrierHandlerTestSuite) SetupTest() {
	s.mock = new(CarrierServiceMock)
	s.handler = NewCarrierDefault(s.mock)
	s.path = "/api/v1/carriers"
}

// Test cases for CarrierHandler

func (s *CarrierHandlerTestSuite) TestGetCarriers_Ok() {
	// Arrange
	var expectedBody []byte
	expectedCarriers := []models.Carrier{
		{ID: 1, CId: "AAA-111", Address: "Boulevard", CompanyName: "Meli", Telephone: "123-456789", LocalityId: 1},
		{ID: 2, CId: "AAA-222", Address: "Plaza", CompanyName: "Meli", Telephone: "223-456789", LocalityId: 1},
	}

	expectedResponse := response.Response{Data: expectedCarriers}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("RetrieveAll").Return(expectedCarriers, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetCarriers(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.Len(resp.Data, 2)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestGetCarriers_InternalError() {
	// Arrange
	var expectedBody []byte
	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("RetrieveAll").Return([]models.Carrier{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetCarriers(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// // GetCarrier tests
func (s *CarrierHandlerTestSuite) TestGetCarrier_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedCarrier := models.Carrier{
		ID: 1, CId: "AAA-111", Address: "Boulevard", CompanyName: "Meli", Telephone: "123-456789", LocalityId: 1,
	}
	expectedResponse := response.Response{Data: expectedCarrier}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("Retrieve", id).Return(expectedCarrier, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "/", id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestGetCarrier_BadRequest_InvalidId() {
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
	s.handler.GetCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestGetCarrier_BadRequest_NegativeId() {
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
	s.handler.GetCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestGetCarrier_NotFound() {
	// Arrange
	var expectedBody []byte
	id := 999
	expectedError := errors.New("entity not found")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNotFound}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("Retrieve", id).Return(models.Carrier{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", s.path, id), nil)
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "999")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.GetCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PostCarrier tests
func (s *CarrierHandlerTestSuite) TestPostCarrier_Ok() {
	// Arrange
	var expectedBody []byte
	cId := "AAA-111"
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1

	requestBody := map[string]interface{}{
		"cid":          cId,
		"company_name": companyName,
		"address":      address,
		"telephone":    telephone,
		"locality_id":  localityId,
	}

	expectedCarrier := models.Carrier{
		ID: 1, CId: cId, CompanyName: companyName, Address: address, Telephone: telephone, LocalityId: localityId,
	}
	expectedResponse := response.Response{Data: expectedCarrier}
	expectedBody, _ = json.Marshal(expectedResponse)

	inputCarrier := models.Carrier{
		CId: cId, CompanyName: companyName, Address: address, Telephone: telephone, LocalityId: localityId,
	}

	s.mock.On("Register", inputCarrier).Return(expectedCarrier, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusCreated, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPostCarrier_BadRequest_MissingCode() {
	// Arrange
	var expectedBody []byte
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1
	requestBody := map[string]interface{}{
		"company_name": companyName,
		"address":      address,
		"telephone":    telephone,
		"locality_id":  localityId,
	}
	expectedResponse := response.Response{Message: "cid code must not be null", StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusUnprocessableEntity, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPostCarrier_CIdAlreadyExists() {
	// Arrange
	var expectedBody []byte
	cId := "AAA-111"
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1

	requestBody := map[string]interface{}{
		"cid":          cId,
		"company_name": companyName,
		"address":      address,
		"telephone":    telephone,
		"locality_id":  localityId,
	}

	expectedError := repository.ErrEntityAlreadyExists

	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusConflict}
	expectedBody, _ = json.Marshal(expectedResponse)

	inputCarrier := models.Carrier{
		CId: cId, CompanyName: companyName, Address: address, Telephone: telephone, LocalityId: localityId,
	}

	s.mock.On("Register", inputCarrier).Return(models.Carrier{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusConflict, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPostCarrier_LocalityDoesNotExist() {
	// Arrange
	var expectedBody []byte
	cId := "AAA-111"
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1

	requestBody := map[string]interface{}{
		"cid":          cId,
		"company_name": companyName,
		"address":      address,
		"telephone":    telephone,
		"locality_id":  localityId,
	}

	expectedError := repository.ErrLocalityNotFound

	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusUnprocessableEntity}
	expectedBody, _ = json.Marshal(expectedResponse)

	inputCarrier := models.Carrier{
		CId: cId, CompanyName: companyName, Address: address, Telephone: telephone, LocalityId: localityId,
	}

	s.mock.On("Register", inputCarrier).Return(models.Carrier{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusUnprocessableEntity, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPostCarrier_InternalError() {
	// Arrange
	var expectedBody []byte
	cId := "AAA-111"
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1

	requestBody := map[string]interface{}{
		"cid":          cId,
		"company_name": companyName,
		"address":      address,
		"telephone":    telephone,
		"locality_id":  localityId,
	}

	expectedError := errors.New("internal error")

	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	inputCarrier := models.Carrier{
		CId: cId, CompanyName: companyName, Address: address, Telephone: telephone, LocalityId: localityId,
	}

	s.mock.On("Register", inputCarrier).Return(models.Carrier{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PutCarrier tests
func (s *CarrierHandlerTestSuite) TestPutCarrier_Ok() {
	// Arrange
	var expectedBody []byte
	carrierId := 1
	cId := "AAA-111"
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1

	requestBody := map[string]interface{}{
		"cid":          cId,
		"company_name": companyName,
		"address":      address,
		"telephone":    telephone,
		"locality_id":  localityId,
	}

	expectedCarrier := models.Carrier{
		ID: carrierId, CId: cId, CompanyName: companyName, Address: address, Telephone: telephone, LocalityId: localityId,
	}
	expectedResponse := response.Response{Data: expectedCarrier}
	expectedBody, _ = json.Marshal(expectedResponse)

	inputCarrier := models.Carrier{
		ID: carrierId, CId: cId, CompanyName: companyName, Address: address, Telephone: telephone, LocalityId: localityId,
	}

	s.mock.On("Modify", inputCarrier).Return(expectedCarrier, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", s.path, carrierId), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPutCarrier_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	cId := "AAA-111"
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1

	requestBody := map[string]interface{}{
		"cid":          cId,
		"company_name": companyName,
		"address":      address,
		"telephone":    telephone,
		"locality_id":  localityId,
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
	s.handler.PutCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPutCarrier_BadRequest_NegativeId() {
	// Arrange
	var expectedBody []byte
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	id := -1
	cId := "AAA-111"
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1

	requestBody := map[string]interface{}{
		"cid":          cId,
		"company_name": companyName,
		"address":      address,
		"telephone":    telephone,
		"locality_id":  localityId,
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
	s.handler.PutCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPutCarrier_BadRequest() {
	// Arrange
	var expectedBody []byte
	carrierId := 1
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1

	requestBody := map[string]interface{}{
		"company_name": companyName,
		"address":      address,
		"telephone":    telephone,
		"locality_id":  localityId,
	}
	expectedResponse := response.Response{Message: "cid code must not be null", StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", s.path, carrierId), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPutCarrier_CIdAlreadyExists() {
	// Arrange
	var expectedBody []byte
	id := 1
	cId := "AAA-111"
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1

	requestBody := map[string]interface{}{
		"cid":          cId,
		"company_name": companyName,
		"address":      address,
		"telephone":    telephone,
		"locality_id":  localityId,
	}

	expectedError := repository.ErrEntityAlreadyExists
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusConflict}
	expectedBody, _ = json.Marshal(expectedResponse)

	inputCarrier := models.Carrier{
		ID: id, CId: cId, CompanyName: companyName, Address: address, Telephone: telephone, LocalityId: localityId,
	}

	s.mock.On("Modify", inputCarrier).Return(models.Carrier{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusConflict, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPutCarrier_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	cId := "AAA-111"
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1

	requestBody := map[string]interface{}{
		"cid":          cId,
		"company_name": companyName,
		"address":      address,
		"telephone":    telephone,
		"locality_id":  localityId,
	}

	expectedError := errors.New("internal error")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	inputCarrier := models.Carrier{
		ID: id, CId: cId, CompanyName: companyName, Address: address, Telephone: telephone, LocalityId: localityId,
	}

	s.mock.On("Modify", inputCarrier).Return(models.Carrier{}, expectedError)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PutCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// PatchCarrier tests
func (s *CarrierHandlerTestSuite) TestPatchCarrier_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	cId := "AAA-111"
	companyName := "Meli"
	address := "Boulevard"
	telephone := "123-456789"
	localityId := 1

	fields := map[string]interface{}{
		"company_name": companyName,
	}

	expectedCarrier := models.Carrier{
		ID: id, CId: cId, CompanyName: companyName, Address: address, Telephone: telephone, LocalityId: localityId,
	}
	expectedResponse := response.Response{Data: expectedCarrier}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("PartialModify", id, fields).Return(expectedCarrier, nil)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPatchCarrier_BadRequest_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedResponse := response.Response{Message: ErrInvalidId.Error(), StatusCode: http.StatusBadRequest}
	expectedBody, _ = json.Marshal(expectedResponse)

	fields := map[string]interface{}{
		"cid": "Partially Updated Company",
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
	s.handler.PatchCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPatchCarrier_BadRequest_NegativeId() {
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
	s.handler.PatchCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPatchCarrier_BadRequest() {
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
	s.handler.PatchCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPatchCarrier_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	fields := map[string]interface{}{
		"address": "Partially Updated Company",
	}

	expectedError := errors.New("internal error")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("PartialModify", id, fields).Return(models.Carrier{}, expectedError)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPatchCarrier_CarrierNotFound() {
	// Arrange
	var expectedBody []byte
	id := 1
	fields := map[string]interface{}{
		"cid": "CID-1",
	}

	expectedError := repository.ErrEntityNotFound
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNotFound}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("PartialModify", id, fields).Return(models.Carrier{}, expectedError)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestPatchCarrier_Conflict_CIdAlreadyExists() {
	// Arrange
	var expectedBody []byte
	id := 1
	fields := map[string]interface{}{
		"cid": "CID-1",
	}

	expectedError := repository.ErrEntityAlreadyExists
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusConflict}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("PartialModify", id, fields).Return(models.Carrier{}, expectedError)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprint(s.path, "/", id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Add chi context for URL parameters
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.PatchCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusConflict, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// DeleteCarrier tests
func (s *CarrierHandlerTestSuite) TestDeleteCarrier_NoContent() {
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
	s.handler.DeleteCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusNoContent, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestDeleteCarrier_InvalidId() {
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
	s.handler.DeleteCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *CarrierHandlerTestSuite) TestDeleteCarrier_InternalError() {
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
	s.handler.DeleteCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Run the test suite
func TestCarrierHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CarrierHandlerTestSuite))
}
