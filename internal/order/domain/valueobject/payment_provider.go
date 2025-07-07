package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

type PaymentProvider string

var providerRe = regexp.MustCompile(`^[A-Z0-9]+$`)

func NewPaymentProvider(val string) (PaymentProvider, error) {
	val = strings.ToUpper(val)
	if val == "" || !providerRe.MatchString(val) {
		return "", errors.New("invalid payment provider")
	}
	return PaymentProvider(val), nil
}
