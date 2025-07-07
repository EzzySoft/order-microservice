package valueobject

import (
	"errors"
	"strings"
)

type Currency string

const (
	CurrencyRUB Currency = "RUB"
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
)

var validCurrencies = map[Currency]struct{}{
	CurrencyRUB: {},
	CurrencyUSD: {},
	CurrencyEUR: {},
}

func NewCurrency(val string) (Currency, error) {
	val = strings.ToUpper(val)
	c := Currency(val)
	if _, ok := validCurrencies[c]; !ok {
		return "", errors.New("invalid currency")
	}
	return c, nil
}
