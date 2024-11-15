package paymentsdk

import "fmt"

type CurrencyEnum string
type PaymentMethodEnum string
type TransactionStateEnum string

func CurrencyFrom(val string) (CurrencyEnum, error) {
	switch val {
	case string(RUB):
		return RUB, nil
	case string(UZS):
		return UZS, nil
	case string(KRW):
		return KRW, nil
	default:
		return "", fmt.Errorf("invalid currency value: %s", val)
	}
}

func PaymentMethodFrom(val string) (PaymentMethodEnum, error) {
	switch val {
	case string(Card):
		return Card, nil
	case string(SberPay):
		return SberPay, nil
	case string(Sbp):
		return Sbp, nil
	case string(SbpFast):
		return SbpFast, nil
	case string(AccountNumber):
		return AccountNumber, nil
	default:
		return "", fmt.Errorf("invalid payment method value: %s", val)
	}
}

func TransactionStateFrom(val string) (TransactionStateEnum, error) {
	switch val {
	case string(WaitingPayment):
		return WaitingPayment, nil
	case string(Paid):
		return Paid, nil
	case string(Failed):
		return Failed, nil
	case string(ClosedFailed):
		return ClosedFailed, nil
	case string(RepeatedClosedFailed):
		return RepeatedClosedFailed, nil
	case string(Successed):
		return Successed, nil
	case string(AcceptedSuccessed):
		return AcceptedSuccessed, nil
	case string(RepeatedAcceptedSuccessed):
		return RepeatedAcceptedSuccessed, nil
	case string(Reviewing):
		return Reviewing, nil
	case string(RepeatedReviewing):
		return RepeatedReviewing, nil
	default:
		return "", fmt.Errorf("invalid transaction state value: %s", val)
	}
}

const (
	RUB CurrencyEnum = "rub"
	UZS CurrencyEnum = "uzs"
	KRW CurrencyEnum = "krw"

	Card          PaymentMethodEnum = "card"
	SberPay       PaymentMethodEnum = "sberpay"
	Sbp           PaymentMethodEnum = "sbp"
	SbpFast       PaymentMethodEnum = "sbp_fast"
	AccountNumber PaymentMethodEnum = "account_number"

	WaitingPayment            TransactionStateEnum = "waiting_payment"
	Paid                      TransactionStateEnum = "paid"
	Failed                    TransactionStateEnum = "failed"
	ClosedFailed              TransactionStateEnum = "closed_failed"
	RepeatedClosedFailed      TransactionStateEnum = "repeated_closed_failed"
	Successed                 TransactionStateEnum = "successed"
	AcceptedSuccessed         TransactionStateEnum = "accepted_successed"
	RepeatedAcceptedSuccessed TransactionStateEnum = "repeated_accepted_successed"
	Reviewing                 TransactionStateEnum = "reviewing"
	RepeatedReviewing         TransactionStateEnum = "repeated_reviewing"
)
