package api

import (
	"net/http"

	"promptpay-payment-gateway/constants"
	"promptpay-payment-gateway/model"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

type PaymentService interface {
	ValidatePayment(paymentReq *model.PaymentValidateReq) (*model.PaymentValidateResp, error)
	ConfirmPayment(paymentReq *model.PaymentConfirmReq) error
}

type Handler struct {
	svc PaymentService
}

func NewHandler(svc PaymentService) *Handler {
	return &Handler{svc}
}

func (h *Handler) ValidateTransaction(c echo.Context) error {
	var req model.PaymentValidateReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":        constants.CodeBadRequest,
			"description": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":        constants.CodeBadRequest,
			"description": err.Error(),
		})
	}

	resp, err := h.svc.ValidatePayment(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":        constants.CodeInternalError,
			"description": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        constants.CodeSuccess,
		"description": "success",
		"data":        resp,
	})
}

func (h *Handler) ConfirmTransaction(c echo.Context) error {
	var req model.PaymentConfirmReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":        constants.CodeBadRequest,
			"description": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":        constants.CodeBadRequest,
			"description": err.Error(),
		})
	}

	if err := h.svc.ConfirmPayment(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":        constants.CodeInternalError,
			"description": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        constants.CodeSuccess,
		"description": "success",
	})
}
