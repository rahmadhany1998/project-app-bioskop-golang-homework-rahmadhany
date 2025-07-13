package main

import (
	"log"
	"project-app-bioskop-golang-homework-rahmadhany/cmd"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/repository"
	"project-app-bioskop-golang-homework-rahmadhany/internal/wire"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/database"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/middleware"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"

	"go.uber.org/zap"
)

func main() {
	// read config
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	// init logger
	logger, err := utils.InitLogger(config.PathLogger, config)
	if err != nil {
		log.Fatal("can't init logger %w", zap.Error(err))
	}

	//Init db
	db, err := database.InitDB(config)
	if err != nil {
		logger.Fatal("can't connect to database ", zap.Error(err))
	}

	repo := repository.NewRepository(db, logger)
	mLogger := middleware.NewLoggerMiddleware(logger)
	router := wire.Wiring(repo, mLogger, logger, config)

	cmd.ApiServer(config, logger, router)
}
