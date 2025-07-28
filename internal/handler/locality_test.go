package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LocalityServiceMock struct {
	mock.Mock
}

func (m *LocalityServiceMock) RetrieveAll() ([]models.Locality, error) {
	args := m.Called()
	return args.Get(0).([]models.Locality), args.Error(1)
}

func (m *LocalityServiceMock) Retrieve(id int) (models.Locality, error) {
	args := m.Called(id)
	return args.Get(0).(models.Locality), args.Error(1)
}

func (m *LocalityServiceMock) RetrieveLocalityBySeller(id int) (models.LocalitySellerCount, error) {
	args := m.Called(id)
	return args.Get(0).(models.LocalitySellerCount), args.Error(1)
}

func (m *LocalityServiceMock) RetrieveCarriers() ([]models.LocalityCarrierCount, error) {
	args := m.Called()
	return args.Get(0).([]models.LocalityCarrierCount), args.Error(1)
}

func (m *LocalityServiceMock) Register(locality models.Locality) (models.Locality, error) {
	args := m.Called(locality)
	return args.Get(0).(models.Locality), args.Error(1)
}

func (m *LocalityServiceMock) RegisterWithNames(locality models.LocalityDoc) (models.LocalityDoc, error) {
	args := m.Called(locality)
	return args.Get(0).(models.LocalityDoc), args.Error(1)
}

func (m *LocalityServiceMock) Modify(locality models.Locality) (models.Locality, error) {
	args := m.Called(locality)
	return args.Get(0).(models.Locality), args.Error(1)
}

func (m *LocalityServiceMock) PartialModify(id int, fields map[string]any) (models.Locality, error) {
	args := m.Called(id, fields)
	return args.Get(0).(models.Locality), args.Error(1)
}

func (m *LocalityServiceMock) Remove(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *LocalityServiceMock) RetrieveAllLocalitiesBySeller() ([]models.LocalitySellerCount, error) {
	args := m.Called()
	return args.Get(0).([]models.LocalitySellerCount), args.Error(1)
}

func (m *LocalityServiceMock) RetrieveCarriersByLocality(id int) ([]models.LocalityCarrierCount, error) {
	args := m.Called(id)
	return args.Get(0).([]models.LocalityCarrierCount), args.Error(1)
}

type LocalityHandlerTestSuite struct {
	suite.Suite
	mock    *LocalityServiceMock
	handler *LocalityHandler
	path    string
}

func (s *LocalityHandlerTestSuite) SetupTest() {
	s.mock = new(LocalityServiceMock)
	s.handler = NewLocalityHandler(s.mock)
	s.path = "/api/v1/localities"
}

