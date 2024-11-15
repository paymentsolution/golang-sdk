package paymentsdk

import (
	"fmt"
	"net/http"
	"time"
)

// PaymentApi представляет структуру SDK с заголовками и URL.
type PaymentApi struct {
	apiURL          string
	secret          string
	logger          Logger
	client          *http.Client
	P2P             *P2P
	MassTransaction *MassTransaction
	Encoder         *Encoder
}

// PaymentApiBuilder помогает построить экземпляр PaymentApi.
type PaymentApiBuilder struct {
	apiURL string
	secret string
	client *http.Client
	logger Logger
}

// NewPaymentApiBuilder создает новый экземпляр PaymentApiBuilder.
func NewPaymentApiBuilder() *PaymentApiBuilder {
	return &PaymentApiBuilder{}
}

func (b *PaymentApiBuilder) ApiURL(apiURL string) *PaymentApiBuilder {
	b.apiURL = apiURL
	return b
}

func (b *PaymentApiBuilder) Secret(secret string) *PaymentApiBuilder {
	b.secret = secret
	return b
}

func (b *PaymentApiBuilder) Client(client *http.Client) *PaymentApiBuilder {
	b.client = client
	return b
}

func (b *PaymentApiBuilder) Logger(logger Logger) *PaymentApiBuilder {
	b.logger = logger
	return b
}

// Build строит и возвращает экземпляр PaymentApi.
func (b *PaymentApiBuilder) Build() (*PaymentApi, error) {
	if b.secret == "" {
		return nil, fmt.Errorf("secret is required")
	}

	if b.apiURL == "" {
		return nil, fmt.Errorf("api URL is required")
	}

	var err error
	if b.logger == nil {
		//по дефолту логгер активен в режиме info
		b.logger, err = NewLogger(true, "info")
		if err != nil {
			return nil, fmt.Errorf("cant build default logger: %s", err.Error())
		}
	}

	if b.client == nil {
		//по дефолту создается кастомный клиент с логом
		b.client = &http.Client{
			Transport: NewLoggingRoundTripper(b.logger, http.DefaultTransport),
			Timeout:   30 * time.Second,
		}
	}

	encoder := NewEncoder(b.secret)

	return &PaymentApi{
		apiURL:          b.apiURL,
		secret:          b.secret,
		client:          b.client,
		logger:          b.logger,
		Encoder:         encoder,
		P2P:             p2pNew(b.apiURL, encoder, b.client),
		MassTransaction: massTransactionNew(b.apiURL, encoder, b.client),
	}, nil
}
