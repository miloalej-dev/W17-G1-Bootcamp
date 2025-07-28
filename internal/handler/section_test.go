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

type SectionServiceMock struct {
	mock.Mock
}

type SectionHandlerTestSuite struct {
	suite.Suite
	mock    *SectionServiceMock
	handler *SectionHandler
	path    string
}

// Mock methods for SectionService

func (s *SectionServiceMock) RetrieveAll() ([]models.Section, error) {
	args := s.Called()
	return args.Get(0).([]models.Section), args.Error(1)
}

func (s *SectionServiceMock) Retrieve(id int) (models.Section, error) {
	args := s.Called(id)
	return args.Get(0).(models.Section), args.Error(1)
}

func (s *SectionServiceMock) Register(section models.Section) (models.Section, error) {
	args := s.Called(section)
	return args.Get(0).(models.Section), args.Error(1)
}

func (s *SectionServiceMock) Modify(section models.Section) (models.Section, error) {
	args := s.Called(section)
	return args.Get(0).(models.Section), args.Error(1)
}

func (s *SectionServiceMock) PartialModify(id int, fields map[string]any) (models.Section, error) {
	args := s.Called(id, fields)
	return args.Get(0).(models.Section), args.Error(1)
}

func (s *SectionServiceMock) Remove(id int) error {
	args := s.Called(id)
	return args.Error(0)
}
func (s *SectionServiceMock) RetrieveSectionReport(sectionId *int) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SectionHandlerTestSuite) SetupTest() {
	s.mock = new(SectionServiceMock)
	s.handler = NewSectionDefault(s.mock)
	s.path = "/api/v1/sellers"
}