// Test GetLocalities - Success
func (s *LocalityHandlerTestSuite) TestGetLocalities_Success() {
	// Arrange
	expectedLocalities := []models.Locality{
		{Id: 1, Locality: "Buenos Aires", ProvinceId: 1},
		{Id: 2, Locality: "Córdoba", ProvinceId: 2},
	}
	expectedResponse := response.Response{Data: expectedLocalities, StatusCode: http.StatusOK}

	s.mock.On("RetrieveAll").Return(expectedLocalities, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetLocalities(recorder, request)

	var resp response.Response
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	s.NoError(err)

	expectedBody, _ := json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Test GetLocalities - Error
func (s *LocalityHandlerTestSuite) TestGetLocalities_Error() {
	// Arrange
	expectedError := errors.New("database connection error")

	s.mock.On("RetrieveAll").Return([]models.Locality{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetLocalities(recorder, request)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.Contains(recorder.Body.String(), "database connection error")
}

// Test GetLocality with ID - Success
func (s *LocalityHandlerTestSuite) TestGetLocality_WithID_Success() {
	// Arrange
	id := 1
	sellerCount := 5
	expectedLocality := models.LocalitySellerCount{
		LocalityDoc: models.LocalityDoc{
			Id:       id,
			Locality: "Buenos Aires",
			Province: "Buenos Aires",
			Country:  "Argentina",
		},
		SellerCount: &sellerCount,
	}
	expectedResult := []models.LocalitySellerCount{expectedLocality}
	expectedResponse := response.Response{Data: expectedResult, StatusCode: http.StatusOK}

	s.mock.On("RetrieveLocalityBySeller", id).Return(expectedLocality, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s?id=%d", s.path, id), nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetLocality(recorder, request)

	var resp response.Response
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	s.NoError(err)

	expectedBody, _ := json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Test GetLocality with ID - Not Found
func (s *LocalityHandlerTestSuite) TestGetLocality_WithID_NotFound() {
	// Arrange
	id := 999
	expectedError := repository.ErrEntityNotFound
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNotFound}

	s.mock.On("RetrieveLocalityBySeller", id).Return(models.LocalitySellerCount{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s?id=%d", s.path, id), nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetLocality(recorder, request)

	var resp response.Response
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	s.NoError(err)

	expectedBody, _ := json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Test GetLocality with Invalid ID
func (s *LocalityHandlerTestSuite) TestGetLocality_InvalidID() {
	// Arrange
	expectedResponse := response.Response{Message: repository.ErrIDInvalid.Error(), StatusCode: http.StatusBadRequest}

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s?id=invalid", s.path), nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetLocality(recorder, request)

	var resp response.Response
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	s.NoError(err)

	expectedBody, _ := json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Test GetLocality without ID - Get All Localities By Seller
func (s *LocalityHandlerTestSuite) TestGetLocality_WithoutID_Success() {
	// Arrange
	sellerCount1 := 5
	sellerCount2 := 3
	expectedLocalities := []models.LocalitySellerCount{
		{
			LocalityDoc: models.LocalityDoc{
				Id:       1,
				Locality: "Buenos Aires",
				Province: "Buenos Aires",
				Country:  "Argentina",
			},
			SellerCount: &sellerCount1,
		},
		{
			LocalityDoc: models.LocalityDoc{
				Id:       2,
				Locality: "Córdoba",
				Province: "Córdoba",
				Country:  "Argentina",
			},
			SellerCount: &sellerCount2,
		},
	}
	expectedResponse := response.Response{Data: expectedLocalities, StatusCode: http.StatusOK}

	s.mock.On("RetrieveAllLocalitiesBySeller").Return(expectedLocalities, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetLocality(recorder, request)

	var resp response.Response
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	s.NoError(err)

	expectedBody, _ := json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Test GetLocality without ID - Error
func (s *LocalityHandlerTestSuite) TestGetLocality_WithoutID_Error() {
	// Arrange
	expectedError := errors.New("database connection error")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusNotFound}

	s.mock.On("RetrieveAllLocalitiesBySeller").Return([]models.LocalitySellerCount{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetLocality(recorder, request)

	var resp response.Response
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	s.NoError(err)

	expectedBody, _ := json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Test PostLocality - Success
func (s *LocalityHandlerTestSuite) TestPostLocality_Success() {
	// Arrange
	localityId := 1
	localityName := "Buenos Aires"
	provinceName := "Buenos Aires"
	countryName := "Argentina"

	requestBody := fmt.Sprintf(`{
		"id": %d,
		"locality_name": "%s",
		"province_name": "%s",
		"country_name": "%s"
	}`, localityId, localityName, provinceName, countryName)

	expectedLocality := models.LocalityDoc{
		Id:       localityId,
		Locality: localityName,
		Province: provinceName,
		Country:  countryName,
	}
	expectedResponse := response.Response{Data: expectedLocality, StatusCode: http.StatusCreated}

	s.mock.On("RegisterWithNames", expectedLocality).Return(expectedLocality, nil)

	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer([]byte(requestBody)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostLocality(recorder, request)

	var resp response.Response
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	s.NoError(err)

	expectedBody, _ := json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusCreated, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Test PostLocality - Invalid JSON
func (s *LocalityHandlerTestSuite) TestPostLocality_InvalidJSON() {
	// Arrange
	requestBody := `{invalid json}`

	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer([]byte(requestBody)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostLocality(recorder, request)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
}

// Test PostLocality - Error (Province not found)
func (s *LocalityHandlerTestSuite) TestPostLocality_ProvinceNotFound() {
	// Arrange
	localityId := 1
	localityName := "Buenos Aires"
	provinceName := "NonExistent Province"
	countryName := "Argentina"

	requestBody := fmt.Sprintf(`{
		"id": %d,
		"locality_name": "%s",
		"province_name": "%s",
		"country_name": "%s"
	}`, localityId, localityName, provinceName, countryName)

	expectedLocality := models.LocalityDoc{
		Id:       localityId,
		Locality: localityName,
		Province: provinceName,
		Country:  countryName,
	}
	expectedError := repository.ErrProvinceNotFound
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusBadRequest}

	s.mock.On("RegisterWithNames", expectedLocality).Return(models.LocalityDoc{}, expectedError)

	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer([]byte(requestBody)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostLocality(recorder, request)

	var resp response.Response
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	s.NoError(err)

	expectedBody, _ := json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Test PostLocality - Entity Already Exists
func (s *LocalityHandlerTestSuite) TestPostLocality_EntityAlreadyExists() {
	// Arrange
	localityId := 1
	localityName := "Buenos Aires"
	provinceName := "Buenos Aires"
	countryName := "Argentina"

	requestBody := fmt.Sprintf(`{
		"id": %d,
		"locality_name": "%s",
		"province_name": "%s",
		"country_name": "%s"
	}`, localityId, localityName, provinceName, countryName)

	expectedLocality := models.LocalityDoc{
		Id:       localityId,
		Locality: localityName,
		Province: provinceName,
		Country:  countryName,
	}
	expectedError := repository.ErrEntityAlreadyExists
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusBadRequest}

	s.mock.On("RegisterWithNames", expectedLocality).Return(models.LocalityDoc{}, expectedError)

	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer([]byte(requestBody)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostLocality(recorder, request)

	var resp response.Response
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	s.NoError(err)

	expectedBody, _ := json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

// Test PostLocality - Missing Required Field
func (s *LocalityHandlerTestSuite) TestPostLocality_MissingRequiredField() {
	// Arrange - Request sin locality_name
	requestBody := `{
		"id": 1,
		"province_name": "Buenos Aires",
		"country_name": "Argentina"
	}`

	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer([]byte(requestBody)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostLocality(recorder, request)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
}

// Test cases for LocalityHandler

func (s *LocalityHandlerTestSuite) TestGetCarriers_Ok() {
	// Arrange
	var expectedBody []byte
	expectedCarriers := []models.LocalityCarrierCount{
		{LocalityID: 1, LocalityName: "L1", TotalCarriers: 3},
		{LocalityID: 2, LocalityName: "L2", TotalCarriers: 0},
	}

	expectedResponse := response.Response{Data: expectedCarriers}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("RetrieveCarriers").Return(expectedCarriers, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.Len(resp.Data, 2)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *LocalityHandlerTestSuite) TestGetCarriers_InternalError() {
	// Arrange
	var expectedBody []byte
	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("RetrieveCarriers").Return([]models.LocalityCarrierCount{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *LocalityHandlerTestSuite) TestGetCarriersById_Ok() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedCarriers := []models.LocalityCarrierCount{
		{LocalityID: 1, LocalityName: "L1", TotalCarriers: 3},
		{LocalityID: 2, LocalityName: "L2", TotalCarriers: 0},
	}

	expectedResponse := response.Response{Data: expectedCarriers}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("RetrieveCarriersByLocality", id).Return(expectedCarriers, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "?id=", id), nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.Len(resp.Data, 2)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *LocalityHandlerTestSuite) TestGetCarriersById_InvalidId() {
	// Arrange
	var expectedBody []byte
	id := "invalid"
	expectedCarriers := []models.LocalityCarrierCount{
		{LocalityID: 1, LocalityName: "L1", TotalCarriers: 3},
		{LocalityID: 2, LocalityName: "L2", TotalCarriers: 0},
	}

	expectedError := ErrInvalidId
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("RetrieveCarriersByLocality", id).Return(expectedCarriers, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "?id=", id), nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *LocalityHandlerTestSuite) TestGetCarriersById_InternalError() {
	// Arrange
	var expectedBody []byte
	id := 1
	expectedCarriers := []models.LocalityCarrierCount{
		{LocalityID: 1, LocalityName: "L1", TotalCarriers: 3},
		{LocalityID: 2, LocalityName: "L2", TotalCarriers: 0},
	}

	expectedError := errors.New("internal error")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}
	expectedBody, _ = json.Marshal(expectedResponse)

	s.mock.On("RetrieveCarriersByLocality", id).Return(expectedCarriers, expectedError)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(s.path, "?id=", id), nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetCarrier(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func TestLocalityHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(LocalityHandlerTestSuite))
}
