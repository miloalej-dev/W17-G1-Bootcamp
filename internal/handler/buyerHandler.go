package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bootcamp-go/web/response"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/buyerService"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"net/http"
)

func NewBuyerDefault(sv buyerService.BuyerService) *BuyerDefault {
	return &BuyerDefault{sv: sv}
}

// BuyerDefault is a struct with methods that represent handlers for buyers
type BuyerDefault struct {
	// sv is the service that will be used by the handler
	sv buyerService.BuyerService
}

// GetAll is a method that returns a handler for the route GET /buyers
func (h *BuyerDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Consultando buyers")
		v, err := h.sv.FindAll()
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

func (h *BuyerDefault) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		value, err := h.sv.Create(DocToAttributes(bodyRequest))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
		}
		response.JSON(w, http.StatusOK, BuyerToDoc(value))

	}

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
