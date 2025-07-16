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

type InboundOrderHandler struct {
	service service.InboundOrderService
}

func NewInboundOrderHandler(service service.InboundOrderService) *InboundOrderHandler {
	return &InboundOrderHandler{
		service: service,
	}
}

func (h *InboundOrderHandler) GetInboundOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	inboundOrders, err := h.service.RetrieveAll()

	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(inboundOrders, http.StatusOK))
}

func (h *InboundOrderHandler) GetInboundOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	inboundOrder, err := h.service.Retrieve(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	_ = render.Render(w, r, response.NewResponse(inboundOrder, http.StatusOK))

}

func (h *InboundOrderHandler) PostInboundOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := &request.InboundOrder{}

	err := render.Bind(r, data)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
	}

	inboundOrder := models.InboundOrder{
		OrderNumber:    *data.OrderNumber,
		EmployeeId:     *data.EmployeeId,
		ProductBatchId: *data.ProductBatchId,
		WarehouseId:    *data.WarehouseId,
	}

	createdInboundOrder, err := h.service.Register(inboundOrder)

	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(createdInboundOrder, http.StatusCreated))
}

func (h *InboundOrderHandler) PutInboundOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	data := &request.InboundOrder{}

	err = render.Bind(r, data)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
	}

	inboundOrder := models.InboundOrder{
		Id:             id,
		OrderNumber:    *data.OrderNumber,
		EmployeeId:     *data.EmployeeId,
		ProductBatchId: *data.ProductBatchId,
		WarehouseId:    *data.WarehouseId,
	}

	updatedInboundOrder, err := h.service.Modify(inboundOrder)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(updatedInboundOrder, http.StatusOK))
}

func (h *InboundOrderHandler) PatchInboundOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
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

	updatedInboundOrder, err := h.service.PartialModify(id, fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		http.Error(w, "Failed to update seller", http.StatusInternalServerError)
		return
	}

	_ = render.Render(w, r, response.NewResponse(updatedInboundOrder, http.StatusOK))
}

func (h *InboundOrderHandler) DeleteInboundOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err = h.service.Remove(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(nil, http.StatusNoContent))
}
