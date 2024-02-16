package service

import (
	"context"
	"errors"
	"time"

	"promptpay-payment-gateway/client"
	"promptpay-payment-gateway/configs"
	"promptpay-payment-gateway/constants"
	"promptpay-payment-gateway/model"
	"promptpay-payment-gateway/repository"
)

type PaymentStorage interface {
	InsertOne(transaction repository.PaymentHistoryEntity) error
	FindOne(reference_id string) (*repository.PaymentHistoryEntity, error)
	UpdateStatus(reference_id string) error
}

type Payment struct {
	cfg             configs.Config
	repo            PaymentStorage
	promptpayClient client.PromptPayClient
	merchantClient  client.MerchantClient
}

func NewPaymentService(cfg configs.Config, repo PaymentStorage, promptpayClient client.PromptPayClient, merchantClient client.MerchantClient) *Payment {
	return &Payment{cfg, repo, promptpayClient, merchantClient}
}

func (p *Payment) ValidatePayment(paymentReq *model.PaymentValidateReq) (*model.PaymentValidateResp, error) {
	promptpayReq := mappingPaymentValidateToPrompPay(paymentReq)
	resp, err := p.promptpayClient.ValidatePaymentPromptPay(context.Background(), &promptpayReq)
	if err != nil {
		return nil, err
	}

	paymentHistory := repository.PaymentHistoryEntity{
		ReferenceID:   resp.Data.ReferenceID,
		CustomerAccNo: resp.Data.AccountNo,
		PromptPayID:   resp.Data.PromptpayID,
		MerchantID:    paymentReq.MerchantID,
		Amount:        resp.Data.Amount,
		PaymentType:   constants.PAYMENT_TYPE_PROMPTPAY,
		Status:        constants.STATUS_PENDING,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err = p.repo.InsertOne(paymentHistory)
	if err != nil {
		return nil, err
	}

	return &model.PaymentValidateResp{
		ReferenceID:     resp.Data.ReferenceID,
		AccountNo:       resp.Data.AccountNo,
		PromptpayID:     resp.Data.PromptpayID,
		FromAccountName: resp.Data.FromAccountName,
		ToAccountName:   resp.Data.ToAccountName,
		Amount:          resp.Data.Amount,
	}, nil
}

func (p *Payment) ConfirmPayment(paymentReq *model.PaymentConfirmReq) error {
	// Find Ref id in DB
	paymentData, err := p.repo.FindOne(paymentReq.ReferenceID)
	if err != nil {
		return errors.New("FindOne Error")
	}

	// Send To merchant
	err = p.merchantClient.ConfirmPaymentMerchant(context.Background(), &model.MerchantConfirmReq{ReferenceID: paymentReq.ReferenceID})
	if err != nil {
		return errors.New("ConfirmPaymentMerchant Error")
	}

	promptPayReq := mappingPaymentConfirmToPromPay(paymentData)

	// Send To merchant
	resp, err := p.promptpayClient.ConfirmPaymentPromptPay(context.Background(), &promptPayReq)
	if err != nil {
		return errors.New("ConfirmPaymentPromptPay Error")
	}

	if resp.Code != "0000" {
		return errors.New("response code not 0000")
	}

	err = p.repo.UpdateStatus(promptPayReq.ReferenceID)
	if err != nil {
		return errors.New("UpdateStatus Error")
	}

	return nil
}

func mappingPaymentValidateToPrompPay(paymentReq *model.PaymentValidateReq) model.PromptPayValidateReq {
	return model.PromptPayValidateReq{
		AccountNO:   paymentReq.AccountNO,
		PromptPayID: paymentReq.PromptPayID,
		Amount:      paymentReq.Amount,
	}
}

func mappingPaymentConfirmToPromPay(paymentReq *repository.PaymentHistoryEntity) model.PromptPayConfirmReq {
	return model.PromptPayConfirmReq{
		ReferenceID: paymentReq.ReferenceID,
		AccountNO:   paymentReq.CustomerAccNo,
		PromptPayID: paymentReq.PromptPayID,
		Amount:      paymentReq.Amount,
	}
}
