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
	var id int
	var err error
	if idParam != "" {
		id, err = strconv.Atoi(idParam)
		if err != nil {
			_ = render.Render(w, r, response.NewErrorResponse(idParam, http.StatusBadRequest))
			return
		}
	} else {
		id = 0 // valor por defecto si no hay query param
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

	orDetail := make([]models.OrderDetail, 0)
	for _, or := range *data.OrderDetails {
		ordt := models.OrderDetail{
			Id:               0,
			Quantity:         or.Quantity,
			CleanLinesStatus: or.CleanLinesStatus,
			Temperature:      or.Temperature,
			ProductRecordID:  or.ProductRecordID,
		}
		orDetail = append(orDetail, ordt)
	}

	purchaseOrders := models.PurchaseOrder{
		OrderNumber:   *data.OrderNumber,
		OrderDate:     *data.OrderDate,
		TracingCode:   *data.TracingCode,
		BuyerID:       *data.BuyersID,
		WarehouseID:   *data.WarehousesID,
		CarrierID:     *data.CarriersID,
		OrderStatusID: *data.OrderStatusID,
		OrderDetails:  &orDetail,
	}

	createdPurchaseOrder, err := h.sv.Register(purchaseOrders)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusConflict))
		return
	}
	_ = render.Render(w, r, response.NewResponse(createdPurchaseOrder, http.StatusOK))

}
