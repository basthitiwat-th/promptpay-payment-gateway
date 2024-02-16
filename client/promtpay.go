package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"promptpay-payment-gateway/configs"
	"promptpay-payment-gateway/model"
)

type PromptPayClient struct {
	client *http.Client
	cfg    configs.PromptPayClient
}

func NewPromptPayClient(cfg configs.PromptPayClient) *PromptPayClient {
	tr := &http.Transport{
		MaxIdleConns: cfg.MaxConns,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   cfg.Timeout,
	}

	return &PromptPayClient{client: client, cfg: cfg}
}

func (c *PromptPayClient) sendRequest(ctx context.Context, method, url string, body interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	httpResp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Status: %s", httpResp.Status)
	}

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTTP response body: %v", err)
	}

	return respBody, nil
}

func (c *PromptPayClient) ValidatePaymentPromptPay(ctx context.Context, req *model.PromptPayValidateReq) (*model.PromptPayValidateResp, error) {
	url := c.cfg.BaseURL + c.cfg.URLValidate
	respBody, err := c.sendRequest(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	resp := &model.PromptPayValidateResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return resp, nil
}

func (c *PromptPayClient) ConfirmPaymentPromptPay(ctx context.Context, req *model.PromptPayConfirmReq) (*model.PromptPayConfirmResp, error) {
	url := c.cfg.BaseURL + c.cfg.URLConfirm
	respBody, err := c.sendRequest(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	resp := &model.PromptPayConfirmResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return resp, nil
}
