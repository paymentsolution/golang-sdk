package paymentsdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"testing"
)

var (
	p2pTransactionRequest = P2PTransactionRequest{
		UserUUID:         userUUID,
		MerchantID:       getRandMerchantID(),
		Amount:           2000,
		CallbackURL:      "https://example.com/callback",
		Currency:         RUB,
		PayeerIdentifier: "mock_identifier",
		PayeerIP:         "127.0.0.1",
		PayeerType:       "trust",
		PaymentMethod:    Card,
	}
)

// TestCreateP2PTransaction tests the CreateP2PTransaction method
func TestP2PTransaction(t *testing.T) {
	sdkBuilder := NewPaymentApiBuilder().
		ApiURL(apiUrl).
		Secret(apiSecret)

	sdk, err := sdkBuilder.Build()
	if err != nil {
		t.Fatalf("Error building SDK: %v", err)
	}

	//создание транзакции
	create, err := sdk.P2P.CreateP2PTransaction(context.Background(), p2pTransactionRequest)
	if err != nil {
		t.Errorf("Error creating p2p transaction: %v", err)
	}
	if create.ResultCode != "ok" {
		t.Errorf("Error creating P2P transaction: %v", err)
	}

	//получение транзакции по id
	get, err := sdk.P2P.GetP2PTransaction(context.Background(), create.Payload.ID)
	if err != nil {
		t.Errorf("Error get p2p transaction: %v, by ID: %s", err, create.Payload.ID)
	}
	if get.ResultCode != "ok" {
		t.Errorf("Error creating P2P transaction: %v", err)
	}
}

// TestCreateP2PDispute tests the CreateP2PDispute method
func TestCreateP2PDispute(t *testing.T) {
	sdkBuilder := NewPaymentApiBuilder().
		ApiURL(apiUrl).
		Secret(apiSecret)

	sdk, err := sdkBuilder.Build()
	if err != nil {
		t.Fatalf("Error building SDK: %v", err)
	}

	//создание транзакции
	create, err := sdk.P2P.CreateP2PTransaction(context.Background(), p2pTransactionRequest)
	if err != nil {
		t.Errorf("Error creating p2p transaction: %v", err)
	}
	if create.ResultCode != "ok" {
		t.Errorf("Error creating P2P transaction: %v", err)
	}

	// Mock file for upload (PDF or JPG format)
	fileContent := []byte("mock file content")
	_ = bytes.NewReader(fileContent)

	// Создаём временный файл без расширения
	tmpfile, err := ioutil.TempFile("", "example-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Переименовываем файл с необходимым расширением (PDF в этом случае)
	newFileName := tmpfile.Name() + ".pdf"
	err = os.Rename(tmpfile.Name(), newFileName)
	if err != nil {
		t.Fatal(err)
	}

	// Открываем файл с новым именем для дальнейших операций
	tmpfile, err = os.OpenFile(newFileName, os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(newFileName) // clean up after test

	// Writing content to the file
	if _, err := tmpfile.Write(fileContent); err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	mockP2PDisputeRequest := NewP2PDisputeRequest(create.Payload.ID, 1000, "first-file.pdf", tmpfile)
	mockP2PDisputeRequest.WithProofImage2("second-file.pdf", tmpfile)

	resp, err := sdk.P2P.CreateP2PDispute(context.Background(), mockP2PDisputeRequest)
	if err != nil {
		t.Errorf("CreateP2PDispute() error = %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("CreateP2PDispute() Status = %v, want %v", resp.Status, "ok")
	}
}
