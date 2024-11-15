package paymentsdk

import (
	"math/rand"
	"time"
)

// Mock API secret
const apiUrl = "https://test.url"
const apiSecret = "secret"
const userUUID = "id"

func getRandMerchantID() string {
	length := 5
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(charset))]
	}
	return string(b)
}
