package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
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

type ProductServiceMock struct {
	mock.Mock
}

func (p *ProductServiceMock) RetrieveRecordsCountByProductId(id int) (models.ProductReport, error) {
	args := p.Called(id)
	return args.Get(0).(models.ProductReport), args.Error(1)
}

func (p *ProductServiceMock) RetrieveRecordsCount() ([]models.ProductReport, error) {
	args := p.Called()
	return args.Get(0).([]models.ProductReport), args.Error(1)
}

type ProductHandlerTestSuite struct {
	suite.Suite
	mock    *ProductServiceMock
	handler *ProductDefault
	path    string
}

// Mock methods for SectionService

func (p *ProductServiceMock) RetrieveAll() ([]models.Product, error) {
	args := p.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (p *ProductServiceMock) Retrieve(id int) (models.Product, error) {
	args := p.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (p *ProductServiceMock) Register(product models.Product) (models.Product, error) {
	args := p.Called(product)
	return args.Get(0).(models.Product), args.Error(1)
}

func (p *ProductServiceMock) Modify(product models.Product) (models.Product, error) {
	args := p.Called(product)
	return args.Get(0).(models.Product), args.Error(1)
}

func (p *ProductServiceMock) PartialModify(id int, fields map[string]any) (models.Product, error) {
	args := p.Called(id, fields)
	return args.Get(0).(models.Product), args.Error(1)
}

func (p *ProductServiceMock) Remove(id int) error {
	args := p.Called(id)
	return args.Error(0)
}
func (p *ProductHandlerTestSuite) SetupTest() {
	p.mock = new(ProductServiceMock)
	p.handler = NewProductDefault(p.mock)
	p.path = "/api/v1/products"
}

func (p *ProductHandlerTestSuite) TestGetProducts_Ok() {
	var expectedBody []byte

	var expectedProducts = []models.Product{
		{
			Id:                             1,
			ProductCode:                    "mock1",
			Description:                    "mock1",
			Width:                          68.64,
			Height:                         185.04,
			Length:                         185.62,
			NetWeight:                      2.93,
			ExpirationRate:                 8.62,
			RecommendedFreezingTemperature: -33.84,
			FreezingRate:                   0.62,
			ProductTypeId:                  6,
		},
		{
			Id:                             2,
			ProductCode:                    "mock2",
			Description:                    "mock2",
			Width:                          68.64,
			Height:                         185.04,
			Length:                         185.62,
			NetWeight:                      2.93,
			ExpirationRate:                 8.62,
			RecommendedFreezingTemperature: -33.84,
			FreezingRate:                   0.62,
			ProductTypeId:                  6,
		},
	}

	expectedResponse := response.Response{Data: expectedProducts}
	p.mock.On("RetrieveAll").Return(expectedProducts, nil)

	request := httptest.NewRequest(http.MethodGet, p.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	p.handler.GetProducts(recorder, request)

	var resp response.Response
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	p.Equal(http.StatusOK, recorder.Code)
	p.Len(resp.Data, 2)
	p.JSONEq(string(expectedBody), recorder.Body.String())

}

func (p *ProductHandlerTestSuite) TestGetProducts_InternalError() {
	// Arrange

	expectedError := errors.New("not found")
	expectedResponse := response.Response{
		Message:    expectedError.Error(),
		StatusCode: http.StatusNotFound,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	p.mock.On("RetrieveAll").Return([]models.Product{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, p.path, nil)
	recorder := httptest.NewRecorder()

	// Act
	p.handler.GetProducts(recorder, request)

	// Assert
	p.Equal(http.StatusNotFound, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())
}

func (p *ProductHandlerTestSuite) TestGetProduct_Ok() {
	// Arrange
	product := models.Product{
		Id:                             1,
		ProductCode:                    "mock1",
		Description:                    "mock1",
		Width:                          68.64,
		Height:                         185.04,
		Length:                         185.62,
		NetWeight:                      2.93,
		ExpirationRate:                 8.62,
		RecommendedFreezingTemperature: -33.84,
		FreezingRate:                   0.62,
		ProductTypeId:                  6,
	}
	expectedResponse := response.Response{Data: product}
	expectedBody, _ := json.Marshal(expectedResponse)

	p.mock.On("Retrieve", product.Id).Return(product, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", p.path, product.Id), nil)
	recorder := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", strconv.Itoa(product.Id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	// Act
	p.handler.GetProduct(recorder, request)

	// Assert
	p.Equal(http.StatusOK, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())
}

func (p *ProductHandlerTestSuite) TestGetProduct_InvalidID() {
	// Arrange
	expectedResponse := response.Response{
		Message:    "strconv.Atoi: parsing \"invalid\": invalid syntax",
		StatusCode: http.StatusBadRequest,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	request := httptest.NewRequest(http.MethodGet, p.path+"/invalid", nil)
	recorder := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalid")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	// Act
	p.handler.GetProduct(recorder, request)

	// Assert
	p.Equal(http.StatusBadRequest, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())
}

func (p *ProductHandlerTestSuite) TestGetSection_NotFound() {
	// Arrange
	id := 999
	expectedErr := errors.New("section not found")
	expectedResponse := response.Response{
		Message:    expectedErr.Error(),
		StatusCode: http.StatusNotFound,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	p.mock.On("Retrieve", id).Return(models.Product{}, expectedErr)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", p.path, id), nil)
	recorder := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

	// Act
	p.handler.GetProduct(recorder, request)

	// Assert
	p.Equal(http.StatusNotFound, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())
}

func (p *ProductHandlerTestSuite) TestPostProduct_Ok() {
	// Arrange
	requestBody := map[string]interface{}{
		"product_code":                     "new product",
		"description":                      "new product",
		"width":                            68.64,
		"height":                           185.04,
		"length":                           185.62,
		"net_weight":                       2.93,
		"expiration_rate":                  8.62,
		"recommended_freezing_temperature": -33.84,
		"freezing_rate":                    0.62,
		"product_type_id":                  6,
	}
	returnedProduct := models.Product{Id: 1, ProductCode: "new product"}
	expectedResponse := response.Response{Data: returnedProduct}
	expectedBody, _ := json.Marshal(expectedResponse)

	p.mock.On("Register", mock.AnythingOfType("models.Product")).Return(returnedProduct, nil)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, p.path, bytes.NewBuffer(requestBodyBytes))

	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	p.handler.PostProduct(recorder, request)

	// Assert
	p.Equal(http.StatusCreated, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())
}

func (p *ProductHandlerTestSuite) TestPostProduct_Conflict_AlreadyExists() {
	// Arrange
	requestBody := map[string]interface{}{
		"id":                               1,
		"product_code":                     "mock1",
		"description":                      "mock1",
		"width":                            68.64,
		"height":                           185.04,
		"length":                           185.62,
		"net_weight":                       2.93,
		"expiration_rate":                  8.62,
		"recommended_freezing_temperature": -33.84,
		"freezing_rate":                    0.62,
		"product_type_id":                  6,
	}

	expectedErr := repository.ErrEntityAlreadyExists
	expectedResponse := response.Response{
		Message:    expectedErr.Error(),
		StatusCode: http.StatusBadRequest,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	p.mock.On("Register", mock.AnythingOfType("models.Product")).Return(models.Product{}, expectedErr)

	requestBodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, p.path, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	p.handler.PostProduct(recorder, request)

	// Assert
	p.Equal(http.StatusBadRequest, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())
}

func (p *ProductHandlerTestSuite) TestPostProduct_BindError() {
	// Arrange
	invalidBody := []byte(`{ "product_code": 123 }`) // string esperado, número enviado

	expectedResponse := response.Response{
		Message:    "json: cannot unmarshal number into Go struct field ProductRequest.product_code of type string",
		StatusCode: http.StatusBadRequest,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	request := httptest.NewRequest(http.MethodPost, p.path, bytes.NewBuffer(invalidBody))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	p.handler.PostProduct(recorder, request)

	// Assert
	p.Equal(http.StatusBadRequest, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())
}

func (p *ProductHandlerTestSuite) TestPatchSection_Ok() {
	// Arrange
	id := 1
	fields := map[string]interface{}{
		"product_code": "updatedMock1",
	}

	expectedProduct := models.Product{
		Id:                             2,
		ProductCode:                    "oldMock1",
		Description:                    "oldMock1",
		Width:                          68.64,
		Height:                         185.04,
		Length:                         185.62,
		NetWeight:                      2.93,
		ExpirationRate:                 8.62,
		RecommendedFreezingTemperature: -33.84,
		FreezingRate:                   0.62,
		ProductTypeId:                  6,
	}
	expectedResponse := response.Response{Data: expectedProduct}
	expectedBody, _ := json.Marshal(expectedResponse)

	p.mock.On("PartialModify", id, fields).Return(expectedProduct, nil)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", p.path, id), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")

	// Add chi context
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	recorder := httptest.NewRecorder()

	// Act
	p.handler.PatchProduct(recorder, request)

	// Assert
	p.Equal(http.StatusOK, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())
}

func (p *ProductHandlerTestSuite) TestPatchProduct_NotFound() {
	// Arrange
	productID := 999
	fields := map[string]interface{}{}

	expectedErr := repository.ErrProductNotFound

	p.mock.On("PartialModify", productID, fields).Return(models.Product{}, expectedErr)

	requestBodyBytes, _ := json.Marshal(fields)
	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", p.path, productID), bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(productID))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	recorder := httptest.NewRecorder()

	// Act
	p.handler.PatchProduct(recorder, request)

	// Assert
	p.Equal(http.StatusNotFound, recorder.Code)
	p.Contains(recorder.Body.String(), expectedErr.Error())
}

func (p *ProductHandlerTestSuite) TestPatchProduct_InternalServerError() {
	// Arrange
	id := 999
	invalidBody := []byte(`{"product_code":`)

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", p.path, id), bytes.NewBuffer(invalidBody))
	request.Header.Set("Content-Type", "application/json")

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))
	recorder := httptest.NewRecorder()
	// Act
	p.handler.PatchProduct(recorder, request)

	// Assert
	p.Equal(http.StatusBadRequest, recorder.Code)
	p.Contains(recorder.Body.String(), "Invalid request body")
}

func (p *ProductHandlerTestSuite) TestPatchSection_BadRequest_InvalidID() {
	// Arrange
	requestBody := []byte(`{"product_code": "updatedMock1"}`)

	request := httptest.NewRequest(http.MethodPatch, p.path+"/bad_id", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "bad_id")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	recorder := httptest.NewRecorder()

	// Act
	p.handler.PatchProduct(recorder, request)

	// Assert
	p.Equal(http.StatusBadRequest, recorder.Code)
	p.Contains(recorder.Body.String(), "Invalid ID format")
}

func (p *ProductHandlerTestSuite) TestPatchProduct_BadRequest_InvalidBody() {
	// Arrange
	id := 4
	invalidBody := []byte(`{"product_code":`)

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", p.path, id), bytes.NewBuffer(invalidBody))
	request.Header.Set("Content-Type", "application/json")

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

	recorder := httptest.NewRecorder()

	// Act
	p.handler.PatchProduct(recorder, request)

	// Assert
	p.Equal(http.StatusBadRequest, recorder.Code)
	p.Contains(recorder.Body.String(), "Invalid request body")
}

func (p *ProductHandlerTestSuite) TestDeleteSection_NoContent() {
	// Arrange
	id := 1
	expectedResponse := response.Response{Data: "product Deleted"}
	expectedBody, _ := json.Marshal(expectedResponse)

	p.mock.On("Remove", id).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/sections/%d", id), nil)
	rec := httptest.NewRecorder()

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	// Act
	p.handler.DeleteProduct(rec, req)

	// Assert
	p.Equal(http.StatusNoContent, rec.Code)
	p.JSONEq(string(expectedBody), rec.Body.String())
}

func (p *ProductHandlerTestSuite) TestDeleteSection_NotFound() {
	// Arrange
	id := 99
	expectedErr := errors.New("product not found")
	expectedResponse := response.Response{Message: expectedErr.Error()}
	expectedBody, _ := json.Marshal(expectedResponse)

	p.mock.On("Remove", id).Return(expectedErr)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/sections/%d", id), nil)
	rec := httptest.NewRecorder()

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", strconv.Itoa(id))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	// Act
	p.handler.DeleteProduct(rec, req)

	// Assert
	p.Equal(http.StatusNotFound, rec.Code)
	p.JSONEq(string(expectedBody), rec.Body.String())
}

func (p *ProductHandlerTestSuite) TestDeleteSection_BadRequest() {
	// Arrange
	expectedErr := "strconv.Atoi: parsing \"abc\": invalid syntax"
	expectedResponse := response.Response{Message: expectedErr}
	expectedBody, _ := json.Marshal(expectedResponse)

	req := httptest.NewRequest(http.MethodDelete, "/products/abc", nil)
	rec := httptest.NewRecorder()

	// Simular contexto con ID inválido
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	// Act
	p.handler.DeleteProduct(rec, req)

	// Assert
	p.Equal(http.StatusBadRequest, rec.Code)
	p.JSONEq(string(expectedBody), rec.Body.String())
}

func (p *ProductHandlerTestSuite) TestGetProductReport_AllProducts_Success() {
	expected := []models.ProductReport{
		{Id: 1, Description: "Coca", RecordsCount: 2},
		{Id: 2, Description: "pepsi", RecordsCount: 4},
	}

	expectedData := response.Response{Data: expected, StatusCode: http.StatusOK}

	p.mock.On("RetrieveRecordsCount").Return(expected, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(p.path, "/reportRecords"), nil)
	recorder := httptest.NewRecorder()

	p.handler.GetProductReport(recorder, request)

	expectedResponse, _ := json.Marshal(expectedData)

	p.Equal(expectedData.StatusCode, recorder.Code)
	p.JSONEq(string(expectedResponse), recorder.Body.String())

}

func (p *ProductHandlerTestSuite) TestGetProductReport_AllProducts_InternalServerError() {

	var expectedBody []byte
	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	p.mock.On("RetrieveRecordsCount").Return([]models.ProductReport{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(p.path, "/reportRecords"), nil)
	recorder := httptest.NewRecorder()

	p.handler.GetProductReport(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	p.Equal(expectedResponse.StatusCode, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())
}

func (p *ProductHandlerTestSuite) TestGetProductReportById_Success() {

	var expectedBody []byte
	id := 1
	expected := models.ProductReport{
		Id: id, Description: "Coca", RecordsCount: 2,
	}

	expectedResponse := response.Response{Data: expected, StatusCode: http.StatusOK}

	p.mock.On("RetrieveRecordsCountByProductId", id).Return(expected, nil)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(p.path, "/reportRecords?id=", id), nil)
	recorder := httptest.NewRecorder()

	p.handler.GetProductReport(recorder, request)
	expectedBody, _ = json.Marshal(expectedResponse)

	p.Equal(expectedResponse.StatusCode, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())

}

func (p *ProductHandlerTestSuite) TestGetProductReportById_BadRequest() {

	var expectedBody []byte
	id := "asc"

	expectedError := errors.New("invalid request")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusBadRequest}

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(p.path, "/reportRecords?id=", id), nil)
	recorder := httptest.NewRecorder()

	p.handler.GetProductReport(recorder, request)
	expectedBody, _ = json.Marshal(expectedResponse)

	p.Equal(expectedResponse.StatusCode, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())

}

func (p *ProductHandlerTestSuite) TestGetProductReportById_NotFound() {

	var expectedBody []byte
	id := 999

	expectedResponse := response.Response{Message: service.ErrProductNotFound.Error(), StatusCode: http.StatusNotFound}

	p.mock.On("RetrieveRecordsCountByProductId", id).Return(models.ProductReport{}, service.ErrProductNotFound)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(p.path, "/reportRecords?id=", id), nil)
	recorder := httptest.NewRecorder()

	p.handler.GetProductReport(recorder, request)
	expectedBody, _ = json.Marshal(expectedResponse)

	p.Equal(expectedResponse.StatusCode, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())

}

func (p *ProductHandlerTestSuite) TestGetProductReportById_InternalServerError() {

	var expectedBody []byte
	id := 1
	expectedError := errors.New("something went wrong")
	expectedResponse := response.Response{Message: expectedError.Error(), StatusCode: http.StatusInternalServerError}

	p.mock.On("RetrieveRecordsCountByProductId", id).Return(models.ProductReport{}, expectedError)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprint(p.path, "/reportRecords?id=", id), nil)
	recorder := httptest.NewRecorder()

	p.handler.GetProductReport(recorder, request)

	expectedBody, _ = json.Marshal(expectedResponse)

	// Assert
	p.Equal(expectedResponse.StatusCode, recorder.Code)
	p.JSONEq(string(expectedBody), recorder.Body.String())
}

// Run the test suite
func TestProductHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}
