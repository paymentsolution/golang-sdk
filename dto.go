package paymentsdk

import (
	"mime/multipart"
)

// P2PTransactionRequest представляет тело запроса для создания P2P транзакции.
type P2PTransactionRequest struct {
	UserUUID         string            `json:"user_uuid"`
	MerchantID       string            `json:"merchant_id"`
	PayeerIdentifier string            `json:"payeer_identifier"`
	PayeerIP         string            `json:"payeer_ip"`
	PayeerType       string            `json:"payeer_type"`
	Currency         CurrencyEnum      `json:"currency"`
	PaymentMethod    PaymentMethodEnum `json:"payment_method"`
	Amount           int               `json:"amount"`
	CallbackURL      string            `json:"callback_url"`
	//not required
	RedirectURL      *string `json:"redirect_url"`
	Email            *string `json:"email"`
	CustomerName     *string `json:"customer_name"`
	PayeerCardNumber *string `json:"payeer_card_number"`
}

// NewP2PTransactionRequest создает новый экземпляр P2PTransactionRequest с обязательными параметрами.
func NewP2PTransactionRequest(userUUID, merchantID, payeerIdentifier, payeerIP, payeerType, callbackURL string, currency CurrencyEnum, paymentMethod PaymentMethodEnum, amount int) *P2PTransactionRequest {
	return &P2PTransactionRequest{
		UserUUID:         userUUID,
		MerchantID:       merchantID,
		PayeerIdentifier: payeerIdentifier,
		PayeerIP:         payeerIP,
		PayeerType:       payeerType,
		Currency:         currency,
		PaymentMethod:    paymentMethod,
		Amount:           amount,
		CallbackURL:      callbackURL,
	}
}

// WithEmail задает email и возвращает обновленный запрос
func (p *P2PTransactionRequest) WithEmail(email string) *P2PTransactionRequest {
	p.Email = &email
	return p
}

// WithCustomerName задает имя клиента и возвращает обновленный запрос
func (p *P2PTransactionRequest) WithCustomerName(customerName string) *P2PTransactionRequest {
	p.CustomerName = &customerName
	return p
}

// WithPayeerCardNumber задает номер карты и возвращает обновленный запрос
func (p *P2PTransactionRequest) WithPayeerCardNumber(payeerCardNumber string) *P2PTransactionRequest {
	p.PayeerCardNumber = &payeerCardNumber
	return p
}

// WithPayeerCardNumber задает номер карты и возвращает обновленный запрос
func (p *P2PTransactionRequest) WithRedirectURL(redirectURL string) *P2PTransactionRequest {
	p.RedirectURL = &redirectURL
	return p
}

// P2PTransactionResponse представляет тело ответа API для создания P2P транзакции.
type P2PTransactionResponse struct {
	ResultCode string `json:"result_code"`
	Payload    struct {
		ID                string               `json:"id"`
		MerchantID        string               `json:"merchant_id"`
		Currency          CurrencyEnum         `json:"currency"`
		FormURL           string               `json:"form_url"`
		State             TransactionStateEnum `json:"state"`
		CreatedAt         string               `json:"created_at"`
		UpdatedAt         string               `json:"updated_at"`
		CloseAt           string               `json:"close_at"`
		CallbackURL       string               `json:"callback_url"`
		RedirectURL       string               `json:"redirect_url"`
		Email             string               `json:"email"`
		CustomerName      string               `json:"customer_name"`
		Rate              string               `json:"rate"`
		Amount            string               `json:"amount"`
		FiatAmount        string               `json:"fiat_amount"`
		OldFiatAmount     string               `json:"old_fiat_amount"`
		ServiceCommission string               `json:"service_commission"`
		TotalAmount       string               `json:"total_amount"`
		PaymentMethod     PaymentMethodEnum    `json:"payment_method"`
		RecipientCard     struct {
			ID            string   `json:"id"`
			Number        string   `json:"number"`
			BankName      string   `json:"bank_name"`
			BankFullName  string   `json:"bank_full_name"`
			BankColors    struct{} `json:"bank_colors"`
			Brand         string   `json:"brand"`
			CardHolder    string   `json:"card_holder"`
			PaymentMethod string   `json:"payment_method"`
			UpdatedAt     string   `json:"updated_at"`
			CreatedAt     string   `json:"created_at"`
			SberpayURL    string   `json:"sberpay_url"`
		} `json:"resipient_card"`
	} `json:"payload"`
}

// P2PDisputeRequest представляет тело запроса для создания диспута по p2p транзакции.
type sdkFile struct {
	file multipart.File
	Name string
}

type P2PDisputeRequest struct {
	TransactionID string
	Amount        int
	ProofImage    sdkFile
	//not required
	ProofImage2 *sdkFile
}

func NewP2PDisputeRequest(TransactionID string, Amount int, fileNameWithFormat string, file multipart.File) *P2PDisputeRequest {
	return &P2PDisputeRequest{TransactionID: TransactionID, Amount: Amount, ProofImage: sdkFile{Name: fileNameWithFormat, file: file}}
}

func (p *P2PDisputeRequest) WithProofImage2(fileNameWithFormat string, file multipart.File) *P2PDisputeRequest {
	p.ProofImage2 = &sdkFile{file: file, Name: fileNameWithFormat}
	return p
}

