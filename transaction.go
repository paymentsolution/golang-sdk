package paymentsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MassTransaction struct {
	apiURL  string
	client  *http.Client
	encoder *Encoder
}

func massTransactionNew(apiURL string, encoder *Encoder, client *http.Client) *MassTransaction {
	return &MassTransaction{apiURL: apiURL, encoder: encoder, client: client}
}

// CreateMassTransaction создает заявку на выплату на карту.
func (mt *MassTransaction) CreateMassTransaction(ctx context.Context, req MassTransactionRequest) (*MassTransactionResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/v1/mass_transactions", mt.apiURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set(signatureHeader, mt.encoder.CalculateSignature(jsonData))

	resp, err := mt.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v, reason: %s", resp.StatusCode, string(respBody))
	}

	var response MassTransactionResponse
	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error Unmarshal response: %v", err)
	}

	return &response, nil
}

// GetMassTransaction получает информацию о транзакции по её ID.
func (mt *MassTransaction) GetMassTransaction(ctx context.Context, transactionID string) (*MassTransactionResponse, error) {
	url := fmt.Sprintf("%s/v1/mass_transactions/%s", mt.apiURL, transactionID)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := mt.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v, reason: %s", resp.StatusCode, string(respBody))
	}

	var response MassTransactionResponse
	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error Unmarshal response: %v", err)
	}

	return &response, nil
}
