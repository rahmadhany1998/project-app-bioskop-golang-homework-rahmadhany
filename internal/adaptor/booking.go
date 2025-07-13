package adaptor

import (
	"encoding/json"
	"net/http"
	"project-app-bioskop-golang-homework-rahmadhany/internal/dto"
	"project-app-bioskop-golang-homework-rahmadhany/internal/usecase"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/middleware"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"

	"go.uber.org/zap"
)

type BookingHandler struct {
	Booking usecase.BookingService
	Logger  *zap.Logger
	Config  utils.Configuration
}

func NewBookingHandler(booking usecase.BookingService, logger *zap.Logger, config utils.Configuration) BookingHandler {
	return BookingHandler{
		Booking: booking,
		Logger:  logger,
		Config:  config,
	}
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req dto.BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middleware.ContextUserID).(int)
	if !ok {
		utils.WriteError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	booking, err := h.Booking.CreateBooking(r.Context(), req, userID)
	if err != nil {
		if err.Error() == "seat already booked" {
			utils.WriteError(w, "The selected seat is already booked.", http.StatusConflict)
			return
		}

		utils.WriteError(w, "Failed to create booking"+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteSuccess(w, "Booking Confirmed", http.StatusOK, booking, nil)
}
