package handler

import (
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
	"strconv"
)

func NewPurchaseOrderDefault(sv service.PurchaseOrderService) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{sv: sv}
}

type PurchaseOrderHandler struct {
	sv service.PurchaseOrderService
}

func (h *PurchaseOrderHandler) GetPurchaseOrdersReport(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse("Invalid buyer ID", http.StatusBadRequest))
		return
	}

	// Suponiendo que tienes un servicio llamado h.sv con método ReportByBuyerID
	report, err := h.sv.RetrieveByBuyer(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	_ = render.Render(w, r, response.NewResponse(report, http.StatusOK))
}

func (h *PurchaseOrderHandler) PostPurchaseOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := &request.PurchaseOrderRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	purchaseOrders := models.PurchaseOrder{
		Id:            *data.Id,
		OrderNumber:   *data.OrderNumber,
		OrderDate:     *data.OrderDate,
		TracingCode:   *data.TracingCode,
		BuyersID:      *data.BuyersID,
		WarehousesID:  *data.WarehousesID,
		CarriersID:    *data.CarriersID,
		OrderStatusID: *data.OrderStatusID,
	}

	createdPurchaseOrder, err := h.sv.Register(purchaseOrders)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	_ = render.Render(w, r, response.NewResponse(createdPurchaseOrder, http.StatusOK))

}
