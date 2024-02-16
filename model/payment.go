package model

type PaymentValidateReq struct {
	AccountNO   string  `json:"account_no" validate:"max=12"`
	MerchantID  int     `json:"merchant_id" validate:"gt=0"`
	PromptPayID string  `json:"prompt_pay_id" validate:"max=12"`
	Amount      float64 `json:"amount" validate:"gt=0"`
}

type PaymentValidateResp struct {
	ReferenceID     string  `json:"reference_id"`
	AccountNo       string  `json:"account_no"`
	PromptpayID     string  `json:"prompt_pay_id"`
	FromAccountName string  `json:"from_account_name"`
	ToAccountName   string  `json:"to_account_name"`
	Amount          float64 `json:"amount"`
}

type PaymentConfirmReq struct {
	ReferenceID string `json:"reference_id"`
}

type PaymentConfirmResp struct {
	AccountNO   string  `json:"account_no"`
	MerchantID  string  `json:"merchant_id"`
	PromptPayID string  `json:"prompt_pay_id"`
	Amount      float64 `json:"amount"`
}

type PromptPayValidateReq struct {
	AccountNO   string  `json:"account_no"`
	PromptPayID string  `json:"prompt_pay_id"`
	Amount      float64 `json:"amount"`
}

type PromptPayValidateResp struct {
	Code        string                    `json:"code"`
	Description string                    `json:"description"`
	Data        PromptPayValidateDataResp `json:"data"`
}

type PromptPayValidateDataResp struct {
	AccountNo       string  `json:"account_no"`
	PromptpayID     string  `json:"prompt_pay_id"`
	FromAccountName string  `json:"from_account_name"`
	ToAccountName   string  `json:"to_account_name"`
	ReferenceID     string  `json:"reference_id"`
	Amount          float64 `json:"amount"`
}

type PromptPayConfirmReq struct {
	ReferenceID string  `json:"reference_id"`
	AccountNO   string  `json:"account_no"`
	PromptPayID string  `json:"prompt_pay_id"`
	Amount      float64 `json:"amount"`
}

type MerchantConfirmReq struct {
	ReferenceID string `json:"reference_id"`
}

type PromptPayConfirmResp struct {
	Code        string                   `json:"code"`
	Description string                   `json:"description"`
	Data        PromptPayConfirmDataResp `json:"data"`
}

type PromptPayConfirmDataResp struct {
	AccountNo       string  `json:"account_no"`
	PromptpayID     string  `json:"prompt_pay_id"`
	FromAccountName string  `json:"from_account_name"`
	ToAccountName   string  `json:"to_account_name"`
	Amount          float64 `json:"amount"`
}
