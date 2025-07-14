package wire

import (
	"project-app-bioskop-golang-homework-rahmadhany/internal/adaptor"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/repository"
	"project-app-bioskop-golang-homework-rahmadhany/internal/usecase"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/middleware"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Wiring(repo repository.Repository, mLogger middleware.LoggerMiddleware, logger *zap.Logger, config utils.Configuration) *chi.Mux {
	router := chi.NewRouter()
	router.Use(mLogger.LoggingMiddleware)
	rV1 := chi.NewRouter()
	wireUser(rV1, repo, logger, config)
	wireCinema(rV1, repo, logger, config)
	wireBooking(rV1, repo, logger, config)
	wirePayment(rV1, repo, logger, config)
	router.Mount("/api/v1", rV1)

	return router
}

func wireUser(router *chi.Mux, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecaseUser := usecase.NewUserService(repo, logger, config)
	adaptorUser := adaptor.NewUserHandler(usecaseUser, logger, config)
	router.Post("/register", adaptorUser.Register)
	router.Post("/login", adaptorUser.Login)
	router.Post("/logout", adaptorUser.Logout)
}

func wireCinema(router *chi.Mux, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecaseCinema := usecase.NewCinemaService(repo, logger, config)
	adaptorCinema := adaptor.NewCinemaHandler(usecaseCinema, logger, config)
	router.Get("/cinemas", adaptorCinema.GetAll)
	router.Get("/cinemas/{id}", adaptorCinema.GetByID)
	router.Get("/cinemas/{id}/seats", adaptorCinema.GetSeatAvailability)

}

func wireBooking(router *chi.Mux, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecaseBooking := usecase.NewBookingService(repo, logger, config)
	adaptorBooking := adaptor.NewBookingHandler(usecaseBooking, logger, config)

	router.Group(func(protected chi.Router) {
		protected.Use(middleware.AuthMiddlewareWithRepo(repo)) // <-- gunakan middleware autentikasi di sini
		protected.Post("/booking", adaptorBooking.CreateBooking)
		protected.Get("/history", adaptorBooking.GetBookingHistory)
	})
}

func wirePayment(router *chi.Mux, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecasePayment := usecase.NewPaymentService(repo, logger, config)
	adaptorPayment := adaptor.NewPaymentHandler(usecasePayment, logger, config)

	router.Get("/payment-methods", adaptorPayment.GetPaymentMethods)
	router.Get("/pay", adaptorPayment.Pay)
}
