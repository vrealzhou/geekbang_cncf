package main

import (
	"context"
	"os"
	"strconv"

	"github.com/vrealzhou/geekbang_cncf/module_08/api"
	"github.com/vrealzhou/geekbang_cncf/module_08/internal/service"
	"github.com/vrealzhou/geekbang_cncf/module_08/log"
)

func main() {
	logger := log.NewLogger()
	defer logger.Shutdown()
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		logger.Fatal(err)
	}
	ctx, m := service.NewServiceManager(context.Background())
	m.Start(ctx, service.NewSignalService(logger))
	m.Start(ctx, service.NewHTTPService(port, api.NewHandler(os.Getenv("VERSION"), logger), logger))
	err = m.Idle()
	if err != nil {
		if service.NormalShutdown(err) {
			logger.Info("Program shutdown normally")
		} else {
			logger.Fatal(err)
		}
	}
}
