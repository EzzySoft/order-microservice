package valueobject

import (
	"errors"
	"strings"
)

type DeliveryEmail string

func NewDeliveryEmail(val string) (DeliveryEmail, error) {
	if len(val) < 3 || !strings.Contains(val, "@") {
		return "", errors.New("invalid email")
	}
	return DeliveryEmail(val), nil
}
