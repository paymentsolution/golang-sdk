package paymentsdk

import (
	"context"
	"testing"
)

// TestCreateP2PTransaction tests the CreateP2PTransaction method
func TestCreateMassTransaction(t *testing.T) {
	sdkBuilder := NewPaymentApiBuilder().
		ApiURL(apiUrl).
		Secret(apiSecret)

	sdk, err := sdkBuilder.Build()
	if err != nil {
		t.Fatalf("Error building SDK: %v", err)
	}

	massTransactionRequest := MassTransactionRequest{
		UserUUID:      userUUID,
		MerchantID:    getRandMerchantID(),
		Amount:        2000,
		CallbackURL:   "https://example.com/callback",
		ToCard:        "4111111111111111",
		Currency:      RUB,
		PaymentMethod: Card,
	}

	//создание транзакции
	create, err := sdk.MassTransaction.CreateMassTransaction(context.Background(), massTransactionRequest)
	if err != nil {
		t.Fatalf("Error creating mass transaction: %v", err)
	}
	if create.ResultCode != "ok" {
		t.Fatalf("Error creating mass transaction: %v", err)
	}

	//получение транзакции по id
	get, err := sdk.MassTransaction.GetMassTransaction(context.Background(), create.Payload.ID)
	if err != nil {
		t.Fatalf("Error get p2p transaction: %v, by ID: %s", err, create.Payload.ID)
	}
	if get.ResultCode != "ok" {
		t.Fatalf("Error creating P2P transaction: %v", err)
	}
}
