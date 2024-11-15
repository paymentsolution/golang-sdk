package paymentsdk

import (
	"crypto/sha1"
	"encoding/hex"
)

type Encoder struct {
	secret string
}

func NewEncoder(secret string) *Encoder {
	return &Encoder{secret: secret}
}

// calculateSignature вычисляет SHA1 подпись для тела запроса и api_secret.
func (e *Encoder) CalculateSignature(body []byte) string {
	hash := sha1.New()
	hash.Write([]byte(e.secret + string(body)))
	return hex.EncodeToString(hash.Sum(nil))
}

// verifySignature проверяет подпись, вычисленную для тела ответа и api_secret.
func (e *Encoder) VerifySignature(body []byte, receivedSignature string) bool {
	expectedSignature := e.CalculateSignature(body)
	return expectedSignature == receivedSignature
}