func (s *SectionHandlerTestSuite) TestGetSections_Ok() {
	var expectedBody []byte

	var expectedSections = []models.Section{
		{
			Id:                 1,
			SectionNumber:      "A01",
			CurrentTemperature: 4.5,
			MinimumTemperature: 2.0,
			CurrentCapacity:    50,
			MinimumCapacity:    10,
			MaximumCapacity:    100,
			WarehouseId:        1,
			ProductTypeId:      101,
		},
		{
			Id:                 2,
			SectionNumber:      "B12",
			CurrentTemperature: 6.0,
			MinimumTemperature: 3.5,
			CurrentCapacity:    75,
			MinimumCapacity:    20,
			MaximumCapacity:    150,
			WarehouseId:        2,
			ProductTypeId:      202,
		},
	}

	expectedResponse := response.Response{Data: expectedSections}
	s.mock.On("RetrieveAll").Return(expectedSections, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetSections(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.Len(resp.Data, 2)
	s.JSONEq(string(expectedBody), recorder.Body.String())

}

func (s *SectionHandlerTestSuite) TestGetSections_InternalError() {
	// Arrange
	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{
		Message:    expectedError.Error(),
		StatusCode: http.StatusInternalServerError,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("RetrieveAll").Return([]models.Section{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetSections(recorder, request)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SectionHandlerTestSuite) TestGetSections_NotFound() {
	// Arrange
	expectedResponse := response.Response{
		Message:    "not found",
		StatusCode: http.StatusNotFound,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("RetrieveAll").Return([]models.Section{}, nil)

	request := httptest.NewRequest(http.MethodGet, s.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	s.handler.GetSections(recorder, request)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SectionHandlerTestSuite) TestGetSection_Ok() {
	// Arrange
	section := models.Section{
		Id:                 1,
		SectionNumber:      "A01",
		CurrentTemperature: 4.5,
		MinimumTemperature: 2.0,
		CurrentCapacity:    50,
		MinimumCapacity:    10,
		MaximumCapacity:    100,
		WarehouseId:        1,
		ProductTypeId:      101,
	}
	expectedResponse := response.Response{Data: section}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("Retrieve", section.Id).Return(section, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", s.path, section.Id), nil)
	recorder := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", strconv.Itoa(section.Id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	// Act
	s.handler.GetSection(recorder, request)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SectionHandlerTestSuite) TestGetSection_InvalidID() {
	// Arrange
	expectedResponse := response.Response{
		Message:    "strconv.Atoi: parsing \"invalid\": invalid syntax",
		StatusCode: http.StatusBadRequest,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	request := httptest.NewRequest(http.MethodGet, s.path+"/invalid", nil)
	recorder := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalid")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	// Act
	s.handler.GetSection(recorder, request)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SectionHandlerTestSuite) TestGetSection_NotFound() {
	// Arrange
	id := 999
	expectedErr := errors.New("section not found")
	expectedResponse := response.Response{
		Message:    expectedErr.Error(),
		StatusCode: http.StatusNotFound,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("Retrieve", id).Return(models.Section{}, expectedErr)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", s.path, id), nil)
	recorder := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	// Act
	s.handler.GetSection(recorder, request)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SectionHandlerTestSuite) TestPostSection_Ok() {
	// Arrange
	requestBody := map[string]interface{}{
		"section_number":      "A01",
		"current_temperature": 4.5,
		"minimum_temperature": 2.0,
		"current_capacity":    50,
		"minimum_capacity":    10,
		"maximum_capacity":    100,
		"warehouse_id":        1,
		"product_type_id":     101,
	}

	inputSection := models.Section{
		SectionNumber:      "A01",
		CurrentTemperature: 4.5,
		MinimumTemperature: 2.0,
		CurrentCapacity:    50,
		MinimumCapacity:    10,
		MaximumCapacity:    100,
		WarehouseId:        1,
		ProductTypeId:      101,
	}

	expectedSection := inputSection
	expectedSection.Id = 1

	expectedResponse := response.Response{Data: expectedSection}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("Register", inputSection).Return(expectedSection, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostSection(recorder, request)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SectionHandlerTestSuite) TestPostSection_Conflict_AlreadyExists() {
	// Arrange
	requestBody := map[string]interface{}{
		"section_number":      "A01",
		"current_temperature": 4.5,
		"minimum_temperature": 2.0,
		"current_capacity":    50,
		"minimum_capacity":    10,
		"maximum_capacity":    100,
		"warehouse_id":        1,
		"product_type_id":     101,
	}

	inputSection := models.Section{
		SectionNumber:      "A01",
		CurrentTemperature: 4.5,
		MinimumTemperature: 2.0,
		CurrentCapacity:    50,
		MinimumCapacity:    10,
		MaximumCapacity:    100,
		WarehouseId:        1,
		ProductTypeId:      101,
	}

	expectedErr := repository.ErrEntityAlreadyExists
	expectedResponse := response.Response{
		Message:    expectedErr.Error(),
		StatusCode: http.StatusConflict,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("Register", inputSection).Return(models.Section{}, expectedErr)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostSection(recorder, request)

	// Assert
	s.Equal(http.StatusConflict, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SectionHandlerTestSuite) TestPostSection_BindError() {
	// Arrange
	invalidBody := []byte(`{ "section_number": 123 }`) // string esperado, número enviado

	expectedResponse := response.Response{
		Message:    "json: cannot unmarshal number into Go struct field SectionRequest.section_number of type string",
		StatusCode: http.StatusBadRequest,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	request := httptest.NewRequest(http.MethodPost, s.path, bytes.NewBuffer(invalidBody))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	s.handler.PostSection(recorder, request)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SectionHandlerTestSuite) TestPatchSection_Ok() {
	// Arrange
	id := 1
	fields := map[string]interface{}{
		"current_capacity": float64(70),
	}

	expectedSection := models.Section{
		Id:                 id,
		SectionNumber:      "A01",
		CurrentTemperature: 4.5,
		MinimumTemperature: 2.0,
		CurrentCapacity:    70,
		MinimumCapacity:    10,
		MaximumCapacity:    100,
		WarehouseId:        1,
		ProductTypeId:      101,
	}
	expectedResponse := response.Response{Data: expectedSection}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("PartialModify", id, fields).Return(expectedSection, nil)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")

	// Add chi context
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	recorder := httptest.NewRecorder()

	// Act
	s.handler.PatchSection(recorder, request)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(string(expectedBody), recorder.Body.String())
}

func (s *SectionHandlerTestSuite) TestPatchSection_NotFound() {
	// Arrange
	id := 2
	fields := map[string]interface{}{
		"current_capacity": float64(30),
	}
	expectedErr := repository.ErrEntityNotFound

	s.mock.On("PartialModify", id, fields).Return(models.Section{}, expectedErr)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	recorder := httptest.NewRecorder()

	// Act
	s.handler.PatchSection(recorder, request)

	// Assert
	s.Equal(http.StatusNotFound, recorder.Code)
	s.Contains(recorder.Body.String(), expectedErr.Error())
}

func (s *SectionHandlerTestSuite) TestPatchSection_InternalServerError() {
	// Arrange
	id := 3
	fields := map[string]interface{}{
		"current_capacity": float64(80),
	}
	expectedErr := errors.New("db connection lost")

	s.mock.On("PartialModify", id, fields).Return(models.Section{}, expectedErr)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	recorder := httptest.NewRecorder()

	// Act
	s.handler.PatchSection(recorder, request)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)
	s.Contains(recorder.Body.String(), expectedErr.Error())
}

func (s *SectionHandlerTestSuite) TestPatchSection_BadRequest_InvalidID() {
	// Arrange
	requestBody := []byte(`{"current_capacity": 70}`)

	request := httptest.NewRequest(http.MethodPatch, s.path+"/bad_id", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "bad_id")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	recorder := httptest.NewRecorder()

	// Act
	s.handler.PatchSection(recorder, request)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.Contains(recorder.Body.String(), "invalid syntax")
}

func (s *SectionHandlerTestSuite) TestPatchSection_BadRequest_InvalidBody() {
	// Arrange
	id := 4
	invalidBody := []byte(`{"current_capacity":}`) // valor faltante

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", s.path, id), bytes.NewBuffer(invalidBody))
	request.Header.Set("Content-Type", "application/json")

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	recorder := httptest.NewRecorder()

	// Act
	s.handler.PatchSection(recorder, request)

	// Assert
	s.Equal(http.StatusBadRequest, recorder.Code)
	s.Contains(recorder.Body.String(), "invalid character")
}

func (s *SectionHandlerTestSuite) TestDeleteSection_NoContent() {
	// Arrange
	id := 1
	expectedResponse := response.Response{Data: nil}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("Remove", id).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/sections/%d", id), nil)
	rec := httptest.NewRecorder()

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteSection(rec, req)

	// Assert
	s.Equal(http.StatusNoContent, rec.Code)
	s.JSONEq(string(expectedBody), rec.Body.String())
}

func (s *SectionHandlerTestSuite) TestDeleteSection_NotFound() {
	// Arrange
	id := 99
	expectedErr := errors.New("section not found")
	expectedResponse := response.Response{Data: expectedErr.Error()}
	expectedBody, _ := json.Marshal(expectedResponse)

	s.mock.On("Remove", id).Return(expectedErr)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/sections/%d", id), nil)
	rec := httptest.NewRecorder()

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteSection(rec, req)

	// Assert
	s.Equal(http.StatusNotFound, rec.Code)
	s.JSONEq(string(expectedBody), rec.Body.String())
}

func (s *SectionHandlerTestSuite) TestDeleteSection_BadRequest() {
	// Arrange
	expectedErr := "strconv.Atoi: parsing \"abc\": invalid syntax"
	expectedResponse := response.Response{Data: expectedErr}
	expectedBody, _ := json.Marshal(expectedResponse)

	req := httptest.NewRequest(http.MethodDelete, "/sections/abc", nil)
	rec := httptest.NewRecorder()

	// Simular contexto con ID inválido
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	// Act
	s.handler.DeleteSection(rec, req)

	// Assert
	s.Equal(http.StatusBadRequest, rec.Code)
	s.JSONEq(string(expectedBody), rec.Body.String())
}

// Run the test suite
func TestSectionHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(SectionHandlerTestSuite))
}
