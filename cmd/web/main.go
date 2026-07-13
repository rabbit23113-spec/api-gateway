package main

import (
	"log"
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
}
