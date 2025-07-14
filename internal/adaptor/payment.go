package adaptor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project-app-bioskop-golang-homework-rahmadhany/internal/dto"
	"project-app-bioskop-golang-homework-rahmadhany/internal/usecase"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"

	"go.uber.org/zap"
)

type PaymentHandler struct {
	Payment usecase.PaymentService
	Logger  *zap.Logger
	Config  utils.Configuration
}

func NewPaymentHandler(payment usecase.PaymentService, logger *zap.Logger, config utils.Configuration) PaymentHandler {
	return PaymentHandler{
		Payment: payment,
		Logger:  logger,
		Config:  config,
	}
}

func (h *PaymentHandler) GetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	methods, err := h.Payment.ListMethods(r.Context())
	if err != nil {
		utils.WriteError(w, "Failed to retrieve payment methods", http.StatusInternalServerError)
	}
	utils.WriteSuccess(w, "", http.StatusOK, methods, nil)
}

func (h *PaymentHandler) Pay(w http.ResponseWriter, r *http.Request) {
	var req dto.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	txnID, err := h.Payment.ProcessPayment(r.Context(), req)
	if err != nil {
		fmt.Println(err.Error())
		utils.WriteError(w, "Payment failed due to insufficient funds."+err.Error(), http.StatusPaymentRequired)
		return
	}

	data := map[string]interface{}{
		"transactionId": txnID,
		"bookingId":     req.BookingID,
	}
	utils.WriteSuccess(w, "Payment successful.", http.StatusOK, data, nil)

}
