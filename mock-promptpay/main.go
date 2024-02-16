package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type User struct {
	AccountNo   string  `json:"account_no"`
	PromptpayID string  `json:"prompt_pay_id"`
	Amount      float64 `json:"amount"`
	AccountName string  `json:"account_name"`
}

var UsersMock = []User{
	{AccountNo: "0588559424", PromptpayID: "0814592555", AccountName: "Mister A", Amount: 500},
	{AccountNo: "2456412355", PromptpayID: "1256313412", AccountName: "Mister B", Amount: 100000},
	{AccountNo: "4501405162", PromptpayID: "0561456778", AccountName: "ร้านเติมเกมส์", Amount: 0},
	{AccountNo: "6123124341", PromptpayID: "1251212342", AccountName: "เต่าบิน", Amount: 0},
}

type ValidateResp struct {
	AccountNo       string  `json:"account_no"`
	PromptpayID     string  `json:"prompt_pay_id"`
	FromAccountName string  `json:"from_account_name"`
	ToAccountName   string  `json:"to_account_name"`
	Amount          float64 `json:"amount"`
	ReferenceID     string  `json:"reference_id"`
}

type ConfirmResp struct {
	ReferenceID string `json:"reference_id"`
}

type PaymentValidateReq struct {
	AccountNO   string  `json:"account_no"`
	PromptPayID string  `json:"prompt_pay_id"`
	Amount      float64 `json:"amount"`
}

type PaymentConfirmReq struct {
	AccountNO   string  `json:"account_no"`
	PromptPayID string  `json:"prompt_pay_id"`
	ReferenceID string  `json:"reference_id"`
	Amount      float64 `json:"amount"`
}

func main() {
	e := echo.New()
	e.POST("promptpay/validate", func(c echo.Context) error {
		fmt.Println("POST: promptpay/validate")
		var req PaymentValidateReq
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}

		accountFrom := findAccFromByAccountNo(req.AccountNO)
		accountTo := findAccToByPromptpayID(req.PromptPayID)

		if accountFrom == nil || accountTo == nil {
			return c.JSON(http.StatusNotFound, "Not found")
		}

		mapResponse := map[string]interface{}{}
		mapResponse["data"] = ValidateResp{
			AccountNo:       accountFrom.AccountNo,
			PromptpayID:     accountTo.PromptpayID,
			FromAccountName: accountFrom.AccountName,
			ToAccountName:   accountTo.AccountName,
			ReferenceID:     uuid.NewString(),
			Amount:          req.Amount,
		}
		mapResponse["code"] = "0000"
		mapResponse["description"] = "success"

		return c.JSON(http.StatusOK, mapResponse)
	})

	e.POST("promptpay/confirm", func(c echo.Context) error {
		fmt.Println("POST: promptpay/confirm")
		var req PaymentConfirmReq
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}

		accountFrom := findAccFromByAccountNo(req.AccountNO)
		accountTo := findAccToByPromptpayID(req.PromptPayID)

		if accountFrom == nil || accountTo == nil {
			return c.JSON(http.StatusNotFound, "Not found")
		}

		mapResponse := map[string]interface{}{}
		mapResponse["data"] = ConfirmResp{
			ReferenceID: req.ReferenceID,
		}
		mapResponse["code"] = "0000"
		mapResponse["description"] = "payment success"

		fmt.Println(mapResponse)
		return c.JSON(http.StatusOK, mapResponse)
	})

	e.Logger.Fatal(e.Start(":9090"))
}

func findAccFromByAccountNo(accountNo string) *User {
	for _, user := range UsersMock {
		if accountNo == user.AccountNo {
			return &user
		}
	}

	return nil
}

func findAccToByPromptpayID(PromptpayID string) *User {
	for _, user := range UsersMock {
		if PromptpayID == user.PromptpayID {
			return &user
		}
	}

	return nil
}
