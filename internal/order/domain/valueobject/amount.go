package valueobject

import "errors"

var ErrAmountNegative = errors.New("amount must not be negative")

type Amount int

func NewAmount(val int) (Amount, error) {
	if val < 0 {
		return 0, ErrAmountNegative
	}
	return Amount(val), nil
}

func (a Amount) Int() int {
	return int(a)
}
