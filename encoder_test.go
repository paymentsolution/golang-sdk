package paymentsdk

import (
	"testing"
)

// TestCalculateSignature tests the calculateSignature function
func TestCalculateSignature(t *testing.T) {
	secret := "mock_api_secret"
	body := []byte(`{"user_uuid":"364dbfc8-ae50-492f-bdd9-748edd84d5c9","amount":300,"callback_url":"https://example.com/callback"}`)

	expectedSignature := "1d41b723b630e0cd790e553b12293995f24a1dd8"

	e := NewEncoder(secret)
	signature := e.CalculateSignature(body)
	if signature != expectedSignature {
		t.Errorf("calculateSignature() = %v, want %v", signature, expectedSignature)
	}
}

// TestVerifySignature tests the verifySignature function
func TestVerifySignature(t *testing.T) {
	secret := "mock_api_secret"
	body := []byte(`{"user_uuid":"364dbfc8-ae50-492f-bdd9-748edd84d5c9","amount":300,"callback_url":"https://example.com/callback"}`)
	e := NewEncoder(secret)
	expectedSignature := e.CalculateSignature(body)

	if !e.VerifySignature(body, expectedSignature) {
		t.Errorf("verifySignature() = false, want true")
	}

	if e.VerifySignature(body, "invalid_signature") {
		t.Errorf("verifySignature() = true, want false")
	}
}
