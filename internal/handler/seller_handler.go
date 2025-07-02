package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
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
		http.Error(w, "Failed to retrieve sellers", http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, sellers)
}

// GetSeller handles GET requests to retrieve a seller by ID
func (h *SellerHandler) GetSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	seller, err := h.service.GetSellerById(id)
	if err != nil {
		http.Error(w, "Seller not found", http.StatusNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, seller)
}

// PostSeller handles POST requests to create a new seller
func (h *SellerHandler) PostSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var seller models.Seller

	err := json.NewDecoder(r.Body).Decode(&seller)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdSeller, err := h.service.RegisterSeller(seller)
	if err != nil {
		http.Error(w, "Failed to create seller", http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, createdSeller)
}

// PutSeller handles PUT requests to update a seller
func (h *SellerHandler) PutSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	var seller models.Seller
	err = json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	seller.Id = id
	updatedSeller, err := h.service.ModifySeller(seller)
	if err != nil {
		http.Error(w, "Failed to update seller", http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, updatedSeller)
}

// PatchSeller handles PATCH requests to partially update a seller
func (h *SellerHandler) PatchSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	var fields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedSeller, err := h.service.UpdateSellerFields(id, fields)
	if err != nil {
		http.Error(w, "Failed to update seller", http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, updatedSeller)
}

// DeleteSeller handles DELETE requests to remove a seller
func (h *SellerHandler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	err = h.service.RemoveSeller(id)
	if err != nil {
		http.Error(w, "Failed to delete seller", http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
}
