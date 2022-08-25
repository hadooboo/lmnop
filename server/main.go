package main

import (
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gin_server "jaehonam.com/lmnop/adapter/in/gin-server"
	solved_ac_client "jaehonam.com/lmnop/adapter/out/solved-ac-client"
	"jaehonam.com/lmnop/application"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	logger := initLogger()

	restyAdapter := initRestyAdapter()
	service := initService(restyAdapter)
	server := initAPIServer(service)

	runServer(server)

	clearLogger(logger)
}

func initLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.Level.SetLevel(zap.DebugLevel)
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	config.DisableCaller = true
	config.DisableStacktrace = true

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	_ = zap.ReplaceGlobals(logger)

	return logger
}

func initRestyAdapter() *solved_ac_client.RestyAdapter {
	return solved_ac_client.NewRestyAdapter()
}

func initService(restyAdapter *solved_ac_client.RestyAdapter) *application.Service {
	return application.NewService(restyAdapter)
}

func initAPIServer(service *application.Service) *gin_server.GinAPIServer {
	return gin_server.NewGinAPIServer(service)
}

func runServer(server *gin_server.GinAPIServer) {
	zap.S().Fatal(server.Serve())
}

func clearLogger(logger *zap.Logger) {
	_ = logger.Sync()
}
