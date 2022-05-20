package api

import (
	"context"

	"github.com/labstack/echo"

	"github.com/nelsonlpco/classic_cc_problens/internal/api/handler"
	"github.com/nelsonlpco/classic_cc_problens/internal/infra/metrics"
)

func Bootstrap() {
	metrics.InitMetrics(context.Background())

	server := NewServer()

	server.WithEcho(echo.New())
	server.WithHandler(handler.NewHandler())

	server.Start()
}
