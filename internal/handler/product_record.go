package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
	"strconv"
)

func NewProductRecordHandler(sv service.ProductRecordService) *ProductRecordHandler {
	return &ProductRecordHandler{
		service: sv,
	}
}

type ProductRecordHandler struct {
	// sv is the service that will be used by the handler
	service service.ProductRecordService
}

// GetAll is a method that returns a handler for the route GET /buyers
func (h *ProductRecordHandler) GetProductRecords(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	value, err := h.service.RetrieveAll()
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(value, http.StatusOK))

}

func (h *ProductRecordHandler) GetProductRecord(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		_ = render.Render(w, r, response.NewErrorResponse(errors.New("invalid request").Error(), http.StatusBadRequest))
		return
	}
	value, err := h.service.Retrieve(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}
	_ = render.Render(w, r, response.NewResponse(value, http.StatusOK))

}

func (h *ProductRecordHandler) PostProductRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bodyRequest := &request.ProductRecordRequest{}

	if err := render.Bind(r, bodyRequest); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusUnprocessableEntity))
		return
	}

	productRecord := models.ProductRecord{
		LastUpdate:    *bodyRequest.LastUpdate,
		PurchasePrice: *bodyRequest.PurchasePrice,
		SalePrice:     *bodyRequest.SalePrice,
		ProductId:     *bodyRequest.ProductId,
	}

	fmt.Println(productRecord)
	value, err := h.service.Register(productRecord)

	if err != nil {

		if errors.Is(err, service.ErrProductIdConflict) {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusConflict))
			return
		}
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = render.Render(w, r, response.NewResponse(value, http.StatusCreated))

}

func (h *ProductRecordHandler) PatchProductRecord(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		_ = render.Render(w, r, response.NewErrorResponse(errors.New("invalid request").Error(), http.StatusBadRequest))
		return
	}

	var fields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(errors.New("unexpected JSON format, check the request body").Error(), http.StatusBadRequest))
		return
	}

	value, err := h.service.PartialModify(id, fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = render.Render(w, r, response.NewResponse(value, http.StatusOK))

}

func (h *ProductRecordHandler) DeleteProductRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		_ = render.Render(w, r, response.NewErrorResponse(errors.New("invalid request").Error(), http.StatusBadRequest))
		return
	}
	err = h.service.Remove(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = render.Render(w, r, response.NewResponse(nil, http.StatusNoContent))

}
