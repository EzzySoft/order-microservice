package valueobject

import (
	"errors"
	"regexp"
)

type DeliveryPhone string

var phoneRe = regexp.MustCompile(`^\+\d{7,15}$`)

func NewDeliveryPhone(val string) (DeliveryPhone, error) {
	if !phoneRe.MatchString(val) {
		return "", errors.New("phone must be in format +[digits], min 7 max 15")
	}
	return DeliveryPhone(val), nil
}
