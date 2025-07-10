package handler

import (
	"encoding/json"

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

func NewBuyerHandler(sv service.BuyerService) *BuyerHandler {
	return &BuyerHandler{
		service: sv,
	}
}

// BuyerHandler is a struct with methods that represent handlers for buyers
type BuyerHandler struct {
	// sv is the service that will be used by the handler
	service service.BuyerService
}

// GetAll is a method that returns a handler for the route GET /buyers
func (h *BuyerHandler) GetBuyers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Consultando buyers")
	value, err := h.service.RetrieveAll()
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	_ = render.Render(w, r, response.NewResponse(value, http.StatusOK))

}

func (h *BuyerHandler) GetBuyer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	value, err := h.service.Retrieve(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}
	_ = render.Render(w, r, response.NewResponse(value, http.StatusOK))

}

func (h *BuyerHandler) PostBuyer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bodyRequest := &request.BuyerRequest{}

	if err := render.Bind(r, bodyRequest); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
	}

	buyer := models.Buyer{
		CardNumberId: *bodyRequest.CardNumberId,
		FirstName:    *bodyRequest.FirstName,
		LastName:     *bodyRequest.LastName,
	}

	value, err := h.service.Register(buyer)

	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = render.Render(w, r, response.NewResponse(value, http.StatusOK))

}

func (h *BuyerHandler) DeleteBuyer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
	}
	err = h.service.Remove(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
	}
	_ = render.Render(w, r, response.NewResponse(nil, http.StatusNoContent))

}

func (h *BuyerHandler) PatchBuyer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
	}

	var fields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	buyer, err := h.service.PartialModify(id, fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
	}
	_ = render.Render(w, r, response.NewResponse(buyer, http.StatusOK))

}
