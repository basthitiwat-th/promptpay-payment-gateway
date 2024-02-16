package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"promptpay-payment-gateway/configs"
	"promptpay-payment-gateway/model"
)

type MerchantClient struct {
	client *http.Client
	cfg    configs.MerchantClient
}

func NewMerchantClient(cfg configs.MerchantClient) *MerchantClient {
	tr := &http.Transport{
		MaxIdleConns: cfg.MaxConns,
	}
	merchantClient := &http.Client{
		Transport: tr,
		Timeout:   cfg.Timeout,
	}

	return &MerchantClient{client: merchantClient, cfg: cfg}
}

func (c *MerchantClient) sendRequest(ctx context.Context, method, url string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return c.client.Do(req)
}

func (c *MerchantClient) ConfirmPaymentMerchant(ctx context.Context, req *model.MerchantConfirmReq) error {
	url := c.cfg.BaseURL + c.cfg.URLConfirm
	httpResp, err := c.sendRequest(ctx, http.MethodPost, url, req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Status: %s", httpResp.Status)
	}

	return nil
}
