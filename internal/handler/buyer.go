package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
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
func (h *BuyerHandler) GetBuyers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		fmt.Println("Consultando buyers")
		value, err := h.service.RetrieveAll()
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    value,
		})

	}

}

func (h *BuyerHandler) GetBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}
		value, err := h.service.Retrieve(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}
		response.JSON(w, http.StatusOK, value)

	}
}

func (h *BuyerHandler) PostBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var bodyRequest models.Buyer

		err := json.NewDecoder(r.Body).Decode(&bodyRequest)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		//err = PostValidator(DocToAttributes(bodyRequest))
		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, nil)
			return
		}

		value, err := h.service.Register(bodyRequest)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
		}
		response.JSON(w, http.StatusOK, value)

	}

}

func (h *BuyerHandler) DeleteBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
		}
		err = h.service.Remove(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
		}
		response.JSON(w, http.StatusNoContent, nil)

	}
}

func (h *BuyerHandler) PatchBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
		}

		var bodyRequest models.Buyer
		err = json.NewDecoder(r.Body).Decode(&bodyRequest)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
		}

		value, err := PutValidator(bodyRequest, id, h)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
		}

		buyer, err := h.service.Modify(value)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
		}
		response.JSON(w, http.StatusOK, buyer)
	}

}
func PutValidator(buyer models.Buyer, id int, h *BuyerHandler) (b models.Buyer, err error) {

	_, err = h.service.Retrieve(id)
	if err != nil {
		return models.Buyer{}, err
	}

	if buyer.CardNumberId != "" && buyer.FirstName != "" && buyer.LastName != "" {
		return models.Buyer{}, errors.New("x")
	}

	return b, nil

}

func PostValidator(buyer models.Buyer) error {
	if buyer.FirstName == "" || buyer.LastName == "" || buyer.CardNumberId == "" {
		return errors.New("First name, Last name or CardNumberId are empty")
	}
	return nil
}
