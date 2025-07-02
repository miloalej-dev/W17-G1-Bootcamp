package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/buyerService"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"net/http"
	"strconv"
)

func NewBuyerHandler(sv buyerService.BuyerService) *BuyerHandler {
	return &BuyerHandler{
		service: sv,
	}
}

// BuyerHandler is a struct with methods that represent handlers for buyers
type BuyerHandler struct {
	// sv is the service that will be used by the handler
	service buyerService.BuyerService
}

// GetAll is a method that returns a handler for the route GET /buyers
func (h *BuyerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		fmt.Println("Consultando buyers")
		v, err := h.service.FindAll()
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		data := make(map[int]models.BuyerDoc)
		for key, value := range v {
			data[key] = models.BuyerDoc{
				Id:           value.Id,
				CardNumberId: value.CardNumberId,
				FirstName:    value.FirstName,
				LastName:     value.LastName,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}

}

func (h *BuyerHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}
		value, err := h.service.FindById(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}
		response.JSON(w, http.StatusOK, BuyerToDoc(value))

	}
}

func (h *BuyerHandler) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var bodyRequest models.BuyerDoc

		err := json.NewDecoder(r.Body).Decode(&bodyRequest)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		err = PostValidator(DocToAttributes(bodyRequest))
		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, nil)
			return
		}

		value, err := h.service.Create(DocToAttributes(bodyRequest))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
		}
		response.JSON(w, http.StatusOK, BuyerToDoc(value))

	}

}

func (h *BuyerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
		}
		value, err := h.service.Delete(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
		}
		response.JSON(w, http.StatusNoContent, BuyerToDoc(value))

	}
}

func (h *BuyerHandler) Patch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
		}

		var bodyRequest models.BuyerDoc
		err = json.NewDecoder(r.Body).Decode(&bodyRequest)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
		}

		value, err := PutValidator(DocToAttributes(bodyRequest), id, h)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
		}

		buyer, err := h.service.Update(value)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
		}
		response.JSON(w, http.StatusOK, BuyerToDoc(buyer))
	}

}
func PutValidator(buyer models.BuyerAtributtes, id int, h *BuyerHandler) (b models.Buyer, err error) {

	value, err := h.service.FindById(id)
	if err != nil {
		return models.Buyer{}, err
	}

	if buyer.CardNumberId != "" && buyer.FirstName != "" && buyer.LastName != "" {
		return models.Buyer{}, errors.New("<UNK>")
	}

	switch {

	case buyer.FirstName != "":
		b = models.Buyer{
			Id: id,
			BuyerAtributtes: models.BuyerAtributtes{
				CardNumberId: value.CardNumberId,
				FirstName:    buyer.FirstName,
				LastName:     value.LastName,
			},
		}
	case buyer.LastName != "":
		b = models.Buyer{
			Id: id,
			BuyerAtributtes: models.BuyerAtributtes{
				CardNumberId: value.CardNumberId,
				FirstName:    value.FirstName,
				LastName:     buyer.LastName,
			},
		}

	case buyer.CardNumberId != "":
		b = models.Buyer{
			Id: id,
			BuyerAtributtes: models.BuyerAtributtes{
				CardNumberId: buyer.CardNumberId,
				FirstName:    value.FirstName,
				LastName:     value.LastName,
			},
		}

	}

	return b, nil

}

func PostValidator(buyer models.BuyerAtributtes) error {
	if buyer.FirstName == "" || buyer.LastName == "" || buyer.CardNumberId == "" {
		return errors.New("First name or Last name is empty")
	}
	return nil
}

func BuyerToDoc(buyer *models.Buyer) (b models.BuyerDoc) {
	b = models.BuyerDoc{
		Id:           buyer.Id,
		CardNumberId: buyer.CardNumberId,
		FirstName:    buyer.FirstName,
		LastName:     buyer.LastName,
	}

	return b

}

func DocToBuyer(buyer models.BuyerDoc) (b models.Buyer) {

	b = models.Buyer{
		Id: buyer.Id,
		BuyerAtributtes: models.BuyerAtributtes{
			CardNumberId: buyer.CardNumberId,
			FirstName:    buyer.FirstName,
			LastName:     buyer.LastName,
		},
	}
	return b
}

func DocToAttributes(buyer models.BuyerDoc) (b models.BuyerAtributtes) {

	b = models.BuyerAtributtes{
		CardNumberId: buyer.CardNumberId,
		FirstName:    buyer.FirstName,
		LastName:     buyer.LastName,
	}
	return b
}