// P2PDisputeResponse представляет тело ответа API для создания диспута по p2p транзакции.
type P2PDisputeResponse struct {
	Data struct {
		ID          int    `json:"id"`
		State       string `json:"state"`
		Repeated    bool   `json:"repeated"`
		UpdatedAt   string `json:"updated_at"`
		CreatedAt   string `json:"created_at"`
		ProofImage  string `json:"proof_image"`
		ProofImage2 string `json:"proof_image2"`
		Amount      int    `json:"amount"`
		P2PTx       struct {
			ID               string `json:"id"`
			MerchantID       string `json:"merchant_id"`
			Currency         string `json:"currency"`
			ToCurrency       string `json:"to_currency"`
			State            string `json:"state"`
			CreatedAt        string `json:"created_at"`
			UpdatedAt        string `json:"updated_at"`
			CloseAt          string `json:"close_at"`
			RedirectURL      string `json:"redirect_url"`
			Email            string `json:"email"`
			CustomerName     string `json:"customer_name"`
			Rate             string `json:"rate"`
			Amount           string `json:"amount"`
			FiatAmount       string `json:"fiat_amount"`
			OldFiatAmount    string `json:"old_fiat_amount"`
			PaymentMethod    string `json:"payment_method"`
			PayeerBankName   string `json:"payeer_bank_name"`
			Comment          string `json:"comment"`
			AntifraudVerdict string `json:"antifraud_verdict"`
			Requisities      struct {
				Number        string                 `json:"number"`
				CardHolder    string                 `json:"card_holder"`
				BankName      string                 `json:"bank_name"`
				BankFullName  string                 `json:"bank_full_name"`
				BankColors    map[string]interface{} `json:"bank_colors"`
				Brand         string                 `json:"brand"`
				PaymentMethod string                 `json:"payment_method"`
				UpdatedAt     string                 `json:"updated_at"`
				CreatedAt     string                 `json:"created_at"`
				ID            string                 `json:"id"`
				SberpayURL    string                 `json:"sberpay_url"`
			} `json:"requisities"`
		} `json:"p2p_tx"`
	} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Errors  interface{} `json:"errors"`
	Meta    interface{} `json:"meta"`
}

// MassTransactionRequest представляет тело запроса для создания массовой транзакции.
type MassTransactionRequest struct {
	UserUUID      string            `json:"user_uuid"`
	MerchantID    string            `json:"merchant_id"`
	Amount        int               `json:"amount"`
	CallbackURL   string            `json:"callback_url"`
	ToCard        string            `json:"to_card"`
	Currency      CurrencyEnum      `json:"currency"`
	PaymentMethod PaymentMethodEnum `json:"payment_method"`

	SbpBankName        *string `json:"sbp_bank_name,omitempty"`
	BankName           *string `json:"bank_name,omitempty"`
	RecipientFirstName *string `json:"recipient_first_name,omitempty"`
	RecipientLastName  *string `json:"recipient_last_name,omitempty"`
}

// NewMassTransactionRequest создает новый экземпляр MassTransactionRequest с обязательными параметрами.
func NewMassTransactionRequest(userUUID, merchantID, toCard, callbackURL string, amount int, currency CurrencyEnum, paymentMethod PaymentMethodEnum) *MassTransactionRequest {
	return &MassTransactionRequest{
		UserUUID:      userUUID,
		MerchantID:    merchantID,
		ToCard:        toCard,
		Amount:        amount,
		CallbackURL:   callbackURL,
		Currency:      currency,
		PaymentMethod: paymentMethod,
	}
}

func (m *MassTransactionRequest) WithSbpBankName(sbpBankName string) *MassTransactionRequest {
	m.SbpBankName = &sbpBankName
	return m
}

func (m *MassTransactionRequest) WithBankName(bankName string) *MassTransactionRequest {
	m.BankName = &bankName
	return m
}

func (m *MassTransactionRequest) WithRecipientFirstName(recipientFirstName string) *MassTransactionRequest {
	m.RecipientFirstName = &recipientFirstName
	return m
}

func (m *MassTransactionRequest) WithRecipientLastName(recipientLastName string) *MassTransactionRequest {
	m.RecipientLastName = &recipientLastName
	return m
}

// MassTransactionResponse представляет тело ответа API для создания массовой транзакции.
type MassTransactionResponse struct {
	ResultCode string `json:"result_code"`
	Payload    struct {
		ID                string            `json:"id"`
		MerchantId        string            `json:"merchant_id"`
		State             string            `json:"state"`
		CreatedAt         string            `json:"created_at"`
		UpdatedAt         string            `json:"updated_at"`
		Currency          string            `json:"currency"`
		CallBackUrl       string            `json:"callback_url"`
		Amount            string            `json:"amount"`
		FiatAmount        string            `json:"fiat_amount"`
		OldFiatAmount     string            `json:"old_fiat_amount"`
		Rate              string            `json:"rate"`
		CommissionType    string            `json:"commission_type"`
		ServiceCommission string            `json:"service_commission"`
		TotalAmount       string            `json:"total_amount"`
		BankName          string            `json:"bank_name"`
		SbpBankName       string            `json:"sbp_bank_name"`
		PaymentMethod     PaymentMethodEnum `json:"payment_method"`
		RecipientCard     string            `json:"recipient_card"`
	} `json:"payload"`
}
