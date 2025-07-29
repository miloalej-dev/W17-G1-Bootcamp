package handler

import (
	"bytes"
	"encoding/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ProductBatchServiceMock struct {
	mock.Mock
}

type ProductBatchHandlerTestSuite struct {
	suite.Suite
	mock    *ProductBatchServiceMock
	handler *ProductBatchDefault
	path    string
}

// Mock methods for SellerService
// Mock methods for SectionService

func (p *ProductBatchServiceMock) RetrieveAll() ([]models.ProductBatch, error) {
	args := p.Called()
	return args.Get(0).([]models.ProductBatch), args.Error(1)
}

func (p *ProductBatchServiceMock) Retrieve(id int) (models.ProductBatch, error) {
	args := p.Called(id)
	return args.Get(0).(models.ProductBatch), args.Error(1)
}

func (p *ProductBatchServiceMock) Modify(productBatch models.ProductBatch) (models.ProductBatch, error) {
	args := p.Called(productBatch)
	return args.Get(0).(models.ProductBatch), args.Error(1)
}

func (p *ProductBatchServiceMock) PartialModify(id int, fields map[string]any) (models.ProductBatch, error) {
	args := p.Called(id, fields)
	return args.Get(0).(models.ProductBatch), args.Error(1)
}
func (p *ProductBatchServiceMock) Remove(id int) error {
	args := p.Called(id)
	return args.Error(0)
}

func (p *ProductBatchServiceMock) Register(productBatch models.ProductBatch) (models.ProductBatch, error) {
	args := p.Called(productBatch)
	return args.Get(0).(models.ProductBatch), args.Error(1)
}

func (p *ProductBatchHandlerTestSuite) SetupTest() {
	p.mock = new(ProductBatchServiceMock)
	p.handler = NewProductBatchDefault(p.mock)
	p.path = "/api/v1/productBatches"
}

// PostSeller tests
func (p *ProductBatchHandlerTestSuite) TestPostSeller_Ok() {
	requestBody := map[string]interface{}{
		"batch_number":        40,
		"current_quantity":    200,
		"current_temperature": 20,
		"due_date":            "2022-04-04",
		"initial_quantity":    10,
		"manufacturing_date":  "2020-04-04",
		"manufacturing_hour":  10,
		"minimum_temperature": 5,
		"product_id":          1,
		"section_id":          1,
	}
	inputProductBatch := models.ProductBatch{
		BatchNumber:        40,
		CurrentQuantity:    200,
		CurrentTemperature: 20,
		DueDate:            "2022-04-04",
		InitialQuantity:    10,
		ManufacturingDate:  "2020-04-04",
		ManufacturingHour:  10,
		MinimumTemperature: 5,
		SectionId:          1,
		ProductId:          1,
	}

	expectedProductBatch := inputProductBatch
	expectedProductBatch.Id = 1

	expectedResponse := response.Response{Data: expectedProductBatch}
	expectedBody, _ := json.Marshal(expectedResponse)

	p.mock.On("Register", inputProductBatch).Return(expectedProductBatch, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, p.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	p.handler.PostProductBatch(recorder, request)

	// Assert
	p.Equal(http.StatusCreated, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())

}

// PostSeller tests
func (p *ProductBatchHandlerTestSuite) TestPostSeller_errorBinder() {
	requestBody := map[string]interface{}{
		"batch_number":        40,
		"current_temperature": 20,
		"due_date":            "2022-04-04",
		"initial_quantity":    10,
		"manufacturing_date":  "2020-04-04",
		"manufacturing_hour":  10,
		"minimum_temperature": 5,
		"product_id":          1,
		"section_id":          1,
	}

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, p.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	p.handler.PostProductBatch(recorder, request)

	// Assert
	p.Equal(http.StatusUnprocessableEntity, recorder.Code)
}

// PostSeller tests
func (p *ProductBatchHandlerTestSuite) TestPostSeller_errorService() {
	requestBody := map[string]interface{}{
		"batch_number":        40,
		"current_quantity":    200,
		"current_temperature": 20,
		"due_date":            "2022-04-04",
		"initial_quantity":    10,
		"manufacturing_date":  "2020-04-04",
		"manufacturing_hour":  10,
		"minimum_temperature": 5,
		"product_id":          1,
		"section_id":          1,
	}
	expectedErr := repository.ErrEntityAlreadyExists
	p.mock.On("Register", mock.AnythingOfType("models.ProductBatch")).Return(models.ProductBatch{}, expectedErr)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, p.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	p.handler.PostProductBatch(recorder, request)

	// Assert
	p.Equal(http.StatusConflict, recorder.Code)
	p.Contains(recorder.Body.String(), expectedErr.Error())

}

// Run the test suite
func TestProductBatchHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductBatchHandlerTestSuite))
}
