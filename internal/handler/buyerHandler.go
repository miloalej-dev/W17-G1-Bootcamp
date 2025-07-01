package handler

import (
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
			response.JSON(w, http.StatusInternalServerError, nil)
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
