package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

type DeliveryService string

var dsRe = regexp.MustCompile(`^[A-Z0-9]+$`)

func NewDeliveryService(val string) (DeliveryService, error) {
	val = strings.ToLower(val)
	if val == "" || !dsRe.MatchString(strings.ToUpper(val)) {
		return "", errors.New("invalid delivery service")
	}
	return DeliveryService(val), nil
}
