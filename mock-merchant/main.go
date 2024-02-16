package main

import (
	"fmt"
	"net/http"

	"promptpay-payment-gateway/constants"

	"github.com/labstack/echo/v4"
)

type ConfirmReq struct {
	ReferenceID string `json:"reference_id"`
}

func main() {
	e := echo.New()
	e.POST("merchant/confirm", func(c echo.Context) error {
		fmt.Println("POST: merchant/confirm")
		var req ConfirmReq
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, constants.CodeBadRequest)
		}
		resp := map[string]string{
			"code":         "0000",
			"description":  "success",
			"reference_id": req.ReferenceID,
		}

		return c.JSON(http.StatusOK, resp)
	})

	e.Logger.Fatal(e.Start(":7070"))
}
