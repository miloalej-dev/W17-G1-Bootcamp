package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
	"strconv"
)

type SellerHandler struct {
	service *service.SellerService
}

func NewSellerHandler(service *service.SellerService) *SellerHandler {
	return &SellerHandler{
		service: service,
	}
}

// GetSellers handles GET requests to retrieve all sellers
func (h *SellerHandler) GetSellers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sellers, err := h.service.GetSellers()

	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(sellers, http.StatusOK))
}

// GetSeller handles GET requests to retrieve a seller by ID
func (h *SellerHandler) GetSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	seller, err := h.service.GetSellerById(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	_ = render.Render(w, r, response.NewResponse(seller, http.StatusOK))
}

// PostSeller handles POST requests to create a new seller
func (h *SellerHandler) PostSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := &request.SellerRequest{}

	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
	}

	seller := models.Seller{
		Name:      *data.Name,
		Address:   *data.Address,
		Telephone: *data.Telephone,
	}

	createdSeller, err := h.service.RegisterSeller(seller)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(createdSeller, http.StatusCreated))
}

// PutSeller handles PUT requests to update a seller
func (h *SellerHandler) PutSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	data := &request.SellerRequest{}

	err = render.Bind(r, data)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
	}

	seller := models.Seller{
		Id:        id,
		Name:      *data.Name,
		Address:   *data.Address,
		Telephone: *data.Telephone,
	}

	updatedSeller, err := h.service.ModifySeller(seller)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(updatedSeller, http.StatusOK))
}

// PatchSeller handles PATCH requests to partially update a seller
func (h *SellerHandler) PatchSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	var fields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	updatedSeller, err := h.service.UpdateSellerFields(id, fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		http.Error(w, "Failed to update seller", http.StatusInternalServerError)
		return
	}

	_ = render.Render(w, r, response.NewResponse(updatedSeller, http.StatusOK))
}

// DeleteSeller handles DELETE requests to remove a seller
func (h *SellerHandler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err = h.service.RemoveSeller(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	render.Status(r, http.StatusNoContent)
}
