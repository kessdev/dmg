package main

import (
	"github.com/dev-kess/dmg/logger"
	"go.uber.org/zap"
)

func main() {
	var log = logger.CreateLogger("./main.go")
	defer log.Sync()
	log.Info("Data Mapper in construction ...", zap.String("method", "main"))
}
