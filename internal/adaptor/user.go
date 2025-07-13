package adaptor

import (
	"encoding/json"
	"net/http"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/entity"
	"project-app-bioskop-golang-homework-rahmadhany/internal/dto"
	"project-app-bioskop-golang-homework-rahmadhany/internal/usecase"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"
	"strings"

	"go.uber.org/zap"
)

type UserHandler struct {
	User   usecase.UserService
	Logger *zap.Logger
	Config utils.Configuration
}

func NewUserHandler(user usecase.UserService, logger *zap.Logger, config utils.Configuration) UserHandler {
	return UserHandler{
		User:   user,
		Logger: logger,
		Config: config,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.Error("error : ", zap.Error(err))
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	validation, err := utils.ValidateData(req)
	if err != nil {
		h.Logger.Error("error : ", zap.Error(err))
		utils.ResponseErrorValidation(w, http.StatusBadRequest, "Validation error", validation)
		return
	}
	if err := h.User.Register(r.Context(), &req); err != nil {
		if err.Error() == "username already exists" {
			utils.WriteError(w, err.Error(), http.StatusConflict)
		} else {
			utils.WriteError(w, "failed to register", http.StatusInternalServerError)
		}
		return
	}
	utils.WriteSuccess(w, "user register successfully", http.StatusCreated, nil, nil)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.Error("error : ", zap.Error(err))
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	validation, err := utils.ValidateData(req)
	if err != nil {
		h.Logger.Error("error : ", zap.Error(err))
		utils.ResponseErrorValidation(w, http.StatusBadRequest, "Validation error", validation)
		return
	}

	user, token, err := h.User.Login(r.Context(), req.Username, req.Password)
	if err != nil {
		h.Logger.Error("error : ", zap.Error(err))
		utils.WriteError(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	utils.WriteSuccess(w, "login success", http.StatusOK, map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
		"token":    token,
	}, nil)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		utils.WriteError(w, "Missing token", http.StatusBadRequest)
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_ = h.User.Logout(r.Context(), token)
	utils.WriteSuccess(w, "User Logout Successfully", http.StatusOK, nil, nil)
}
