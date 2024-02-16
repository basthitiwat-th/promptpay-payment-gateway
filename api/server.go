package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"promptpay-payment-gateway/constants"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	Echo *echo.Echo
}

func NewServer(timeout time.Duration) *server {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		slog.Error(err.Error())
		c.JSON(code, constants.CodeInternalError)
	}

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: timeout,
	}))

	return &server{e}
}

func (s *server) Start(port string) {
	address := fmt.Sprintf(":%v", port)
	s.Echo.Start(address)
}

func (s *server) Shutdown() {
	slog.Info("==== Server Shutdown ====")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Echo.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("shutdown server: %s", err.Error()))
	}
}
