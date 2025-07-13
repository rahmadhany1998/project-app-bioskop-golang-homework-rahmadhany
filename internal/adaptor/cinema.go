package adaptor

import (
	"net/http"
	"project-app-bioskop-golang-homework-rahmadhany/internal/usecase"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type CinemaHandler struct {
	Cinema usecase.CinemaService
	Logger *zap.Logger
	Config utils.Configuration
}

func NewCinemaHandler(cinema usecase.CinemaService, logger *zap.Logger, config utils.Configuration) CinemaHandler {
	return CinemaHandler{
		Cinema: cinema,
		Logger: logger,
		Config: config,
	}
}

func (h *CinemaHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit := h.Config.Limit
	cinemas, totalRecords, totalPages, err := h.Cinema.GetAll(r.Context(), page, limit)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data processed successfully.", http.StatusOK, cinemas, &utils.Pagination{
		CurrentPage:  page,
		Limit:        limit,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	})
}

func (h *CinemaHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	cinema, err := h.Cinema.GetByID(r.Context(), id)
	if err != nil {
		utils.WriteError(w, "Data not found", http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w, "Cinema found", http.StatusOK, cinema, nil)
}
