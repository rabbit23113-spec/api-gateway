package main

import (
	"log"
	"main/internal/adapters"
	logger "main/internal/adapters"

	"go.uber.org/zap"
)

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Error when initializing a logger: %v", err)
	}
	defer zapLogger.Sync()
	lg := new(logger.ZapLogger)
	logger := lg.NewZapLogger(zapLogger)

	svc := new(adapters.Service)
	handler := adapters.Handler{
		Logger:  logger,
		Service: svc,
	}

	srv := new(adapters.Server)
	if err := srv.Serve("11030", handler.InitRoutes()); err != nil {
		logger.Error(err.Error())
	}
}
