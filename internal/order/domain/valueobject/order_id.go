package valueobject

import "errors"

type OrderID string

func NewOrderID(val string) (OrderID, error) {
	if len(val) < 8 {
		return "", errors.New("order id too short")
	}
	return OrderID(val), nil
}
