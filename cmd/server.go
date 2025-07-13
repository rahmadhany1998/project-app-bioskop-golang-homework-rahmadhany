package cmd

import (
	"fmt"
	"net/http"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func ApiServer(config utils.Configuration, logger *zap.Logger, h *chi.Mux) {
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", h); err != nil {
		logger.Fatal("can't run service", zap.Error(err))
	}
}
