package paymentsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
)

type P2P struct {
	apiURL  string
	client  *http.Client
	encoder *Encoder
}

func p2pNew(apiURL string, encoder *Encoder, client *http.Client) *P2P {
	return &P2P{apiURL: apiURL, client: client, encoder: encoder}
}

// CreateP2PTransaction создает платеж p2p и получает ссылку на пополнение.
func (p2p *P2P) CreateP2PTransaction(ctx context.Context, req P2PTransactionRequest) (*P2PTransactionResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	// Создаем HTTP запрос
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p2p.apiURL+"/v1/p2p_transactions", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Устанавливаем заголовки
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set(signatureHeader, p2p.encoder.CalculateSignature(jsonData))

	// Отправляем запрос
	resp, err := p2p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v, reason: %s", resp.StatusCode, string(respBody))
	}

	var response P2PTransactionResponse
	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error Unmarshal response: %v", err)
	}

	return &response, nil
}

// GetP2PTransaction получает информацию о p2p транзакции по её ID.
func (p2p *P2P) GetP2PTransaction(ctx context.Context, transactionID string) (*P2PTransactionResponse, error) {
	url := fmt.Sprintf("%s/v1/p2p_transactions/%s", p2p.apiURL, transactionID)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := p2p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v, reason: %s", resp.StatusCode, string(respBody))
	}

	var response P2PTransactionResponse
	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error Unmarshal response: %v", err)
	}

	return &response, nil
}

// CreateP2PDispute создаем диспут по p2p транзакции.
func (p2p *P2P) CreateP2PDispute(ctx context.Context, req P2PDisputeRequest) (*P2PDisputeResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField(transactionIdForm, req.TransactionID)
	_ = writer.WriteField(p2pDisputeAmountForm, strconv.Itoa(req.Amount))

	part, err := writer.CreateFormFile(p2pDisputeProofImageForm, filepath.Base(req.ProofImage.Name))
	if err != nil {
		return nil, fmt.Errorf("error creating form File with name: %s, err: %v", req.ProofImage.Name, err)
	}
	defer req.ProofImage.file.Close()
	if _, err = io.Copy(part, req.ProofImage.file); err != nil {
		return nil, fmt.Errorf("error copying File: %v", err)
	}
	if req.ProofImage2 != nil {

		part2, err := writer.CreateFormFile(p2pDisputeProofImageForm2, filepath.Base(req.ProofImage2.Name))
		if err != nil {
			return nil, fmt.Errorf("error creating form File2 with name: %s, err: %v", req.ProofImage2.Name, err)
		}
		defer req.ProofImage2.file.Close()
		if _, err = io.Copy(part2, req.ProofImage2.file); err != nil {
			return nil, fmt.Errorf("error copying File2: %v", err)
		}
	}

	if err = writer.Close(); err != nil {
		return nil, fmt.Errorf("error closing writer: %v", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/v1/p2p_disputes/from_client", p2p.apiURL), body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	httpReq.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := p2p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v, reason: %s", resp.StatusCode, string(respBody))
	}
	var response P2PDisputeResponse
	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error Unmarshal response: %v", err)
	}

	return &response, nil
}
