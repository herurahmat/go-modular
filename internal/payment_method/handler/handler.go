package handler

import (
	"go-moduler/internal/payment_method/service"
	pkgHttp "go-moduler/pkg/http"
	"net/http"
)

type Handler struct {
	service service.PaymentMethodService
}

func (h Handler) ModemHandler(w http.ResponseWriter, r *http.Request) {

	data, err := h.service.PurchaseModem(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(pkgHttp.Response{
			Status: "NO",
			Data:   nil,
		}.ToJSON())

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(pkgHttp.Response{
		Status: "OK",
		Data:   responsePaymentMethod(data),
	}.ToJSON())
}

func NewHandler(service service.PaymentMethodService) *Handler {

	return &Handler{service}
}
