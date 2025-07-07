package valueobject

import "errors"

type PaymentID string

func NewPaymentID(val string) (PaymentID, error) {
	if len(val) < 8 {
		return "", errors.New("payment id too short")
	}
	return PaymentID(val), nil
}
