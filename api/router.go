package api

import "github.com/labstack/echo/v4"

func Route(e *echo.Echo, h *Handler) {
	e.POST("payment-gateway/validate", h.ValidateTransaction)
	e.POST("payment-gateway/confirm", h.ConfirmTransaction)
}
